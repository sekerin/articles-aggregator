package main

import (
	"context"
	api "github.com/sekerin/artticles-aggregator/api/parser"
	domain "github.com/sekerin/artticles-aggregator/internal/articles"
	"github.com/sekerin/artticles-aggregator/internal/config"
	"github.com/sekerin/artticles-aggregator/internal/parser"
	"github.com/sekerin/artticles-aggregator/pkg/grpc/factory"
	"github.com/sekerin/artticles-aggregator/pkg/http"
	"google.golang.org/grpc"
	"log"
	"time"
)

var conf config.Options

type ParserServer struct {
}

func main() {
	var err error
	conf, err = config.NewConfig()
	if err != nil {
		log.Fatalln(err)
	}

	opts := factory.ServerOptions{
		Network: "tcp",
		Address: conf.ParserAddress,
	}

	err = factory.MakeServer(
		func(s *grpc.Server) {
			api.RegisterParserServer(s, &ParserServer{})
		},
		opts,
	)

	if err != nil {
		log.Fatalln(err)
	}
}

func (s ParserServer) Parse(ctx context.Context, request *api.ParseRequest) (*api.ParseResponse, error) {
	newParser := parser.NewParser(parser.Settings{Request: http.Request})
	timeoutContext, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	opts := parser.Options{
		PageUrl:  request.PageUrl,
		Title:    request.Title,
		FullText: request.FullText,
	}

	data, err := newParser.Parse(timeoutContext, request.Endpoint, opts)
	if err != nil {
		log.Println(err)
	}

	repo, err := domain.NewArticleRepository(domain.ConnectOptions{
		User:     conf.DBUser,
		Password: conf.DBPassword,
		Name:     conf.DBName,
		Hostname: conf.DBHostname,
		Port:     conf.DBPort,
	})
	if err != nil {
		log.Fatalln(err)
	}

	parsedCount := int32(0)
	for _, articleData := range data {
		if articleData.Err != nil {
			log.Println(articleData.Article)
		}

		err := repo.Save(articleData.Article)

		if err != nil {
			log.Println(err)
		}

		parsedCount++
	}

	response := &api.ParseResponse{
		Parsed: parsedCount,
	}

	return response, nil
}

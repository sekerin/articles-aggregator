package main

import (
	"context"
	api "github.com/sekerin/artticles-aggregator/api/articles"
	domain "github.com/sekerin/artticles-aggregator/internal/articles"
	"github.com/sekerin/artticles-aggregator/internal/config"
	"github.com/sekerin/artticles-aggregator/pkg/grpc/factory"
	"google.golang.org/grpc"
	"log"
)

var conf config.Options

type ArticleServer struct {
}

func main() {
	var err error
	conf, err = config.NewConfig()
	if err != nil {
		log.Fatalln(err)
	}

	opts := factory.ServerOptions{
		Network: "tcp",
		Address: conf.SearchAddress,
	}

	err = factory.MakeServer(
		func(s *grpc.Server) {
			api.RegisterArticlesServer(s, &ArticleServer{})
		},
		opts,
	)

	if err != nil {
		log.Fatalln(err)
	}
}

func (s ArticleServer) Search(ctx context.Context, request *api.SearchRequest) (*api.SearchResponse, error) {
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

	defer repo.Db.Close()

	findedArticles, err := repo.GetByTitle(request.Message)
	if err != nil {
		log.Println(err)
	}

	var out []string

	for _, article := range findedArticles {
		out = append(out, article.Title)
	}

	response := &api.SearchResponse{
		Message: out,
	}

	return response, nil
}

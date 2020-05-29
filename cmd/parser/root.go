package parser

import (
	"context"
	api "github.com/sekerin/artticles-aggregator/api/parser"
	"github.com/sekerin/artticles-aggregator/internal/config"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"log"
)

var (
	endpoint string
	pageUrl  string
	title    string
	fullText string
)

var Cmd = &cobra.Command{
	Use:   "parser",
	Short: "Run parse endpoint",
	Long:  "Run parser",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		opts := []grpc.DialOption{
			grpc.WithInsecure(),
		}

		cnf, err := config.NewConfig()
		if err != nil {
			log.Fatalln(err)
		}

		conn, err := grpc.Dial(cnf.ParserAddress, opts...)
		if err != nil {
			log.Fatalln(err)
		}

		defer func() {
			err := conn.Close()

			if err != nil {
				log.Fatalln(err)
			}
		}()

		client := api.NewParserClient(conn)

		request := &api.ParseRequest{
			Endpoint: endpoint,
			PageUrl:  pageUrl,
			Title:    title,
			FullText: fullText,
		}

		parser, err := client.Parse(ctx, request)

		if err != nil {
			log.Fatalln(err)
		}

		log.Println(parser.Parsed)
	},
}

func init() {
	Cmd.PersistentFlags().StringVar(&endpoint, "endpoint", "", "Endpoint for start parse")
	if err := Cmd.MarkPersistentFlagRequired("endpoint"); err != nil {
		log.Fatalln(err)
	}

	Cmd.PersistentFlags().StringVar(&pageUrl, "urls", "", "XPath for get article urls")
	if err := Cmd.MarkPersistentFlagRequired("urls"); err != nil {
		log.Fatalln(err)
	}

	Cmd.PersistentFlags().StringVar(&title, "title", "", "XPath for get title from article page")
	if err := Cmd.MarkPersistentFlagRequired("title"); err != nil {
		log.Fatalln(err)
	}

	Cmd.PersistentFlags().StringVar(&fullText, "text", "", "XPath for get article text")
	if err := Cmd.MarkPersistentFlagRequired("text"); err != nil {
		log.Fatalln(err)
	}
}

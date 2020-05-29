package search

import (
	"context"
	api "github.com/sekerin/artticles-aggregator/api/articles"
	"github.com/sekerin/artticles-aggregator/internal/config"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"log"
)

var (
	search string
)

var Cmd = &cobra.Command{
	Use:   "search",
	Short: "Search by title",
	Long:  "Search by title",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		opts := []grpc.DialOption{
			grpc.WithInsecure(),
		}

		cnf, err := config.NewConfig()
		if err != nil {
			log.Fatalln(err)
		}

		conn, err := grpc.Dial(cnf.SearchAddress, opts...)
		if err != nil {
			log.Fatalln(err)
		}

		defer func() {
			err := conn.Close()

			if err != nil {
				log.Fatalln(err)
			}
		}()

		client := api.NewArticlesClient(conn)

		request := &api.SearchRequest{
			Message: search,
		}

		response, err := client.Search(ctx, request)

		if err != nil {
			log.Fatalln(err)
		}

		for _, mess := range response.Message {
			log.Println(mess)
		}
	},
}

func init() {
	Cmd.PersistentFlags().StringVar(&search, "search", "", "Search string")
	if err := Cmd.MarkPersistentFlagRequired("search"); err != nil {
		log.Fatalln(err)
	}
}

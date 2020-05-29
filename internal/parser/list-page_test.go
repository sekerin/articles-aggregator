package parser

import (
	"context"
	"github.com/sekerin/artticles-aggregator/internal/articles"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParse(t *testing.T) {
	dataPage := `<html><body><div id="url">urlToArticle<div></body></html>`
	dataArticlePage := `<html><body><div><div id="title">Title</div></div><div><div id="text">Text</div></div></body></html>`

	result := Data{
		Article: articles.Article{
			Title: "Title",
			Url:   "urlToArticle",
			Text:  "Text",
		},
		Err: nil,
	}

	opts := Options{
		PageUrl:  "//*/div[@id='url']",
		Title:    "//*/div[@id='title']",
		FullText: "//*/div[@id='text']",
	}

	ctx := context.Background()

	lp := ListPage{request: func(ctx context.Context, url string) (string, error) {
		if url == "urlToArticle" {
			return dataArticlePage, nil
		} else {
			return dataPage, nil
		}
	}}

	data, err := lp.Parse(ctx, "url", opts)
	require.NoError(t, err)

	for _, item := range data {
		require.Equal(t, result, item)
	}
}

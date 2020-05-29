package parser

import (
	"context"
	"github.com/sekerin/artticles-aggregator/internal/articles"
)

type Request func(ctx context.Context, url string) (string, error)

type Options struct {
	PageUrl  string
	Title    string
	FullText string
}

type Settings struct {
	Request Request
}

type Data struct {
	Article articles.Article
	Err     error
}

type Parser interface {
	Parse(url string, opts []Options)
}

func NewParser(settings Settings) ListPage {
	return ListPage{request: settings.Request}
}

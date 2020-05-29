package parser

import (
	"context"
	_ "github.com/lib/pq"
	"github.com/sekerin/artticles-aggregator/internal/articles"
	"gopkg.in/xmlpath.v2"
	"strings"
)

type ListPage struct {
	request Request
}

func (lp ListPage) Parse(ctx context.Context, url string, opts Options) ([]Data, error) {
	list, err := lp.request(ctx, url)

	if err != nil {
		return nil, err
	}

	urls, err := lp.getPagesUrl(list, opts)

	if err != nil {
		return nil, err
	}

	ch := make(chan Data, len(urls))

	for _, itemUrl := range urls {
		go lp.getArticlesAsync(ctx, ch, itemUrl, opts)
	}

	results := make([]Data, len(urls))

	for i := range results {
		results[i] = <-ch
	}

	return results, nil
}

func (lp *ListPage) getPagesUrl(data string, opts Options) ([]string, error) {
	url, err := xmlpath.Compile(opts.PageUrl)
	if err != nil {
		return nil, err
	}

	root, err := xmlpath.ParseHTML(strings.NewReader(data))

	if err != nil {
		return nil, err
	}

	iter := url.Iter(root)
	var urls []string

	for iter.Next() {
		url := iter.Node().String()

		urls = append(urls, url)
	}

	return urls, nil
}

func (lp *ListPage) getArticlesAsync(ctx context.Context, ch chan Data, url string, opts Options) {
	var resp Data
	data, err := lp.request(ctx, url)

	if err != nil {
		resp.Err = err
		ch <- resp
	}

	article, err := lp.getArticles(data, opts)

	if err != nil {
		resp.Err = err
		ch <- resp
	}

	resp.Article = article
	resp.Article.Url = url
	ch <- resp
}

func (lp *ListPage) getArticles(data string, opts Options) (articles.Article, error) {
	title, err := xmlpath.Compile(opts.Title)
	if err != nil {
		return articles.Article{}, err
	}

	fullText, err := xmlpath.Compile(opts.FullText)
	if err != nil {
		return articles.Article{}, err
	}

	root, err := xmlpath.ParseHTML(strings.NewReader(data))

	article := articles.Article{}

	if err != nil {
		return articles.Article{}, err
	}

	if value, ok := title.String(root); ok {
		article.Title = value
	}

	if value, ok := fullText.String(root); ok {
		article.Text = value
	}

	return article, nil
}

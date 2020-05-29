package http

import (
	"context"
	"crypto/tls"
	"io/ioutil"
	"net/http"
)

func Request(ctx context.Context, url string) (string, error) {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{},
	}
	client := &http.Client{Transport: transport}
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	response, err := client.Do(req)

	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}

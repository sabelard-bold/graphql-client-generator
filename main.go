package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	libflag "github.com/Wryte/graphql-client-generator/flag"
	generate "github.com/Wryte/graphql-client-generator/generate"
	graphql "github.com/Wryte/graphql-client-generator/graphql"
)

var (
	endpoint string
	headers  libflag.HTTPHeaders
)

type schemaWrap struct {
	Data struct {
		Schema graphql.Schema `json:"__schema"`
	} `json:"data"`
}

func main() {
	flag.StringVar(&endpoint, "endpoint", "", "graphql endpoint for introspection query")
	flag.Var(&headers, "header", "header for introspection query")
	flag.Parse()

	bytes, err := getSchema()
	if err != nil {
		panic(fmt.Errorf("getting schema: %w", err))
	}

	w := schemaWrap{}

	err = json.Unmarshal(bytes, &w)
	if err != nil {
		panic(fmt.Errorf("unmarshalling json: %w", err))
	}

	err = generate.Write(os.Stdout, w.Data.Schema)

	if err != nil {
		panic(fmt.Errorf("generating files: %w", err))
	}
}

func getSchema() ([]byte, error) {
	schemaQueryFile, err := os.Open("schemaQuery.graphql")

	if err != nil {
		return nil, fmt.Errorf("reading schema file: %w", err)
	}
	defer schemaQueryFile.Close()

	bytes, err := ioutil.ReadAll(schemaQueryFile)
	if err != nil {
		return nil, fmt.Errorf("reading whole file: %w", err)
	}

	url, err := url.Parse(endpoint)
	if err != nil {
		return nil, fmt.Errorf("parsing url: %w", err)
	}

	req := graphql.Request{
		URL:   url,
		Query: string(bytes),
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
	}

	for _, h := range headers {
		req.Header.Add(h.Name, h.Value)
	}

	c := graphql.Client{}

	bytes, err = c.Query(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("running introspection query: %w", err)
	}

	return bytes, nil
}

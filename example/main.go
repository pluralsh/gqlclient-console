package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	gqlclient "github.com/pluralsh/console-client-go"
)

type authedTransport struct {
	key     string
	wrapped http.RoundTripper
}

func (t *authedTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "Token "+t.key)
	return t.wrapped.RoundTrip(req)
}

func main() {
	var err error
	defer func() {
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()

	key := "deploy-h89xne0p2efknbud7n26nzc2rym9anhwppj3mifwrc2bzyiafe"

	httpClient := http.Client{
		Transport: &authedTransport{
			key:     key,
			wrapped: http.DefaultTransport,
		},
	}
	graphqlClient := gqlclient.NewClient(&httpClient, "https://console.cdaws.onplural.sh/gql/ext")
	meResp, err := graphqlClient.ListClusterRestore(context.Background(), "daa2b5f7-a0e7-4e11-896a-6e5389b3d8aa")
	if err != nil {
		return
	}
	fmt.Println("my cluster", meResp)

}

package traqforest

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/traPtitech/go-traq"
)

func TestNewForest(t *testing.T) {
	// load ACCESS_TOKEN from .env
	err := godotenv.Load()
	if err != nil {
		t.Fatal(err)
	}
	// create api_client
	api_key := traq.APIKey{
		Prefix: "Bearer",
		Key:    os.Getenv("ACCESS_TOKEN"),
	}

	api_conf := traq.NewConfiguration()
	api_conf.DefaultHeader["Authorization"] = fmt.Sprintf("%s %s", api_key.Prefix, api_key.Key)
	api_client := traq.NewAPIClient(api_conf)
	// create context
	ctx := context.Background()
	// create forest
	forest, err := NewForest(api_client, &ctx)
	if err != nil {
		t.Fatal(err)
	}
	// get all channels
	channels_request := api_client.ChannelApi.GetChannels(ctx)
	channels, _, err := channels_request.Execute()
	if err != nil {
		t.Fatal(err)
	}
	// get all channels path
	for _, channel := range channels.Public {
		path, ok := forest.GetPath(channel.Id)
		if !ok {
			t.Fatal("GetPath failed")
		}
		fmt.Println(path)
	}
}

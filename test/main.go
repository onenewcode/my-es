package main

import (
	"context"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

func main() {

	typedClient, err := elasticsearch.NewTypedClient(elasticsearch.Config{
		CloudID: "https://my-deployment-e278e0.es.us-central1.gcp.cloud.es.io",
		APIKey:  "essu_V1VOSU4zaHZORUpWWldVMFJFVlFWVEZRWkZZNmVraHVSV3RTUW1KU1RFZEdTRVJ4VlZabWNVTkNVUT09AAAAAOgzkUM=",
	})
	if err != nil {
		fmt.Println("es 连接失败")
	}
	// 创建索引库，类似mysql中的数据库
	typedClient.Indices.Create("my_index").Do(context.TODO())
	// 创建一条

	document := struct {
		Name string `json:"name"`
	}{
		"go-elasticsearch",
	}
	// 创建一条数据
	_, err = typedClient.Index("my_index").
		Id("1").
		Request(document).
		Do(context.TODO())
	if err != nil {
		return
	}
	// 从my_index索引库中，获取一个id为1的文档
	typedClient.Get("my_index", "1").Do(context.TODO())

	// // 从my_index索引库中，获取一个id为1的文档
	typedClient.Search().
		Index("my_index").
		Request(&search.Request{
			Query: &types.Query{MatchAll: &types.MatchAllQuery{}},
		}).
		Do(context.TODO())
}

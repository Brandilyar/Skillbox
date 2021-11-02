package main

import (
	"context"
	"fmt"

	//"http_service/pkg/api"
	"http_service/pkg/api"
	"http_service/pkg/repositores"
)

func main() {
	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()
	client, err := repositores.NewDateBase(context.TODO())
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	users := client.Database("Users").Collection("users")
	api := api.NewApiUser(context.TODO(), users)
	if err := api.Run("8080"); err != nil {
		fmt.Println(err)
	}

}

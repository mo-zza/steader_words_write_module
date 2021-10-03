package main

import (
	"context"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/mo-zza/steader_words_write_module/application/handler"
	"github.com/mo-zza/steader_words_write_module/dependency"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
		return
	}

	app, err := dependency.NewConnection()
	if err != nil {
		fmt.Println("%s", err.Error())
		return 
	}

	client, err := app.Firestore(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	defer client.Close()

	err = handler.Save()
	if err != nil {
		fmt.Println(err)
	}
}

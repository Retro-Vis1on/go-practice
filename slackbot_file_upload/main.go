package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Env error")
		return
	}
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"test.pdf"}

	for _, file := range fileArr {
		params := slack.FileUploadParameters{Channels: channelArr, File: file}
		fileUp, err := api.UploadFile(params)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Name:%s\n", fileUp.Name)
	}
}

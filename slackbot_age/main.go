package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"
)

func printCommandEvents(analytics <-chan *slacker.CommandEvent) {
	for event := range analytics {
		fmt.Println("Command Event")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Env error")
		return
	}
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))
	go printCommandEvents(bot.CommandEvents())
	bot.Command("My yob is <year>", &slacker.CommandDefinition{
		Description: "Calculates the Age ",
		Example:     "My yob is 2020",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				fmt.Println("Error in parsing")
			}
			curYear, _, _ := time.Now().Date()
			age := curYear - yob
			r := fmt.Sprintf("Age is %d", age)
			response.Reply(r)

		},
	})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if err := bot.Listen(ctx); err != nil {
		log.Fatal(err)
	}
}

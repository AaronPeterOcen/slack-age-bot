package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent){
	for  event := range analyticsChannel {
		fmt.Println("command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		// fmt.Println(event.Event.ChannelID)
	}
}

func main(){
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-2883289740131-6686343114342-Iwy9G9I1XJn8I2f5imLG3Wql")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A06L6970DCN-6692846912995-2e4b42386bcd01c685667e66ad5a76844dc56cbeffdc1a24441f264368133713")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("YOB is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		// Examples: "yob is 2010",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)

			if err != nil {
				fmt.Println("error")
			}
			age := 2023-yob
			r := fmt.Sprintf("age is %d", age)
			response.Reply(r)

		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil{
		log.Fatal(err)
	}
}
package main

import (
	"gfeed/bot"
	"gfeed/utils"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ev events.CloudWatchAlarmTrigger) error {
	token := utils.GetEnv("TELEGRAM_TOKEN", "")
	user := utils.GetEnv("TELEGRAM_USER", "")
	channel := utils.GetEnv("TELEGRAM_CHANNEL", "@GamerFeed")

	bot.SendNews(bot.Config{
		Token:   token,
		Channel: channel,
		User:    user,
		DryRun:  false,
	})

	return nil
}

func main() {
	lambda.Start(handler)
}

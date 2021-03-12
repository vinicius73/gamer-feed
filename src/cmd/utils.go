package cmd

import "gfeed/domains/bot"

func getBotConfig() bot.Config {
	cfg := bot.Config{
		Token:   token,
		Channel: channel,
		User:    user,
		DryRun:  dryRun,
	}

	return cfg
}

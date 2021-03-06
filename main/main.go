package main

import (
	"fmt"
	"os"

	"github.com/Goryudyuma/GithHubActionsTest/main/actiontype"
	slack5min "github.com/Goryudyuma/GithHubActionsTest/main/slack/5min"
)

func readArgument() actiontype.Argument {
	return actiontype.NewArgument(
		os.Getenv("SLACK_WEBHOOK_ULR_5MIN"),
	)
}

func main() {
	args := readArgument()

	if err := slack5min.Run(args.GetSlackWebhookURL5min()); err != nil {
		fmt.Errorf("error: %v", err)
	}

	fmt.Println("Hello golang!")
}

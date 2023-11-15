package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	// Set Slack bot token and channel ID as environment variables
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-6196336599091-6182383703927-kHVn2Zewliau2oMdjFanlXan")
	os.Setenv("CHANNEL_ID", "C065PEW8WMB")

	// Create a new Slack API client
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))

	// Define the channel ID and file names
	channelID := os.Getenv("CHANNEL_ID")
	fileArr := []string{"MGTS.pdf"}

	// Iterate through each file in the fileArr
	for i := 0; i < len(fileArr); i++ {
		// Set up parameters for file upload
		params := slack.FileUploadParameters{
			Channels: []string{channelID},
			File:     fileArr[i],
		}

		// Upload the file to Slack
		file, err := api.UploadFile(params)
		if err != nil {
			// Handle the error and print a message
			fmt.Printf("Error uploading file: %s\n", err)
			return
		}

		// Print information about the uploaded file
		fmt.Printf("Name: %s, URL: %s\n", file.Name, file.URL)
	}
}

package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/nicklaw5/helix"
	"github.com/gempir/go-twitch-irc"
)

func main() {
	// Replace with your own Twitch client ID and secret
	clientID := "your_client_id"
	clientSecret := "your_client_secret"

	// Replace with your own Twitch username and OAuth token
	username := "your_username"
	oauthToken := "oauth:your_oauth_token"

	// Connect to Twitch IRC server
	client := twitch.NewClient(username, oauthToken)

	// Join a Twitch channel
	channel := "your_channel"
	client.Join(channel)

	// Create a Helix client
	helixClient, err := helix.NewClient(&helix.Options{
		ClientID:     clientID,
		ClientSecret: clientSecret,
	})
	if err != nil {
		log.Fatalf("Error creating Helix client: %s", err.Error())
	}

	// Listen for incoming messages
	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		// Print the message to the console
		fmt.Printf("[%s] %s: %s\n", message.Channel, message.User.DisplayName, message.Message)

		// Check if the message is a command
		if strings.HasPrefix(message.Message, "!game") {
			// Get the current game being played on the channel
			game, err := helixClient.GetGames(&helix.GamesParams{
				Names: []string{channel},
			})
			if err != nil {
				log.Printf("Error getting game: %s", err.Error())
				return
			}

			// Send a response to the chat
			client.Say(message.Channel, fmt.Sprintf("The current game is %s", game.Data.Games[0].Name))
		}
	})

	// Start the client
	err = client.Connect()
	if err != nil {
		log.Fatalf("Error connecting to Twitch IRC server: %s", err.Error())
	}

	// Wait for incoming messages
	err = helixClient.Run()
	if err != nil {
		log.Fatalf("Error running Helix client: %s", err.Error())
	}
}
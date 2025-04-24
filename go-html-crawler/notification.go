package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func SendTelegramNotification(message string) error {
	// Build the Telegram API URL
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", BOT_TOKEN)

	// Create request with URL-encoded parameters
	resp, err := http.PostForm(apiURL, url.Values{
		"chat_id": {CHAT_ID},
		"text":    {message},
	})
	if err != nil {
		return fmt.Errorf("failed to send Telegram notification: %w", err)
	}
	defer resp.Body.Close()

	// Check if the request was successful
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Telegram API returned non-200 status code: %d", resp.StatusCode)
	}

	return nil
}

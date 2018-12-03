package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"time"
)

type Leaderboard struct {
	Event   string         `json:"event"`
	OwnerID string         `json:"owner_id"`
	Members map[int]Member `json:"members"`
}
type Member struct {
	ID                 string         `json:"id"`
	CompletionDayLevel map[string]Day `json:"completion_day_level"`
	GlobalScore        int            `json:"global_score"`
	LocalScore         int            `json:"local_score"`
	Name               string         `json:"name"`
	Stars              int            `json:"stars"`
}
type Day map[string]Part
type Part struct {
	GetStarTs string `json:"get_star_ts"`
}

func main() {
	flagYear := flag.Int("year", time.Now().Year(), "Leaderboard Year")
	flagLeaderboardID := flag.Int("id", 0, "Leaderboard ID")
	flagSessionCookie := flag.String("session", "", "Session Cookie")
	flagSlackWebhookURL := flag.String("slack-webhook-url", "", "Slack Incoming Webhook URL")
	flagSlackChannel := flag.String("slack-channel", "", "Slack Channel")
	flagSlackIconEmoji := flag.String("slack-icon-emoji", "", "Slack Icon Emoji")
	flagSlackUsername := flag.String("slack-username", "", "Slack Username")

	flag.Parse()

	year := *flagYear
	leaderboardID := *flagLeaderboardID
	sessionCookie := *flagSessionCookie

	if year == 0 || leaderboardID == 0 || sessionCookie == "" || *flagSlackWebhookURL == "" {
		flag.Usage()
		return
	}

	leaderboard, err := getLeaderboard(year, leaderboardID, sessionCookie)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	for {
		time.Sleep(60 * time.Second)

		newLeaderboard, err := getLeaderboard(year, leaderboardID, sessionCookie)
		if err != nil {
			fmt.Println("error:", err)
			continue
		}

		for memberNumber, member := range newLeaderboard.Members {
			for dayNumber, day := range member.CompletionDayLevel {
				for partNumber := range day {
					publish := false
					if _, ok := leaderboard.Members[memberNumber]; !ok {
						message := fmt.Sprintf(
							"%s joined the <%s|leaderboard>.",
							member.Name,
							leaderboardURL(year, leaderboardID),
						)
						fmt.Println(message)
						publishToSlack(
							*flagSlackChannel,
							*flagSlackIconEmoji,
							*flagSlackUsername,
							message,
							*flagSlackWebhookURL,
						)
					} else if _, ok := leaderboard.Members[memberNumber].CompletionDayLevel[dayNumber]; !ok {
						publish = true
					} else if _, ok := leaderboard.Members[memberNumber].CompletionDayLevel[dayNumber][partNumber]; !ok {
						publish = true
					}
					if publish {
						message := fmt.Sprintf(
							"%s completed day %s part %s (<%s|leaderboard>).",
							member.Name,
							dayNumber,
							partNumber,
							leaderboardURL(year, leaderboardID),
						)
						fmt.Println(message)
						publishToSlack(
							*flagSlackChannel,
							*flagSlackIconEmoji,
							*flagSlackUsername,
							message,
							*flagSlackWebhookURL,
						)
					}
				}
			}
		}

		leaderboard = newLeaderboard
	}
}

func leaderboardURL(year, id int) string {
	return fmt.Sprintf("https://adventofcode.com/%d/leaderboard/private/view/%d.json", year, id)
}

func getLeaderboard(year, id int, sessionCookie string) (*Leaderboard, error) {
	url := leaderboardURL(year, id)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.AddCookie(&http.Cookie{Name: "session", Value: sessionCookie})

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	leaderboard := &Leaderboard{}
	err = dec.Decode(leaderboard)
	if err != nil {
		return nil, err
	}

	return leaderboard, nil
}

func publishToSlack(channel, iconEmoji, username, message, slackWebhookURL string) error {
	type Payload struct {
		Channel   string `json:"channel"`
		IconEmoji string `json:"icon_emoji"`
		Username  string `json:"username"`
		Text      string `json:"text"`
	}

	payload := Payload{
		Channel:   channel,
		IconEmoji: iconEmoji,
		Username:  username,
		Text:      message,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", slackWebhookURL, bytes.NewReader(payloadBytes))
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	resp.Body.Close()
	return nil
}

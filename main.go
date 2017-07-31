package main

import (
	"log"
	"os"

	"github.com/nlopes/slack"
)

type SlackInviter struct {
	api   *slack.Client
	from  slack.Channel
	to    slack.Channel
	users map[string]string
}

func (slack *SlackInviter) setUsers() {
	users, err := slack.api.GetUsers()
	if err != nil {
		log.Fatalf("%s\n", err)
	}

	slack.users = make(map[string]string)
	for _, user := range users {
		slack.users[user.ID] = user.Name
	}
}

func (slack *SlackInviter) setChannels(from, to string) {
	channels, err := slack.api.GetChannels(true)
	if err != nil {
		log.Fatalf("%s\n", err)
	}

	for _, channel := range channels {
		if channel.Name == from {
			slack.from = channel
		} else if channel.Name == to {
			slack.to = channel
		}
	}
}

func NewSlackInviter(from, to string) *SlackInviter {
	s := &SlackInviter{api: slack.New(os.Getenv("SLACK_TOKEN"))}
	s.setChannels(from, to)
	s.setUsers()
	return s
}

func main() {
	if len(os.Args) != 3 {
		log.Fatalln(os.ErrInvalid.Error())
	}
	slack := NewSlackInviter(os.Args[1], os.Args[2])

	var toChanMembers = strSliceToMap(slack.to.Members)
	for _, member := range slack.from.Members {
		if _, ok := toChanMembers[member]; !ok {
			if _, err := slack.api.InviteUserToChannel(slack.to.ID, member); err != nil {
				log.Printf("Failed to invite %s (%s) to #%s (%s)",
					slack.users[member], member, slack.to.Name, slack.to.ID)
			} else {
				log.Printf("Invited %s (%s) to #%s (%s)",
					slack.users[member], member, slack.to.Name, slack.to.ID)
			}
		}
	}
}

func strSliceToMap(slice []string) map[string]struct{} {
	s := make(map[string]struct{}, len(slice))
	for _, str := range slice {
		s[str] = struct{}{}
	}
	return s
}

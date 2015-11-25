package main

import (
	"fmt"
	"regexp"
	"strings"
	"os"
	"github.com/nlopes/slack"
)

const (
	botID = "U0F8EH14Y"
	botChannelID = "C0F8FQ2GH"
	mxTeamID = "T0F8CMAU9"
)

func checkMessage(msg string) (string, bool) {
	r := regexp.MustCompile("^<@([\\d\\w]+)>:(.*)")
	s := r.FindStringSubmatch(msg)
	return strings.TrimSpace(s[2]), s[1] == botID
}

func handleCommand(sender, cmd string) {
	fmt.Printf("用户 %s 发送指令 %s\n", sender, cmd)
	return
}

func handleMessage(rtm *slack.RTM) {
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			fmt.Print("Event Received: ")
			switch evt := msg.Data.(type) {
			case *slack.HelloEvent:
			// Ingore hello
			case *slack.ConnectedEvent:
				fmt.Println("Info:", evt.Info)
				fmt.Println("Connection counter:", evt.ConnectionCount)
				rtm.SendMessage(rtm.NewOutgoingMessage("Hello world", "#general"))
			case *slack.MessageEvent:
				fmt.Printf("Message: %v\n", evt)
				if evt.Channel == botChannelID && evt.Team == mxTeamID {
					cmd, sendToMXBot := checkMessage(evt.Text)
					if sendToMXBot {
						//fmt.Printf("接收到机器人指令： %v", evt)
						go handleCommand(evt.User, cmd)
					}
				}
			case *slack.PresenceChangeEvent:
				fmt.Printf("Presence Change: %v\n", evt)
			case *slack.LatencyReport:
				fmt.Printf("Current latency: %v\n", evt.Value)
			case *slack.RTMError:
				fmt.Printf("Error: %s\n", evt.Error())
			case *slack.InvalidAuthEvent:
				fmt.Printf("Invalid credentials")
				return
			default:
			// Ignore other events...
			}
		}
	}
}

func main() {
	token := os.Getenv("SLACK_BOT_TOKEN")
	api := slack.New(token)
	api.SetDebug(true)
	rtm := api.NewRTM()
	go rtm.ManageConnection()
	handleMessage(rtm)
}

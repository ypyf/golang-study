package main

import (
	"fmt"
	"github.com/nlopes/slack"
	"os"
	"regexp"
	"strings"
)

const (
	botID = "U0F8EH14Y"
	botDMChannelID = "D0F8EH15W"
	botChannelID = "C0F8FQ2GH"
	mxTeamID = "T0F8CMAU9"
)

type Message struct {
	*slack.RTM
	text    string
	channel string
}

func (m *Message) Send(msg string) {
	m.SendMessage(m.NewOutgoingMessage(msg, m.channel))
}

func (m *Message) Done() {

}

type Commander interface {
	Matches(text string) bool
	Respond(msg *Message)
	Help() string
}

type Help struct {
	sender string // sender id
}

func (self *Help) Matches(text string) bool {
	return strings.HasPrefix(text, "help")
}

func (self *Help) Respond(msg *Message) {
	tokens := strings.Fields(msg.text)
	if len(tokens) >= 2 {
		for _, v := range botCommands {
			if v.Matches(tokens[1]) {
				msg.Send(v.Help())
				break
			}
		}
	} else {
		msg.Send("当前可用的命令:\nhelp\nhello\n")
	}
	msg.Done()
}

func (self *Help) Help() string {
	return "你可以尝试输入 help <cmd> 来获得各种命令的帮助信息."
}

func checkMessage(msg string) (string, bool) {
	r := regexp.MustCompile("^<@([\\d\\w]+)>:(.*)")
	s := r.FindStringSubmatch(msg)
	if len(s) == 0 {
		return "", false
	}
	return s[2], s[1] == botID
}

func handleCommand(rtm *slack.RTM, channel, sender, text string) {
	//fmt.Printf("用户 %s 发送指令 %s\n", sender, text)

	//target = "<@" + sender + ">: "

	for _, v := range botCommands {
		if v.Matches(text) {
			v.Respond(&Message{rtm, text, channel})
			break
		}
	}
	return
}

type Hello struct {
	name string
}

func (w *Hello) Matches(text string) bool {
	return strings.HasPrefix(text, "hello")
}

func (w *Hello) Respond(msg *Message) {
	msg.Send(fmt.Sprintf("你好，我是机器人 %s，有什么需要我帮助的吗？\n", w.name))
	msg.Done()
}

func (w *Hello) Help() string {
	return "打招呼."
}

type BotCommands []Commander
var botCommands BotCommands

func handleMessage(rtm *slack.RTM) {
	// 注册命令处理器
	user, err := rtm.GetUserInfo(botID)
	if err != nil {
		fmt.Println(err)
		return
	}
	fullname := user.Profile.FirstName + " " + user.Profile.LastName
	botCommands = append(botCommands, &Help{})
	botCommands = append(botCommands, &Hello{name: fullname})

	for {
		select {
		case msg := <-rtm.IncomingEvents:
			fmt.Print("Event Received: ")
			switch evt := msg.Data.(type) {
			case *slack.HelloEvent:
			// Ignore hello
			case *slack.ConnectedEvent:
				fmt.Println("Info:", evt.Info)
				fmt.Println("Connection counter:", evt.ConnectionCount)
				rtm.SendMessage(rtm.NewOutgoingMessage("Hello world", "#general"))
			case *slack.MessageEvent:
				fmt.Printf("Message: %v\n", evt)
				if evt.Channel == botChannelID && evt.Team == mxTeamID {
					text, sendToMXBot := checkMessage(evt.Text)
					if sendToMXBot && len(text) > 0 {
						//fmt.Printf("接收到机器人指令： %v", evt)
						go handleCommand(rtm, evt.Channel, evt.User, strings.TrimSpace(text))
					}
				} else if evt.Channel == botDMChannelID && evt.Team == mxTeamID {
					go handleCommand(rtm, evt.Channel, evt.User, strings.TrimSpace(evt.Text))
				}
			case *slack.ChannelJoinedEvent:
			// Ignore
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
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	api.SetDebug(true)
	rtm := api.NewRTM()
	go rtm.ManageConnection()
	handleMessage(rtm)
}

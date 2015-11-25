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
	botDMChannelID = "D0F8EH15W"
	botChannelID = "C0F8FQ2GH"
	mxTeamID = "T0F8CMAU9"
)

//func initSego() sego.Segmenter {
//	// 载入词典
//	var segmenter sego.Segmenter
//	segmenter.LoadDictionary("E:/projects/Go/src/github.com/huichen/sego/data/dictionary.txt")
//	//	segmenter := initSego()
//	//	segments := segmenter.Segment([]byte(cmd))
//	//	output := sego.SegmentsToString(segments, false)
//	//	//输出分词后的成语
//	//	for _, v := range strings.Fields(output) {
//	//		fmt.Println(v)
//	//	}
//	return segmenter
//}

func checkMessage(msg string) (string, bool) {
	r := regexp.MustCompile("^<@([\\d\\w]+)>:(.*)")
	s := r.FindStringSubmatch(msg)
	if len(s) == 0 {
		return "", false
	}
	return s[2], s[1] == botID
}

//
// show consul members 显示所有consul节点
//
//
const help = `
	me -- 显示当前用户信息
	show all nodes -- 显示 consul 节点信息
	show node health -- 显示节点健康监控
`
func handleCommand(rtm *slack.RTM, dm bool, sender, cmd string) {
	fmt.Printf("用户 %s 发送指令 %s\n", sender, cmd)
	tokens := strings.Fields(cmd)
	var target, channel string
	if dm {
		target = ""
		channel = botDMChannelID
	} else {
		target = "<@" + sender + ">: "
		channel = botChannelID
	}

	if strings.HasPrefix(tokens[0], "你是谁") || strings.HasPrefix(tokens[0], "你叫什么") {
		user, _ := rtm.GetUserInfo(botID)
		fullname := user.Profile.FirstName + " " + user.Profile.LastName
		rtm.SendMessage(rtm.NewOutgoingMessage(target + "我是 " + fullname, channel))
		return
	}

	if strings.HasPrefix(tokens[0], "我是谁") {
		user, _ := rtm.GetUserInfo(sender)
		fullname := user.Profile.FirstName + " " + user.Profile.LastName
		reply := fmt.Sprintf("你是 %s，以下是你的个人资料:\nID: %s\n用户名: %s\nEmail: %s\n电话: %s\n",
			fullname, user.ID, user.Name, user.Profile.Email, user.Profile.Phone)
		rtm.SendMessage(rtm.NewOutgoingMessage(target + reply, channel))
		return
	}

	switch tokens[0] {
	case "help":
		rtm.SendMessage(rtm.NewOutgoingMessage(target + help, channel))
	default:
		// echo unknown command
		rtm.SendMessage(rtm.NewOutgoingMessage(target + cmd, channel))
	}

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
					if sendToMXBot && len(cmd) > 0 {
						//fmt.Printf("接收到机器人指令： %v", evt)
						go handleCommand(rtm, false, evt.User, cmd)
					}
				} else if evt.Channel == botDMChannelID && evt.Team == mxTeamID {
					go handleCommand(rtm, true, evt.User, evt.Text)
				}
			case *slack.ChannelJoinedEvent:
			//
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

package douyu

import (
	"fmt"
	"github.com/songtianyi/barrage/douyu"
	"github.com/songtianyi/rrframework/logs"
)

const (
	xiao8  = 64609
	weiwei = 293829
)

func chatmsg(msg *douyu.Message) {
	level := msg.GetStringField("level")
	nn := msg.GetStringField("nn")
	txt := msg.GetStringField("txt")
	logs.Info(fmt.Sprintf("level(%s) - %s >>> %s", level, nn, txt))
}

func main() {
	client, err := douyu.Connect("openbarrage.douyutv.com:8601", nil)
	if err != nil {
		logs.Error(err)
		return
	}

	client.HandlerRegister.Add("chatmsg", douyu.Handler(chatmsg), "chatmsg")
	if err := client.JoinRoom(xiao8); err != nil {
		logs.Error(fmt.Sprintf("Join room fail, %s", err.Error()))
		return
	}
	client.Serve()
}

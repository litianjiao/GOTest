package douyu

import (
	//"bufio"
	"fmt"
	//"github.com/derekparker/delve/dwarf/reader"
	"github.com/songtianyi/barrage/douyu"
	"github.com/songtianyi/rrframework/logs"
	//	"os"
)

const (
	yyf     = 58428
	xiao8   = 64609
	weiwei  = 293829
	wt55kai = 138286
	momo    = 1975380
)

func chatmsg(msg *douyu.Message) {
	level := msg.GetStringField("level")
	nn := msg.GetStringField("nn")
	txt := msg.GetStringField("txt")
	logs.Info(fmt.Sprintf("level(%s) - %s >>> %s", level, nn, txt))
}

func main() {
	var str string
	fmt.Printf("choose your MC room.\n\r01 yyf\n\r02 xiao8\n\r03 55kai\n\r04 momo\n\r ")
	room_id := yyf
	fmt.Scanf("%sd", &str)
	switch str {
	case "01":
		room_id = yyf
	case "02":
		room_id = xiao8
	case "03":
		room_id = wt55kai
	case "04":
		room_id = momo
	case "05":
		room_id = weiwei
	default:
		fmt.Println("err chosen. please input again")

	}

	client, err := douyu.Connect("openbarrage.douyutv.com:8601", nil)
	if err != nil {
		logs.Error(err)
		return
	}

	client.HandlerRegister.Add("chatmsg", douyu.Handler(chatmsg), "chatmsg")
	if err := client.JoinRoom(room_id); err != nil {
		logs.Error(fmt.Sprintf("Join room fail, %s", err.Error()))
		return
	}
	client.Serve()
}

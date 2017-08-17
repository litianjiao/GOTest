package smog

import (
	"bufio"
	"fmt"

	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

/*
******************************************************************
  * @brief     loader device config
  * @param
  * @ret
  * @author    TRoy
  * @date      2017/8/4 2:45
******************************************************************
*/
func Loader() (new_dev Device) {
	var (
		dev_name string
		sn       string
		err      error
		dev_type int
	)

	for {
		var flag bool
		reader := bufio.NewReader(os.Stdin)
		sn, err = reader.ReadString('\n')
		if err != nil {
			log.Println("输入流转换异常", err)
		}
		//	fmt.Println("转换前:", sn)
		sn = strings.Replace(sn, " ", "", -1) //去除空格
		//	fmt.Println("转换后:", sn)
		sn1 := []byte(sn)
		dev_type = int(sn1[3] - 48)
		flag = true
		switch dev_type {
		case SMOG:
			dev_name = "Smog"
		case CO:
			dev_name = "CO"
		case CH4:
			dev_name = "CH4"
		default:
			log.Println("传感器SN号输入错误")
			flag = false

		}

		if flag {
			break
		}
	}

	now := time.Now()
	Device := CreateDevice(dev_name, dev_type, sn, now.Format("2006年01月02日 15:04:05"))
	fmt.Println(" ")
	fmt.Println("please enter your device state\n\r")
	fmt.Println("0 Init   ")
	fmt.Println("1 Zigbee_net   ")
	fmt.Println("2 Bind_Active    ")
	fmt.Println("3 Bind_Sleep     ")
	fmt.Println("4 Alarm           ")
	fmt.Println("5 Alarm_Has_Clear")
	reader := bufio.NewReader(os.Stdin)
	sn, err = reader.ReadString('\n')
	if err != nil {
		log.Println("输入流转换异常", err)
	}
	sn = strings.TrimSpace(sn)
	Device.State, err = strconv.Atoi(sn)
	if err != nil {
		log.Println("输入状态错误", err)
	}
	dev_sta, _ := Get_Device_State(Device)
	fmt.Printf("create %s_01 succ\n type: %s\nsn: %s\nstate is %s\n ", Device.Name, Device.Name, Device.SN, dev_sta)
	fmt.Println(now.Format("2006年01月02日 15:04:05"))
	return
}

/*
******************************************************************
  * @brief     get Device state
  * @param
  * @ret
  * @author    TRoy
  * @date      2017/8/4 2:45
******************************************************************
*/
func Get_Device_State(dev *Device) (sta string, err bool) {
	switch dev.State {
	case Init:
		sta = "Init"
	case Zigbee_net:
		sta = "Zigbee_net"
	case Bind_Active:
		sta = "Bind_Active"
	case Bind_Sleep:
		sta = "Bind_Sleep"
	case Alarm:
		sta = "Alarm"
	case Alarm_Has_Clear:
		sta = "Alarm_Has_Clear"
	default:
		err = true
	}
	return
}

/*
******************************************************************
  * @brief     Device test process
  * @param
  * @ret
  * @author    TRoy
  * @date      2017/8/4 2:46
******************************************************************
*/
func Test_Process(dev *Device) {
	switch dev.State {
	case Init:

	case Zigbee_net:

	case Bind_Active:

	case Bind_Sleep:

	case Alarm:

	case Alarm_Has_Clear:

	default:
	}
}

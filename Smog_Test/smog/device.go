package smog

const (
	SMOG = 1
	CO   = 2
	CH4  = 4
)
const gateway_SN = "00H10904160006"

const (
	Init            = iota
	Zigbee_net      //zigbee已入网
	Bind_Active     //已绑定在平台        开始有心跳
	Bind_Sleep      //已绑定在平台，睡眠  每次唤醒有心跳
	Alarm           //警报状态 未消警
	Alarm_Has_Clear //警报状态 已消警  此处检测环境如果有灾情不再报88 而是在心跳中上报83

)

type Device struct {
	Name       string
	Type       int
	SN         string
	CreateTime string
	State      int
}

type Gateway struct {
	SN string
}

func CreateDevice(name string, dev_type int, sn string, createTime string) (d *Device) {

	d = &Device{
		Name:       name,
		Type:       dev_type,
		SN:         sn,
		CreateTime: createTime,
	}
	return
}

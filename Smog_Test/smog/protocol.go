package smog

type Test_buff struct {
	head uint8
	len  uint8
	des  string `00000000000000`
	src  string `00000000000000`
	cmd  uint8
	msg  []uint8
	xor  uint8
	end  uint8
}

/*
******************************************************************
  * @brief    07 Search
  * @param
  * @ret
  * @author    TRoy
  * @date      2017/8/4 2:45
     单纯串口先发数据而不是阻塞等待 收到数据不解析直接显示在log中
******************************************************************
*/
func Search(gtw *Gateway) (b *Test_buff) {
	// var buff string
	b = &Test_buff{
		head: 0xff,
		len:  0x21,
		src:  gtw.SN,
		xor:  0x11,
	}

	return
}

/*
******************************************************************
  * @brief    01  Dev_Bind
  * @param
  * @ret
  * @author    TRoy
  * @date      2017/8/5 11:13
******************************************************************
*/
func Dev_Bind(dev *Device, sleep_sta bool) (sta int, err bool) {

	return
}

/*
******************************************************************
  * @brief  02  Dev_Unbind
  * @param
  * @ret
  * @author    TRoy
  * @date      2017/8/5 11:15
******************************************************************
*/
func Dev_Unbind(dev *Device) (sta int, err bool) {

	return
}

/*
******************************************************************
  * @brief   03  HeartBeat
  * @param
  * @ret
  * @author    TRoy
  * @date      2017/8/5 11:17
******************************************************************
*/
func GetHeartBeat(dev *Device) {

}

/*
******************************************************************
  * @brief    04 Val_Set
  * @param
  * @ret
  * @author    TRoy
  * @date      2017/8/5 11:19
******************************************************************
*/
func Val_Set(dev *Device) {
	return
}

/*
******************************************************************
  * @brief      05 Val_Get
  * @param
  * @ret
  * @author    TRoy
  * @date      2017/8/5 11:21
******************************************************************
*/
func Val_Get(dev *Device) {
	return
}

/*
******************************************************************
  * @brief  06 Alarm_Clear
  * @param
  * @ret
  * @author    TRoy
  * @date      2017/8/5 23:02
******************************************************************
*/
func Alarm_Clear(d *Device, sleep_sta bool) {

}

/*
******************************************************************
  * @brief   Alarm_Upload
  * @param
  * @ret
  * @author    TRoy
  * @date      2017/8/5 23:04
******************************************************************
*/
func Alarm_Upload() {

}

/*
******************************************************************
  * @brief   Get_Xor
  * @param
  * @ret
  * @author    TRoy
  * @date      2017/8/5 11:47
******************************************************************
*/
func Get_Xor(in []uint8, len uint8) (x uint8, err bool) {
	var (
		i uint8
	)
	if len > 0 {
		for i = 0; i < len; i++ {
			x ^= in[i]
		}
		return
	} else {
		err = true
		return
	}
}

package balance

import (
	"errors"
)

func init() {

}

var (
	ErrNoInstance       = errors.New("No Instance")
	ErrOutRangeInstance = errors.New("Instance out range")
)

type PollBalance struct {
	curIndex int //当前实例序号
}

func (p *PollBalance) DoBalance(insts []*Instance) (inst *Instance, err error) {
	if len(insts) == 0 {
		err = ErrNoInstance
		return
	}
	lens := len(insts)
	if p.curIndex > lens {
		err = ErrOutRangeInstance
		return
	}
	inst = insts[p.curIndex]
	p.curIndex = (p.curIndex + 1) % lens //从头再来 防越界
	return
}

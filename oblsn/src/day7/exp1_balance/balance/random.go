package balance

import (
	"math/rand"
)

type RandomBalance struct {
}

func (p *RandomBalance) DoBalance(insts []*Instance) (inst *Instance, err error) {
	if len(insts) == 0 {
		err = ErrNoInstance
		return
	}

	lens := len(insts) //获取实例数量
	index := rand.Intn(lens)
	inst = insts[index] //随机取其中一个
	return
}

package balance

import "fmt"

/*算法管理*/
//算法名-对应接口
type BalanceMgr struct {
	allBalancer map[string]Balancer
}

//map初始化
var mgr = BalanceMgr{
	make(map[string]Balancer),
}

//内部注册
func (p *BalanceMgr) registerBalancer(name string, b Balancer) {
	p.allBalancer[name] = b
}

//给外部调用的API
func RegisterBalancer(name string, b Balancer) {
	mgr.registerBalancer(name, b)
}

//通过输入指定负载均衡名称来调用相应算法
func DoBalance(name string, insts []*Instance) (inst *Instance, err error) {
	balancer, ok := mgr.allBalancer[name]
	if !ok {
		err = fmt.Errorf("Not found %s balancer", name)
	}
	fmt.Printf("use %s balancer\n", name)
	inst, err = balancer.DoBalance(insts)
	return
}

package balance

type Balancer interface {
	DoBalacne([]*Instance) *Instance
}

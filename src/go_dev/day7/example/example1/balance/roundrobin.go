package balance

import (
	"errors"
)

type RoundRobinBalance struct {
	curIndex int 
}

func (p *RoundRobinBalance) DoBalance(insts []*Instance) (inst *Instance, err error) {
	if len(insts) == 0 {
		err = errors.New("No instance")
		return
	}

	lens := len(insts)
	if p.curIndex >= lens {
		p.curIndex = 0 
	}

	inst = insts[p.curIndex]
	p.curIndex = (p.curIndex + 1) % lens 
	return
}

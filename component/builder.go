package ping

import "github.com/sarchlab/akita/v4/sim"

// type Builder struct {
// 	Engine sim.Engine
// 	Freq   sim.Freq
// }

func MakeBuilder() Builder {
	return Builder{}
}

func (b Builder) WithEngine(engine sim.Engine) Builder {
	b.Engine = engine
	return b
}

func (b Builder) WithFreq(freq sim.Freq) Builder {
	b.Freq = freq
	return b
}

func (b Builder) Build(name string) *Comp {
	pingAgent := &Comp{}
	pingAgent.ComponentBase = sim.NewComponentBase(name)
	pingAgent.OutPort = sim.NewPort(pingAgent, 4, 4, name+".OutPort")
	pingAgent.Engine = b.Engine

	return pingAgent
}

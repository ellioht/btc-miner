package cpu

import (
	"sync"
	"sync/atomic"
)

type Miner struct {
	sync.Mutex
	started atomic.Bool
	mining  atomic.Bool
}

func NewCpuMiner() *Miner {
	return &Miner{}
}

func (m *Miner) Start() {
	m.Lock()
	defer m.Unlock()

	m.started.Store(true)
}

func (m *Miner) Stop() {
	m.Lock()
	defer m.Unlock()

	m.started.Store(false)
}

func (m *Miner) IsStarted() bool {
	return m.started.Load()
}

func (m *Miner) IsMining() bool {
	return m.mining.Load()
}

func (m *Miner) mine() {
	m.mining.Store(true)
}

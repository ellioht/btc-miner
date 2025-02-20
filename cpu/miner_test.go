package cpu

import "testing"

func TestMiner(t *testing.T) {
	miner := NewCpuMiner()
	if miner.IsStarted() {
		t.Errorf("expected false, got true")
	}

	if miner.IsMining() {
		t.Errorf("expected false, got true")
	}

	miner.Start()

	if !miner.IsStarted() {
		t.Errorf("expected true, got false")
	}

	miner.Stop()

	if miner.IsStarted() {
		t.Errorf("expected false, got true")
	}
}

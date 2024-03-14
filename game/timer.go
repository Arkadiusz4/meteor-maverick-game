package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"time"
)

type Timer struct {
	CurrentTicks int
	TargetTicks  int
}

func NewTimer(d time.Duration) *Timer {
	return &Timer{
		CurrentTicks: 0,
		TargetTicks:  int(d.Milliseconds()) * ebiten.TPS() / 1000,
	}
}

func (t *Timer) Update() {
	if t.CurrentTicks < t.TargetTicks {
		t.CurrentTicks++
	}
}

func (t *Timer) IsReady() bool {
	return t.CurrentTicks >= t.TargetTicks
}

func (t *Timer) Reset() {
	t.CurrentTicks = 0
}

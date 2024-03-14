package tests

import (
	"github.com/Arkadiusz4/meteor-maverick-game/game"
	"github.com/hajimehoshi/ebiten/v2"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewTimer(t *testing.T) {
	duration := time.Second
	timer := game.NewTimer(duration)
	assert.NotNil(t, timer)
	assert.Equal(t, int(duration.Milliseconds())*ebiten.TPS()/1000, timer.TargetTicks)
}

func TestTimerUpdate(t *testing.T) {
	timer := game.Timer{CurrentTicks: 0, TargetTicks: 10}

	timer.Update()
	assert.Equal(t, 1, timer.CurrentTicks)

	for i := 0; i < 10; i++ {
		timer.Update()
	}
	assert.Equal(t, 10, timer.CurrentTicks)

	for i := 0; i < 10; i++ {
		timer.Update()
	}
	assert.Equal(t, 10, timer.CurrentTicks)
}

func TestTimerIsReady(t *testing.T) {
	timer := game.Timer{CurrentTicks: 5, TargetTicks: 10}
	assert.False(t, timer.IsReady())

	timer.CurrentTicks = 10
	assert.True(t, timer.IsReady())

	timer.CurrentTicks = 15
	assert.True(t, timer.IsReady())
}

func TestTimerReset(t *testing.T) {
	timer := game.Timer{CurrentTicks: 10, TargetTicks: 20}
	timer.Reset()
	assert.Equal(t, 0, timer.CurrentTicks)
}

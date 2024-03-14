package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewTimer(t *testing.T) {
	duration := time.Second
	timer := NewTimer(duration)
	assert.NotNil(t, timer)
	assert.Equal(t, int(duration.Milliseconds())*ebiten.TPS()/1000, timer.targetTicks)
}

func TestTimerUpdate(t *testing.T) {
	timer := Timer{currentTicks: 0, targetTicks: 10}

	timer.Update()
	assert.Equal(t, 1, timer.currentTicks)

	for i := 0; i < 10; i++ {
		timer.Update()
	}
	assert.Equal(t, 10, timer.currentTicks)

	for i := 0; i < 10; i++ {
		timer.Update()
	}
	assert.Equal(t, 10, timer.currentTicks)
}

func TestTimerIsReady(t *testing.T) {
	timer := Timer{currentTicks: 5, targetTicks: 10}
	assert.False(t, timer.IsReady())

	timer.currentTicks = 10
	assert.True(t, timer.IsReady())

	timer.currentTicks = 15
	assert.True(t, timer.IsReady())
}

func TestTimerReset(t *testing.T) {
	timer := Timer{currentTicks: 10, targetTicks: 20}
	timer.Reset()
	assert.Equal(t, 0, timer.currentTicks)
}

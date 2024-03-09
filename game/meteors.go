package game

import (
	"github.com/Arkadiusz4/meteor-maverick-game/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"math/rand"
)

type Meteor struct {
	position Vector
	sprite   *ebiten.Image
}

var meteorSprites = assets.MeteorSprites

func NewMeteor() *Meteor {
	sprite := *meteorSprites[rand.Intn(len(meteorSprites))]
	return &Meteor{
		position: Vector{},
		sprite:   *sprite, // Dereference the pointer here
	}
}

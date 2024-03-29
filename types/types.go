package types

import (
	"fmt"
)

type Coord struct{ X, Y int }

func C(X, Y int) Coord {
	return Coord{X, Y}
}

func (c Coord) Add(d Coord) Coord {
	return Coord{
		X: c.X + d.X,
		Y: c.Y + d.Y,
	}
}

func (c Coord) String() string {
	return fmt.Sprintf("%d,%d", c.X, c.Y)
}

type Position struct {
	Coord
	Theta int
}

func Pos(X, Y, Theta int) Position {
	return Position{Coord{X, Y}, Theta}
}

func (p Position) Add(c Coord) Position {
	return Pos(p.X+c.X, p.Y+c.Y, p.Theta)
}

type Tile byte

const (
	Invalid Tile = iota + 1
	Water
	Grass
	GrassUp = 12 + Water
)

func NewWorld(ts []Tile, width, height int) World {
	return World{
		ts, width, height,
	}
}

type World struct {
	tiles         []Tile
	Width, Height int
}

func (w World) TileBytes() []byte {
	var bs []byte
	for _, t := range w.tiles {
		bs = append(bs, byte(t))
	}
	return bs
}

func (w World) ValidTarget(t Position) bool {
	if t.X < 0 || t.X >= w.Width || t.Y < 0 || t.Y >= w.Height {
		return false
	}
	return w.tiles[t.Y*w.Width+t.X] == Grass
}

func (w World) At(p Coord) Tile {
	if p.X < 0 || p.X >= w.Width || p.Y < 0 || p.Y >= w.Height {
		return Water
	}
	return w.tiles[p.Y*w.Width+p.X]
}

func (w World) Set(p Coord, t Tile) {
	if p.X < 0 || p.X >= w.Width || p.Y < 0 || p.Y >= w.Height {
		return
	}

	w.tiles[p.Y*w.Width+p.X] = t
}

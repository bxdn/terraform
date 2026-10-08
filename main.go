package main

import (
	"time"

	"github.com/bxdn/flite/server"
)

type Ore int

const (
	BAUXITE Ore = iota
	HEMATITE
	ICE
	AMONIA
	LIMESTONE
	QUARTZ
)

type BuildingType int

const (
	POWER_BANK = iota
	FABRICATOR
	MINER
	SEPARATOR
	SILO
	CHARGING_STATION
	SCANNER
)

type ActivityType int

const (
	CONSTRUCTING = iota
	MOVING
	IDLE
	LOADING
	UNLOADING
	CHARGING
)

type Point struct {
	X, Y int
}

type Building struct {
	Coords       Point
	BuildingType BuildingType
}

type Patch struct {
	Coords  Point
	Ore     Ore
	Density float64
}

type Agent struct {
	Coords         Point
	ActivityType   ActivityType
	TimeSpentDoing time.Duration
}

type Game struct {
	Patches   []Patch
	Buildings map[BuildingType]*Building
	Agents    []*Agent
}

type Locatable interface {
	GetLocation() Point
}

type Ticker interface {
	Tick(game *Game)
}

func main() {
	server.GET("/hello", func(f *server.F[server.No]) error {
		return f.ReturnTextOk("Hi!")
	})
	server.Serve(8000)
}

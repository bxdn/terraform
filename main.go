package main

import (
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
)

type ActivityType int

const (
	CONSTRUCTING = iota
	MOVING
	IDLE
	LOADING
	UNLOADING
)

type Point struct {
	X, Y int
}

type Building struct {
	coords       Point
	buildingType BuildingType
}

type Patch struct {
	coords  Point
	ore     Ore
	density float64
}

type Agent struct {
	coords Point
}

type Game struct {
	Patches   []Patch
	Buildings []Building
	Agents    []Agent
}

func main() {
	server.GET("/hello", func(f *server.F[server.No]) error {
		return f.ReturnTextOk("Hi!")
	})
	server.Serve(8000)
}

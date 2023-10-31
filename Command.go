package main

import (
	"fmt"
)

type LightOutside struct {
	intensity float64
}

func (light *LightOutside) switchOn() {
	fmt.Println("Light's switched on")
}

func (light *LightOutside) switchOff() {
	fmt.Println("Light's switched off")
}

type HeatingCooling struct {
	temperature float64
}

func (hc *HeatingCooling) start() {
	mode := "cooling"
	if hc.temperature >= 25 {
		mode = "heating"
	}
	fmt.Printf("Start %s\n", mode)
}

func (hc *HeatingCooling) stop() {
	mode := "cooling"
	if hc.temperature >= 25 {
		mode = "heating"
	}
	fmt.Printf("Stop %s\n", mode)
}

type Command interface {
	execute()
}

type SwitchOnLightCommand struct {
	light *LightOutside
}

func (cmd *SwitchOnLightCommand) execute() {
	cmd.light.switchOn()
}

type SwitchOffLightCommand struct {
	light *LightOutside
}

func (cmd *SwitchOffLightCommand) execute() {
	cmd.light.switchOff()
}

type StartHeatingCommand struct {
	heater *HeatingCooling
}

func (cmd *StartHeatingCommand) execute() {
	if cmd.heater.temperature < 25 {
		cmd.heater.temperature = 25
	}
	cmd.heater.start()
}

type StartCoolingCommand struct {
	cooler *HeatingCooling
}

func (cmd *StartCoolingCommand) execute() {
	if cmd.cooler.temperature > 25 {
		cmd.cooler.temperature = 26
	}
	cmd.cooler.start()
}

type Programm struct {
	commands []Command
}

func (p *Programm) start() {
	for _, cmd := range p.commands {
		cmd.execute()
	}
}

func main() {
	light := &LightOutside{}
	heater := &HeatingCooling{}
	cooler := &HeatingCooling{}

	lightOnCommand := &SwitchOnLightCommand{light: light}
	lightOffCommand := &SwitchOffLightCommand{light: light}
	heatCommand := &StartHeatingCommand{heater: heater}
	coolCommand := &StartCoolingCommand{cooler: cooler}

	eveningProgram := &Programm{}
	eveningProgram.commands = append(eveningProgram.commands, lightOnCommand, lightOffCommand, heatCommand, coolCommand)

	eveningProgram.start()
}

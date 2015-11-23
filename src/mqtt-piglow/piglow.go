package main

import (
	"fmt"
	"github.com/schoentoon/piglow"
	"math"
	"strings"
)

func GlowOff() {
	fmt.Println("Going to turn off the lights")
		
	piglow.ShutDown()
}

func checkPiGlow() bool {
	if !piglow.HasPiGlow() {
		fmt.Println("piglow is not connected")
		return false
	}
	return true
}

func glowToColor(color byte, brightness float64) error {
	brightnessByte := byte(math.Floor(brightness*255.0 + 0.5))
	return piglow.Ring(color, brightnessByte)

}

func setLegOn(leg int, brightness float64) error {
	return piglow.Leg(byte(leg), byte(brightness*255+0.5))

}

func setLedOn(leg int, color string, brightness float64) error {
	return piglow.Led(byte(leg), getColorFromString(color), byte(brightness*255+0.5))
}

func TurnAllOn() {
	fmt.Println("Going to turn on the lights")

	for i := 0; i < 3; i++ {
		setLegOn(i, 0.3)
	}
}

func Flare() {
	for i := 0; i < 3; i++ {
		setLegOn(i, 1)
	}
}

func getColorFromString(color string) byte {
	if strings.Contains(color, "red") {
		return piglow.Red
	}
	if strings.Contains(color, "orange") {
		return piglow.Orange
	}
	if strings.Contains(color, "yellow") {
		return piglow.Yellow
	}
	if strings.Contains(color, "green") {
		return piglow.Green
	}
	if strings.Contains(color, "blue") {
		return piglow.Blue
	}
	if strings.Contains(color, "white") {
		return piglow.White
	}
	return 0x00
}

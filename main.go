package main

import (
	"fmt"
	"math"
)

func calcDripInfusion() {
	var amountInfusion, dripTotalTime float64
	var tubeForAdult bool
	var dripPerMinutes float64

	fmt.Print("Amount of Infusion(ml) : ")
	fmt.Scan(&amountInfusion)

	fmt.Print("Tube for ADULT or CHILD (if ADULT type 1, CHILD type 0) : ")
	fmt.Scan(&tubeForAdult)

	fmt.Print("How long hours the infusion needs? :")
	fmt.Scan(&dripTotalTime)

	if tubeForAdult {
		dripPerMinutes = amountInfusion * 20 / (dripTotalTime * 60)
	} else {
		dripPerMinutes = amountInfusion * 60 / (dripTotalTime * 60)
	}

	fmt.Println(math.Round(dripPerMinutes*10) / 10)

}

func main() {
	calcDripInfusion()
}

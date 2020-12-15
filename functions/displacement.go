package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func GenDisplaceFn(acceleration, velocity, displacement float64) func(float64) float64 {

	/* calculates the displacement of a body given the acceleration, velocity and initial
	displacement */

	fn := func(time float64) float64 {

		/* calculates the displacement of a body given the time */

		acc_time := (acceleration * math.Pow(time, 2)) / 2
		vel_time := (velocity * time)

		return acc_time + vel_time + displacement
	}
	return fn
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {

		// prompt users to enter values
		fmt.Println("enter values for acceleration, initial velocity, initial displacement and enter 'stop intial' to stop")
		scanner.Scan()
		acceleration, _ := strconv.ParseFloat(scanner.Text(), 64)
		scanner.Scan()
		velocity, _ := strconv.ParseFloat(scanner.Text(), 64)
		scanner.Scan()
		displacement, _ := strconv.ParseFloat(scanner.Text(), 64)

		fn := GenDisplaceFn(acceleration, velocity, displacement)

		for {
			fmt.Println("enter values for time and enter 'stop time' to stop")
			scanner.Scan()
			time, _ := strconv.ParseFloat(scanner.Text(), 64)
			final_displacement := fn(time)
			fmt.Println(final_displacement)

			if scanner.Text() == "stop time" {
				break
			} else {

			}

		}

		if scanner.Text() == "stop initial" {
			break
		}

	}

}

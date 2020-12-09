package main

import "fmt"
import "os"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	// nbFloors: number of floors
	// width: width of the area
	// nbRounds: maximum number of rounds
	// exitFloor: floor on which the exit is found
	// exitPos: position of the exit on its floor
	// nbTotalClones: number of generated clones
	// nbAdditionalElevators: ignore (always zero)
	// nbElevators: number of elevators
	var nbFloors, width, nbRounds, exitFloor, exitPos, nbTotalClones, nbAdditionalElevators, nbElevators int
	fmt.Scan(&nbFloors, &width, &nbRounds, &exitFloor, &exitPos, &nbTotalClones, &nbAdditionalElevators, &nbElevators)

	map_elevators := make(map[int]int)

	for i := 0; i < nbElevators; i++ {
		// elevatorFloor: floor on which this elevator is found
		// elevatorPos: position of the elevator on its floor
		var elevatorFloor, elevatorPos int
		fmt.Scan(&elevatorFloor, &elevatorPos)
		map_elevators[elevatorFloor] = elevatorPos
	}

	for {
		// cloneFloor: floor of the leading clone
		// clonePos: position of the leading clone on its floor
		// direction: direction of the leading clone: LEFT or RIGHT
		var cloneFloor, clonePos int
		var direction string
		fmt.Scan(&cloneFloor, &clonePos, &direction)

		/*
		   fmt.Fprintln(os.Stderr, "nbFloors:", nbFloors, "width:", width, "nbRounds:", nbRounds, "exitFloor:", exitFloor, "exitPos:", exitPos)
		   fmt.Fprintln(os.Stderr, "nbTotalClones:", nbTotalClones, "nbAdditionalElevators:", nbAdditionalElevators, "nbElevators:", nbElevators)
		   fmt.Fprintln(os.Stderr, "cloneFloor:", cloneFloor, "clonePos:", clonePos, "direction:", direction)
		   fmt.Fprintln(os.Stderr, "map_elevators:", map_elevators)
		*/

		a := needBlock(direction, cloneFloor, clonePos, exitFloor, exitPos, map_elevators)
		//fmt.Fprintln(os.Stderr, "a:", a)
		if a {
			fmt.Println("BLOCK")
		} else {
			fmt.Println("WAIT") // action: WAIT or BLOCK
		}

		// fmt.Fprintln(os.Stderr, "Debug messages...")
		//fmt.Println("WAIT") // action: WAIT or BLOCK
	}
}

func needBlock(direction string, cloneFloor, clonePos, exitFloor, exitPos int, map_elevators map[int]int) (answer bool) {
	target := -1
	if cloneFloor == exitFloor {
		target = exitPos
	} else {
		target = map_elevators[cloneFloor]
	}

	//fmt.Fprintln(os.Stderr, "target:", target)

	if clonePos < target {
		switch direction {
		case "LEFT":
			answer = true
		case "RIGHT":
			answer = false
		default:
			fmt.Fprintln(os.Stderr, "needBlock default case: Issue !!!")
			answer = false
		}
	} else if clonePos > target {
		switch direction {
		case "LEFT":
			answer = false
		case "RIGHT":
			answer = true
		default:
			fmt.Fprintln(os.Stderr, "needBlock default case: Issue !!!")
			answer = false
		}
	} else {
		fmt.Fprintln(os.Stderr, "needBlock, clonePos == target: Do nothing !!!")
		answer = false
	}
	return
}

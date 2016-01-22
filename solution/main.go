package main

import "fmt"

func weeksBeforeSustenance(mouthsToFeed int, fruitProvided int) int {
	harvest := plant(fruitProvided)
	weeks := 1

	for yield := <-harvest; yield < mouthsToFeed; yield = <-harvest {
		weeks++
	}
	return weeks
}

func plant(fruitProvided int) <-chan int {
	harvests := make(chan int)

	go func() {
		numberOfPlants := fruitProvided
		numberOfFruit := 0

		for {
			harvests <- numberOfFruit
			numberOfFruit += numberOfPlants
			numberOfPlants += numberOfFruit
		}
	}()
	return harvests
}

func main() {
	fmt.Println(weeksBeforeSustenance(200, 15))
	fmt.Println(weeksBeforeSustenance(50000, 1))
	fmt.Println(weeksBeforeSustenance(150000, 250))
}

// Source: https://www.reddit.com/r/dailyprogrammer/comments/3twuwf/20151123_challenge_242_easy_funny_plant/cywdv4h

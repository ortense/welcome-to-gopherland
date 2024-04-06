package main

import "fmt"

func main() {
	eveelutions := map[int]string{
		134: "vaporeon",
		135: "jolteon",
		136: "flareon",
		196: "espeon",
		197: "umbreon",
		470: "leafeon",
		471: "glaceon",
	}

	eveelutions[700] = "sylveon"

	name, exist := eveelutions[197]

	if exist {
		fmt.Println(name)
	}

	fmt.Println(eveelutions, len(eveelutions))
}

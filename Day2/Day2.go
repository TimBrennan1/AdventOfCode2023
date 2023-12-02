package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//Only primitive types can be constants
var reNum = regexp.MustCompile(`[^0-9]+`)
var reString = regexp.MustCompile(`[^A-Za-z]+`)
const part = 2;

func main() {
	t0 := time.Now();
	sum := ParseDocument()
	t1 := time.Now();

	fmt.Printf("Sum: %v Duration: %v", sum, t1.Sub(t0) );
}

//parse per line
//parse per semi colon
//check the values based on a dictionary
//check the number based on the dictionary and number of items vs expected


func ParseDocument() (sum int){

	file, _ := os.Open("Day2.txt")
	defer file.Close()

	s := bufio.NewScanner(file)

	for s.Scan() {
		gameNumber, ok := LineInfo(s.Text())

		if ok{
			sum += gameNumber
		}
	}
	return sum
}

func LineInfo(in string) (gameNumber int, ok bool) {
	var maxBlue, maxRed, maxGreen int = 0,0,0

	truncateGame := strings.Split(in, ":")
	gameNumber, _ = strconv.Atoi(reNum.ReplaceAllString(truncateGame[0], ""));

	//gets the round
	splitByRound := strings.Split(truncateGame[1], ";")

	for _, v := range splitByRound {
		//Split on the type
		splitByType := strings.Split(v, ",")
		for _, val := range splitByType {
			//Check to see if all the numbers fit the criteria
			if part == 1{
				valid := ValidityChecker(val)
				if !valid{
					return 0, false
				}
			} 
			//calculate power
			blue, green, red := getVals(val);
			if blue > maxBlue{
				maxBlue = blue;
			}
			if green > maxGreen{
				maxGreen = green
			}
			if red> maxRed{
				maxRed = red
			}
			
		}
	}
	if part == 1{
	return gameNumber, true
	} 
		if maxBlue <= 0 {
			maxBlue = 1
		}
		if maxRed <= 0 {
			maxRed = 1
		}
		if maxGreen <= 0{
			maxGreen = 1
		}
	return maxBlue*maxGreen*maxRed, true
}

func getVals(in string) (blue int, green int, red int ){
	//get type
	cubeType := reString.ReplaceAllString(in,"")
	//get just number
	cubeNumber, _ := strconv.Atoi(reNum.ReplaceAllString(in,""))

	if cubeType == "blue"{
		blue = cubeNumber
	} else if cubeType == "red"{
		red = cubeNumber
	} else {
		green = cubeNumber
	}

	return blue, red, green
}

func ValidityChecker(in string) (valid bool){
	bag := map[string]int{
		"blue": 14,
		"red" : 12,
		"green" : 13,
	}

	//Prolly would of been better to split on the space
	// get just text
	cubeType := reString.ReplaceAllString(in,"")
	//get just number
	cubeNumber, _ := strconv.Atoi(reNum.ReplaceAllString(in,""))
	return cubeNumber <= bag[cubeType]
}

//Part 1 Answer: 2265
//Part 2 Answer: 64097
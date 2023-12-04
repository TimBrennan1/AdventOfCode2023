package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"
)

func main() {
	t0 := time.Now();
	sum := Process()
	t1 := time.Now()

	fmt.Println("Sum: ", sum, "Duration: ", t1.Sub(t0))
}

func Process() (value int){
	file,_ := os.Open("Day4.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var reNum = regexp.MustCompile(`[^0-9]+`)
	gameMap := make(map[int]int)

	for scanner.Scan(){
		text:= scanner.Text();
		card := strings.Split(text, ":")
		scores := card[1]
		numbers := strings.Split(scores, "|")

		winners:= numbers[0]
		ourNums:= numbers[1]

		winnerTokens := RemoveWhitespace(strings.Split(winners, " "))
		ourTokens := RemoveWhitespace(strings.Split(ourNums, " "))

		num, _ := strconv.Atoi(reNum.ReplaceAllString(card[0], ""))
		//value += eval(ourTokens, winnerTokens)
		Incr(gameMap, num)

		times := gameMap[num]
		for i := 0; i < times; i++{
			eval2(ourTokens, winnerTokens, num, gameMap)
		}
		

	}
	
	for i := 0; i <= len(gameMap); i++ {
		value += gameMap[i]
	}
	return value
}

func RemoveWhitespace(in []string) []string{
	re := regexp.MustCompile(`\s+`)
	newSlice := []string{}
	for _, v := range in {
		resultString := re.ReplaceAllString(v, "")
		if resultString != ""{
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}

func eval(ourTokens []string, winnerTokens []string) (count int){

	for _, v := range winnerTokens {

	
		ok := slices.Contains(ourTokens, v) 

		if ok {
			if count == 0{
				count = 1
			} else{
				count *= 2
			}

		}
	}
	return count

}

func eval2(ourTokens []string, winnerTokens []string, gameNum int, gameMap map[int] int,){

		for _, v := range winnerTokens {
		
			ok := slices.Contains(ourTokens, v) 

			if ok {
				Incr(gameMap, gameNum+1)
				gameNum++
			}
		}
	

}

func Incr(counterMap map[int]int, gameNum int){
	counterMap[gameNum] += 1

}

//Part1: 25010
//Part2: 9924412
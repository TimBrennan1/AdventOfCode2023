package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)
	func main(){
		t0 := time.Now()
		out, err := NumerizeText()
		t1 := time.Now();
		if err != nil{
			fmt.Print("Yikes")
		}

		fmt.Printf("Sum: %v Duration: %v", out, t1.Sub(t0) );
	}

	func openFile(in string) (file *os.File, err error){
		file, err = os.Open(in);
		if err != nil {
			return nil, err;
		}
		return file, nil;
	}

	func retriveNumericValue(in string) (int, error){

		if len(in) == 0 {
			return 0, nil
		}
		
		first := string(in[0])
		last := string(in[len(in)-1])
	
		v, err := strconv.Atoi(first + last)

		if err != nil {
			return 0, err
		}
		
		return v, nil;

	}

	func NumerizeText() (sum int, err error) {
		sum =0
		file, err := openFile("Day1.txt")
		defer file.Close()
	
		if err != nil {
			return 0, errors.New("Can't parse home-slice")
		}
	
		scanner := bufio.NewScanner(file)
	
		for scanner.Scan() {
			v := sanitizeInput(scanner.Text())
			num, err := retriveNumericValue(v)
	
			if err != nil {
				return 0, err
			}
	
			sum += num
		}
	
		return sum, nil
	}

	func sanitizeInput(input string) string{

		vals := map[string]string{
			"one" : "o1e",
			"two" : "t2o",
			"three" : "t3e",
			"four" : "f4r",
			"five" : "f5e",
			"six" : "s6x",
			"seven" : "s7n",
			"eight" : "e8t",
			"nine" : "n9e",
		}

		for word, digit := range vals {
			input = strings.ReplaceAll(input, word, digit)
		}

		re := regexp.MustCompile(`[^0-9]+`)
   		input = re.ReplaceAllString(input, "")

		return input;
		
	}
	
	//Part1 Answer: 56108
	//Part2 Answer: 55652
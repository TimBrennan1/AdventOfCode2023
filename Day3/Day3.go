package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"time"
	"unicode"
)

func main() {
	t0 := time.Now();
	//sum := Process();
	sum := Process2()
	t1 := time.Now()

	fmt.Print("Sum: ", sum, " Duration ", t1.Sub(t0))

}


func ToMatrix() (matrix [][]rune, fileStringified string){
	file, err := os.Open("Day3.txt")
	defer file.Close()

	if err != nil{
		return
	}
	scanner := bufio.NewScanner(file)
	index := 0
	matrix = [][]rune{}
	for scanner.Scan() {
		text := scanner.Text()
		slice := []rune{}

		for j := 0; j < len(text); j++{
			slice = append(slice,rune(text[j]))
		}
		matrix = append(matrix,slice)
		index++
	}
	v, _ := os.ReadFile("Day3.txt")
	return matrix, string(v)

}


func removeDuplicates(input []string) (result []string) {
	uniqueElements := make(map[string]bool)

	for _, value := range input {
		if !uniqueElements[value] {
			uniqueElements[value] = true
			result = append(result, value)
		}
	}

	return result
}

func GetValidSymbols(in string) (symbols []string){
	re := regexp.MustCompile(`[^\s\w.]+`)
	symbols = re.FindAllString(in, -1)
	symbols = removeDuplicates(symbols)

	return symbols

}

//find a number, check if there is a symbol nearby
func CalcSum (matrix [][]rune, validSymbols []string) (sum int){

	matrixHeight:= len(matrix[0])
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < matrixHeight; j++ {
			val, skip := GetValue(i,j,matrix,validSymbols)
			sum += val
			j += skip

		}
	}
	return sum
}

func GetValue(i int, j int, matrix [][]rune, validSymbols []string)(value int, skip int){
	v := ""
	index := j
	hasSymbol := false

	for index < len(matrix[0]){
		ok := unicode.IsDigit(matrix[i][index])
		if !ok{
			break
		}
		v += string(matrix[i][index])
		valid := CheckNeighbors(i,index,matrix,validSymbols)
		if valid {
			hasSymbol = true
		}
		index++
	}

	if hasSymbol{
		value, _ = strconv.Atoi(v)
	}

	return value, index-j
}

//prolly should of made seperate function for similar code below
func CheckNeighbors(i int, j int, matrix [][]rune, validSymbols []string)(ok bool){
	topBound := 0
	leftBound := 0
	bottomBound := len(matrix) -1
	rightBound := len(matrix[0]) -1
	valid := false

	//check top
	if i-1 >= topBound{
		//check above
		valid = slices.Contains(validSymbols, string(matrix[i-1][j]))
		if valid {
			return true
		}
		//check top-left
		if j-1 >= leftBound{
			valid = slices.Contains(validSymbols, string(matrix[i-1][j-1]))
			if valid {
				return true
			}
		}
		//check top-right
		if j+1 <= rightBound{
			valid = slices.Contains(validSymbols, string(matrix[i-1][j+1]))
			if valid {
				return true
			}
		}
	}
	//check bottom
	if i+1 <= bottomBound{
		//check right
		valid = slices.Contains(validSymbols, string(matrix[i+1][j]))
		if valid {
			return true
		}
		// check bottom-left
		if j-1 >= leftBound{
			valid = slices.Contains(validSymbols, string(matrix[i+1][j-1]))
			if valid {
				return true
			}
		}
		//check left down
		if j+1 <= rightBound{
			valid = slices.Contains(validSymbols, string(matrix[i+1][j+1]))
			if valid {
				return true
			}
		}
	}
	//check left
	if j-1 >= leftBound{
		valid = slices.Contains(validSymbols, string(matrix[i][j-1]))
		if valid {
			return true
		}
	}
	//check right
	if j+1 <= rightBound{
		valid = slices.Contains(validSymbols, string(matrix[i][j+1]))
			if valid {
				return true
			}

	}
	return false

}

func Process() (sum int){
	matrix, fileStringified := ToMatrix()
	if matrix == nil{
		return 0
	}
	symbols := GetValidSymbols(fileStringified)
	sum = CalcSum(matrix, symbols)
	return sum
}

func Process2() (value int){
	matrix, _ := ToMatrix()
	if matrix == nil{
		return 0
	}
	return GetValue2(matrix)
	//find by *
	//check neighbors, and get their values
}

func GetValue2(matrix [][]rune)(value int){
	nums := []string{"1","2","3","4","5","6","7","8","9","0"}
	matrixHeight:= len(matrix[0])
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < matrixHeight; j++ {
			if string(matrix[i][j]) == "*"{
				n1, n2 := CheckNeighbors2(i,j,matrix,nums)
				if n1 != 0 && n2 !=0{
					value += n1 * n2
				}
			}

		}
	}
	return value
}

func GetFullNumber(matrix [][]rune, i int, j int) (result string, iRange int, jRange []int){

	before := ""
	after := ""
	index := j + 1
	iRange = i

	// print("hi")

	for index < len(matrix[0]){
		val, err := strconv.Atoi(string(matrix[i][index]))
		if err != nil{
			break
		}
		after = fmt.Sprint(after,val)

		jRange = append(jRange, index)
		index++
	}

	index = j
	for index >= 0{
		val, err := strconv.Atoi(string(matrix[i][index]))
		if err != nil{
			break
		}
		before = fmt.Sprint(before,val)
		

		jRange = append(jRange, index)
		index--
	}
	before = reverseString(before)
	return before+after, iRange, jRange
}

func reverseString(s string) string {
    // Convert the string to a slice of runes
    runes := []rune(s)

    // Get the length of the slice
    length := len(runes)

    // Reverse the slice of runes
    for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }

    // Convert the slice of runes back to a string
    reversedString := string(runes)

    return reversedString
}


func CheckNeighbors2(i int, j int, matrix [][]rune, validSymbols []string)(n1 int, n2 int){
	topBound := 0
	leftBound := 0
	bottomBound := len(matrix) -1
	rightBound := len(matrix[0]) -1
	valid := false
	var iRange int = -1
	jRange := []int{}

	//check top
	if i-1 >= topBound{
		//check above
		valid = slices.Contains(validSymbols, string(matrix[i-1][j])) && !(i-1 == iRange && slices.Contains(jRange, j))
		if valid {
			result, r1, r2 := GetFullNumber(matrix, i-1,j)
			if result != ""{
				if n1 == 0{
					iRange = r1
					jRange = r2
					n1,_ = strconv.Atoi(result)
				} else if n2 == 0{
					n2,_ = strconv.Atoi(result)
				}
			}

		}
		//check top-left
		if j-1 >= leftBound{
			valid = slices.Contains(validSymbols, string(matrix[i-1][j-1])) && !(i-1 == iRange && slices.Contains(jRange, j-1))

			if valid {
				result, r1, r2 := GetFullNumber(matrix, i-1,j-1)
				if result != ""{
					if n1 == 0{
						iRange = r1
						jRange = r2
						n1,_ = strconv.Atoi(result)
					} else if n2 == 0{
						n2,_ = strconv.Atoi(result)
					}
				}
	
			}
		}
		//check top-right
		if j+1 <= rightBound{
			valid = slices.Contains(validSymbols, string(matrix[i-1][j+1])) && !(i-1 == iRange && slices.Contains(jRange, j+1))
			if valid {
				result, r1, r2 := GetFullNumber(matrix, i-1,j+1)

				if result != ""{
					if n1 == 0{
						iRange = r1
						jRange = r2
						n1,_ = strconv.Atoi(result)
					} else if n2 == 0{
						n2,_ = strconv.Atoi(result)
					}
				}
	
			}
		}
	}
	//check bottom
	if i+1 <= bottomBound{
		//check bottom
		valid = slices.Contains(validSymbols, string(matrix[i+1][j])) && !(i+1 == iRange && slices.Contains(jRange, j))
		if valid {
			result, r1, r2 := GetFullNumber(matrix, i+1,j)
			if result != ""{
				if n1 == 0{
					iRange = r1
					jRange = r2
					n1,_ = strconv.Atoi(result)
				} else if n2 == 0{
					n2,_ = strconv.Atoi(result)
				}
			}

		}
		// check bottom-left
		if j-1 >= leftBound{
			valid = slices.Contains(validSymbols, string(matrix[i+1][j-1])) && !(i+1 == iRange && slices.Contains(jRange, j-1))
			if valid {
				result, r1, r2 := GetFullNumber(matrix, i+1,j-1)
				if result != ""{
					if n1 == 0{
						iRange = r1
						jRange = r2
						n1,_ = strconv.Atoi(result)
					} else if n2 == 0{
						n2,_ = strconv.Atoi(result)
					}
				}
	
			}
		}
		//check bottom right
		if j+1 <= rightBound{
			valid = slices.Contains(validSymbols, string(matrix[i+1][j+1])) && !(i+1 == iRange && slices.Contains(jRange, j+1))

			if valid {
				result, r1, r2 := GetFullNumber(matrix, i+1,j+1)
				if result != ""{
					if n1 == 0{
						iRange = r1
						jRange = r2
						n1,_ = strconv.Atoi(result)
					} else if n2 == 0{
						n2,_ = strconv.Atoi(result)
					}
				}
	
			}
		}
	}
	//check left
	if j-1 >= leftBound{
		valid = slices.Contains(validSymbols, string(matrix[i][j-1])) && !(i == iRange && slices.Contains(jRange, j-1))
		if valid {
			result, r1, r2 := GetFullNumber(matrix, i,j-1)
			if result != ""{
				if n1 == 0{
					iRange = r1
					jRange = r2
					n1,_ = strconv.Atoi(result)
				} else if n2 == 0{
					n2,_ = strconv.Atoi(result)
				}
			}

		}
	}
	//check right
	if j+1 <= rightBound{
		valid = slices.Contains(validSymbols, string(matrix[i][j+1])) && !(i == iRange && slices.Contains(jRange, j+1))
		if valid {
			result, r1, r2 := GetFullNumber(matrix, i,j+1)
			if result != ""{
				if n1 == 0{
					iRange = r1
					jRange = r2
					n1,_ = strconv.Atoi(result)
				} else if n2 == 0{
					n2,_ = strconv.Atoi(result)
				}
			}

		}

	}
	//println(n1, n2)
	return n1, n2

}

//Part 1: 557705
//Part 2: 84266818





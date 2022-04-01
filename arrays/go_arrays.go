package arrays

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func ArrayMethods() []string {
	// slice_name := []datatype{values}
	mySlice := []int{1, 2, 3}
	mySliceStr := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println(mySlice)
	fmt.Println(mySliceStr)
	return mySliceStr

}

func AppendToSlice(array []string, toAppend string) []string {
	return append(array, toAppend)
}

func Maps() {
	aMap := map[string]int{
		"key1": -1,
		"key2": 123,
	}
	fmt.Println(aMap)
	_, ok := aMap["key3"]
	fmt.Println(ok)

	for _, v := range aMap {
		fmt.Println(" # ", v)
	}

}

func AddGenerics[T float64 | int | string](a, b T) {
	fmt.Println(a + b)
}

func PrintSlice[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

type Ordered interface {
	int | float32 | ~string | float64
}

func Same[T Ordered](a, b T) bool {
	return a == b
}

type Numeric interface {
	int | int8 | int16 | int32 | int64 | float64
}

func Add[T Numeric](a, b T) T {
	return a + b
}

func MaxNumberInList(list []int) int {
	maxNum := list[0]
	for _, num := range list {
		if num > maxNum {
			maxNum = num
		}
	}
	return maxNum
}

func DifferenceMaxMin(list []int) int {
	maxNum, minNum := list[0], list[0]
	for _, num := range list {
		if num > maxNum {
			maxNum = num
		}
		if num < minNum {
			minNum = num
		}
	}
	fmt.Println("max is ", maxNum, "min is ", minNum)
	return maxNum - minNum
}

func EvenOddTransform(array []int, n int) []int {
	for i := 0; i < n; i++ {
		for index, num := range array {
			if num%2 == 0 {
				array[index] = num - 2
			} else {
				array[index] = num + 2
			}
		}
	}
	return array
}

func LetterCounter(nestedArray [][]string, letter string) int {
	counter := 0
	for _, array := range nestedArray {
		for _, char := range array {
			if char == letter {
				counter += 1
			}
		}
	}
	return counter
}

func SpinAround(array []string) int {
	rightCounter, leftCounter := 0, 0
	for _, value := range array {
		if value == "right" {
			rightCounter += 1
		} else {
			leftCounter += 1
		}
	}
	return int(math.Ceil(float64((rightCounter - leftCounter) * 90 / 360)))
}

func SquarePatch(n int) [][]int {
	var (
		OuterSlice [][]int
		innerSlice []int
	)
	for i := 0; i < n; i++ {
		innerSlice = append(innerSlice, n)
	}
	for i := 0; i < n; i++ {
		OuterSlice = append(OuterSlice, innerSlice)
	}
	if n == 0 {
		return [][]int{}
	}
	return OuterSlice
}

func RemoveSpecialCharacters(str string) string {
	res := str
	for _, char := range "!@#$%^&\\*()./+,:;<=>?`|{}~[]" {
		res = strings.Replace(res, string(char), "", 5)
	}
	return res
}

func ArrayElementSum(array []int) []int {
	sum := 0
	var res []int
	for _, num := range array {
		sum += num
	}
	for _, num := range array {
		res = append(res, sum-num)
	}
	return res
}

func AccumulatingSumArray(array []int) []int {
	if len(array) == 0 {
		return []int{}
	}
	sum := 0
	var res []int
	for _, num := range array {
		sum += num
		res = append(res, sum)
	}
	return res
}

func AccumulatingProduct(array []int) []int {
	// https://edabit.com/challenge/qhhDLY3mtyZXikS3S
	if len(array) == 0 {
		return []int{}
	}
	sum := 1
	var res []int
	for _, num := range array {
		sum *= num
		res = append(res, sum)
	}
	return res
}

func Boxes(weights []int) int {
	countOfBoxes, cumWeight := 0, 0
	for _, item := range weights {
		if cumWeight == 10 || cumWeight+item > 10 {
			countOfBoxes += 1
			cumWeight = 0
		}
		cumWeight += item
	}
	return countOfBoxes + 1
}

func IsSubset(arr1 []int, arr2 []int) bool {
	allMatchCounter := 0
	for _, num2 := range arr2 {
		for _, num1 := range arr1 {
			if num1 == num2 {
				allMatchCounter += 1
			}
		}
	}
	return len(arr1) == allMatchCounter
}

func FilterUnique(arr []string) []string {
	// https://edabit.com/challenge/n3uivD8qXEq2D3ZMh
	var res []string
	for _, word := range arr {
		letterCounter := 0
		for _, letter := range word {
			if strings.Count(word, string(letter)) > 1 {
				letterCounter += 1
			}
		}
		if letterCounter == 0 {
			res = append(res, word)
		}
	}
	fmt.Println(res)
	return res
}

func RemoveDuplicates(arr []string) []string {
	// https://edabit.com/challenge/eLH7c93ujbacHTwmy
	var uniqueStrings []string
	mapOfStructs := map[string]struct{}{}
	// https://golangbot.com/maps/#:~:text=no%20runtime%20error.-,Map%20of%20structs,-So%20far%20we
	for _, word := range arr {
		if _, exists := mapOfStructs[word]; exists {
			continue
		}
		uniqueStrings = append(uniqueStrings, word)
		mapOfStructs[word] = struct{}{}
	}
	fmt.Println(uniqueStrings)
	return uniqueStrings
}

func SortByLength(arr []string) []string {
	// https://edabit.com/challenge/LDNzJS27GY7aWNcBv
	// https://edabit.com/challenge/PM6j8xFZCuJmLwZkx
	// https://yourbasic.org/golang/how-to-sort-in-go/
	sort.Slice(arr, func(i, j int) bool {
		return len(arr[i]) < len(arr[j])
	})
	fmt.Println(arr)
	return arr
}

func CountDuplicates(str string) int {
	// https://edabit.com/challenge/SS2T2QAsf5jWWZJ24
	var duplicatesCounter int
	mapOfStructs := map[string]struct{}{}
	// https://golangbot.com/maps/#:~:text=no%20runtime%20error.-,Map%20of%20structs,-So%20far%20we
	for _, letter := range str {
		if _, exists := mapOfStructs[string(letter)]; exists {
			duplicatesCounter++
		}
		mapOfStructs[string(letter)] = struct{}{}
	}
	fmt.Println("Count of duplicates is: ", duplicatesCounter)
	return duplicatesCounter
}

func StringToPhoneNumber(str string) string {
	// https://edabit.com/challenge/ckM4CHgAn9pcNAaSH
	conversionMap := map[string]int{
		"abc": 2, "def": 3, "ghi": 4, "jkl": 5,
		"mno": 6, "pqrs": 7, "tuv": 8, "wxyz": 9,
	}
	converted := strings.ToLower(str)
	fmt.Println(converted)
	for _, letter := range converted {
		for key, value := range conversionMap {
			if strings.Contains(key, string(letter)) {
				// https://stackoverflow.com/a/10105983
				fmt.Println(key, " Contains ", string(letter), "substitute is ", strconv.Itoa(value))
				converted = strings.Replace(converted, string(letter), strconv.Itoa(value), 1)
			}
		}
	}
	fmt.Println("Converted is ", converted)
	return converted
}

func NthDifference(array []int) int {
	// https://edabit.com/challenge/jXAutCihEANSFJdCj
	for len(array) > 0 {
		for id := range array {
			if id < len(array)-1 {
				array[id] = array[id+1] - array[id]
			}
		}
		array = array[0 : len(array)-1]
		if len(array) == 1 {
			break
		}
	}
	fmt.Println(array[0])
	return array[0]
}

func AscDescNone(arr []int, str string) []int {
	if str == "Asc" || str == "Des" {
		sort.Slice(arr, func(i, j int) bool {
			if str == "Asc" {
				return (arr[i]) < arr[j]
			} else {
				return arr[i] > arr[j]
			}
		})
	}
	fmt.Println(arr)
	return arr
}

// https://go.dev/doc/code#PackagePaths
// https://www.golangprograms.com/go-language/struct.html

// Employee Declare an Interface Type and methods does not have a body
type Employee interface {
	PrintName() string                // Method with string return type
	PrintAddress(id int)              // Method with int parameter
	PrintSalary(b int, t int) float64 // Method with parameters and return type
}

// Emp user-defined type
type Emp int

// PrintName method to print employee name
func (e Emp) PrintName(name string) {
	fmt.Println("Employee Id:\t", e)
	fmt.Println("Employee Name:\t", name)
}

// PrintSalary method to calculate employee salary
func (e Emp) PrintSalary(basic int, tax int) int {
	var salary = (basic * tax) / 100
	return basic - salary
}

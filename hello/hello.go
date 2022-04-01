package main

import (
	"example.com/arrays"
	"example.com/greetings"
	"fmt"
	"log"
)

import "rsc.io/quote"

// https://go.dev/tour/list
func main() {
	mySliceStr := []string{"Go", "Slices", "Are", "Powerful"}
	newSlice := arrays.AppendToSlice(mySliceStr, "magical")
	fmt.Printf("mySlice = %v\n", newSlice)
	arrays.Maps()
	arrays.AddGenerics(6.3, 10.9)
	arrays.PrintSlice(mySliceStr)
	arrays.FilterUnique([]string{"ABCDE", "DDEB", "BED", "CCA", "BAC"})
	arrays.RemoveDuplicates([]string{"javascript", "python", "python", "ruby", "javascript", "c", "ruby"})
	arrays.SortByLength([]string{"Turing", "Einstein", "Jung"})
	arrays.CountDuplicates("Gen'rals gathered in their masses,Just like witches at black masses")
	arrays.StringToPhoneNumber("abc-def-ghi-jkl!!")
	arrays.NthDifference([]int{8, 9, 2, 5, 4, 3, 3, 1, 6})
	arrays.AscDescNone([]int{7, 8, 11, 66}, "None")

	fmt.Println(arrays.MaxNumberInList([]int{1000, 1001, 857, 1}))
	fmt.Println(arrays.DifferenceMaxMin([]int{44, 32, 86, 19}))
	fmt.Println(arrays.EvenOddTransform([]int{1, 2, 3}, 1))
	fmt.Println(arrays.LetterCounter([][]string{
		{"D", "E", "Y", "H", "A", "D"},
		{"C", "B", "Z", "Y", "J", "K"},
		{"D", "B", "C", "A", "M", "N"},
		{"F", "G", "G", "R", "S", "R"},
		{"V", "X", "H", "A", "S", "S"},
	}, "D"))
	fmt.Println(arrays.SpinAround([]string{"right", "right", "right", "right", "right", "right", "right", "right"}))
	fmt.Println(arrays.SquarePatch(5))
	fmt.Println(arrays.RemoveSpecialCharacters("%fd76$fd(-)6GvKlO."))
	fmt.Println(arrays.ArrayElementSum([]int{10, 20, 30, 40, 50, 60}))
	fmt.Println(arrays.AccumulatingSumArray([]int{1, 5, 7}))
	fmt.Println(arrays.AccumulatingProduct([]int{1, 2, 3, 4}))
	fmt.Println(arrays.Boxes([]int{1, 3, 3, 3, 2, 1, 1, 9, 7, 10, 8, 6, 1, 2, 9}))
	fmt.Println(arrays.IsSubset([]int{1, 2}, []int{3, 5, 9, 1}))

	fmt.Println("4 = 3 is", arrays.Same(4, 3))
	fmt.Println("aa = aa is", arrays.Same("aa", "aa"))
	fmt.Println("4.1 = 4.15 is", arrays.Same(4.1, 4.15))

	fmt.Println("4 + 3 =", arrays.Add(4, 3))
	fmt.Println("4.1 + 3.2 =", arrays.Add(4.1, 3.2))

	v := greetings.Vertex{X: 3, Y: 4}
	v.Scale(10)
	greetings.Scale(&v, 10)
	fmt.Println(v.Abs())
	fmt.Println(greetings.AbsNormal(v))
	fmt.Println(quote.Go())
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	// A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin", "Sesan"}
	//message, err := greetings.Hello("Sesan")

	// Request greeting messages for the names.
	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}
	for _, message := range messages {
		fmt.Println(message)
	}
	// https://go.dev/blog/maps#:~:text=Exploiting%20zero%20values

	var student1 = "John" //type is string
	var student2 = "Jane" //type is inferred
	x := 2                //type is inferred

	//var a, b, c, d = 1, 3, 5, 7
	const PI = 3.14
	// https://www.w3schools.com/go/go_constants.php#:~:text=Multiple%20Constants%20Declaration
	var (
		a int
		b int    = 1
		c string = "hello"
	)
	const (
		A int = 1
		B     = 3.14
		C     = "Hi!"
	)

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// Array of integers
	arr2 := [5]int{4, 5, 6, 7, 8}
	fmt.Println(arr2)

	// Array of Strings
	var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}
	fmt.Println(cars)

	// https://www.w3schools.com/go/go_arrays.php#:~:text=With%20the%20%3A%3D%20sign%3A-,Syntax,-array_name%20%3A%3D%20%5Blength%5Ddatatype
	// array_name := [length]datatype{values} // here length is defined
	// array_name := [...]datatype{values} // here length is inferred

	arr3 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(len(arr3))
	TestCount(1)

}

func TestCount(x int) int {
	if x == 11 {
		return 0
	}
	fmt.Println(x)
	return TestCount(x + 1)
}

package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"time"

	"github.com/martinskou/safari"
)

/*
func compare(a, b int32) bool {
	return a < b
}
*/

func RandIntSlice(length int) []int {
	slc := []int{}
	for i := 0; i < length; i++ {
		slc = append(slc, rand.Intn(length))
	}
	return slc
}

func test1(listInt []int) {
	defer safari.TimeTrack(time.Now(), "Quicksort Generic")
	//	listInt := []int{3, 5, 1, 2, 6, 7, 4, 2}
	//fmt.Println(listInt)
	safari.SortAll(listInt, safari.OpCompAsc[int])
	//	quicksort.SortAll(listInt, func(a, b int) bool { return a > b })
	fmt.Println(listInt)
}
func test2(listInt []int) {
	defer safari.TimeTrack(time.Now(), "Build in sort")
	//	listInt := []int{3, 5, 1, 2, 6, 7, 4, 2}
	//	fmt.Println(listInt)
	sort.Ints(listInt)
	//	fmt.Println(listInt)
}
func test3() {
	lst := []string{"martin", "lone", "liva", "marcus", "albert", "valdemar"}
	safari.SortAll(lst, safari.OpCompAsc[string])
	fmt.Println(lst)
}

type Person struct {
	Name string
	Age  int
}

func AgeComp(a, b Person) bool {
	return a.Age < b.Age
}

func test4() {
	lst := []Person{{"martin", 49}, {"lone", 51}, {"liva", 12}, {"marcus", 9}, {"albert", 9}, {"valdemar", 1}}
	safari.SortAll(lst, AgeComp)
	fmt.Println(lst)
}

func test5() {
	lst := []Person{{"martin", 49}, {"lone", 51}, {"liva", 12}, {"marcus", 9}, {"albert", 9}, {"valdemar", 1}}
	mp := safari.SliceToMap(lst, func(p Person) string { return p.Name })
	fmt.Println(mp)

	keys := safari.MapKeys(mp)
	fmt.Println(keys)

	values := safari.MapValues(mp)
	fmt.Println(values)

	adults := safari.Filter(lst, func(p Person) bool { return p.Age > 17 })
	fmt.Println(adults)

	adultsmap := safari.FilterMap(mp, func(p Person) bool { return p.Age > 17 })
	fmt.Println(adultsmap)

	adultages := safari.Map(lst, func(p Person) int { return p.Age })
	fmt.Println(adultages)

}

func test6() {
	// y, e := strconv.Atoi(x)
	x := "15"
	fmt.Println(safari.Must(strconv.Atoi(x)))
}

func main() {
	items := RandIntSlice(100)
	s1 := make([]int, len(items))
	copy(s1, items)
	test1(s1)
	test2(s1)

	test3()
	test4()

	test5()

	test6()
	/*
		listFloat := []float32{3.3, .5, 1.5, 22.3, 64.2, 7.1, 1.4, 2}
		sortedFloats := quicksort.SortAll(listFloat, quicksort.CompareNumber[float32])
		fmt.Println(sortedFloats)
	*/
}

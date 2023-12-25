package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Arrays struct {
	Arr []int
}

func (A *Arrays) Init(count int) {
	A.Arr = make([]int, count)
	for i := 0; i < count; i++ {
		A.Arr[i] = rand.Intn(10)
	}
}

func (a *Arrays) Bubble() {
	for i := len(a.Arr) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if a.Arr[j] > a.Arr[j+1] {
				a.Arr[j], a.Arr[j+1] = a.Arr[j+1], a.Arr[j]
			}
		}
	}
}

func (a *Arrays) BubbleOpt() {
	for i := len(a.Arr) - 1; i > 0; i-- {
		isSorted := true
		for j := 0; j < i; j++ {
			if a.Arr[j] > a.Arr[j+1] {
				a.Arr[j], a.Arr[j+1] = a.Arr[j+1], a.Arr[j]
				isSorted = false
			}
		}
		if isSorted {
			return
		}
	}
}

func (a *Arrays) Insertion() {
	for i := 1; i < len(a.Arr); i++ {
		for j := i - 1; j >= 0 && (a.Arr[j] > a.Arr[j+1]); j-- {
			a.Arr[j+1], a.Arr[j] = a.Arr[j], a.Arr[j+1]
		}
	}
	//	fmt.Println("arr=", a.Arr)
}

func (a *Arrays) InsertionBin() {
	var j int
	for i := 1; i < len(a.Arr); i++ {
		t := a.Arr[i]
		for j = i - 1; j >= 0 && (a.Arr[j] > a.Arr[j+1]); j-- {
			//	a.Arr[j+1], a.Arr[j] = a.Arr[j], a.Arr[j+1]
			a.Arr[j+1] = a.Arr[j]
		}
		a.Arr[j+1] = t
	}
	fmt.Println("arr=", a.Arr)
}

func binSearch(x int) int {
	return 0
}

func (a *Arrays) Shell() {
	for gap := len(a.Arr) / 2; gap > 0; gap /= 2 {
		for i := gap; i < len(a.Arr); i++ {
			for j := i; j >= gap && (a.Arr[j] < a.Arr[j-gap]); j -= gap {
				a.Arr[j], a.Arr[j-gap] = a.Arr[j-gap], a.Arr[j]
			}
		}
	}
	//	fmt.Println("shell=", a)
}

func main() {
	//	count := 1000
	var Arr Arrays
	for i := 10; i <= 10; i *= 10 {
		Arr.Init(i)
		start := time.Now()
		Arr.InsertionBin()
		duration := time.Since(start)
		fmt.Println("count=", i, "time = ", duration)

	}
}

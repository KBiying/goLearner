package main

import (
	"reflect"
	"testing"
)


func Sum(numbers []int) int{
	sum:= 0
	for _, number := range numbers{
		sum+= number
	}
	return sum
}
func TestSum(t *testing.T){
	t.Run("collection of 5 numbers", func(t *testing.T){
		numbers := []int{1,2,3,4,5}
		got := Sum(numbers)
		want:=15
		if want!=got{
			t.Errorf("got %d want %d given, %v",want, got, numbers)
		}
	})
}
func SumAll(numbersToSum ...[]int) (sums []int){
	// 1. index
	// lenofSlice:=len(numbersToSum)
	// sums = make([]int,lenofSlice)

	// for i, slicesOfNumbers := range numbersToSum{
	// 	sums[i] = Sum(slicesOfNumbers)
	// }
	// 2. append
	for _, slicesOfNumbers := range numbersToSum{
		sums = append(sums, Sum(slicesOfNumbers))
	}

	return
}

func TestSumAll(t *testing.T){
	got := SumAll([]int{1,2}, []int{0,9})
	want:=[]int{3,9}
	// want:="bob"
	if !reflect.DeepEqual(got,want){
		t.Errorf("got %v want %v", got, want)
	}
}

func SumTail(numbersToSum ...[]int) (sum []int){
	// add all number expect the first one of a slice
	for _, slices := range(numbersToSum){
		if len(slices) == 0{
			sum = append(sum, 0)
		} else{
			tail := slices[1:]
			sum = append(sum, Sum(tail))
		}
		
	} 
	return
}
func TestSumTail(t *testing.T){

	checkSum := func(t *testing.T, got, want []int){
		if !reflect.DeepEqual(got,want){
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("make the sums of the slices", func(t *testing.T) {
		got:= SumTail([]int{1,2}, []int{0,9})
		want:= []int{2,9}
		checkSum(t, got, want)
	})
	t.Run("safely sum empty slices", func(t *testing.T) {
		got:= SumTail([]int{}, []int{0,9})
		want:= []int{0,9}
		checkSum(t, got, want)

	})

}
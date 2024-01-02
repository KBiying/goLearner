package main

import (
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
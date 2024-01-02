package interation

import (
	"fmt"
	"testing"
)

const repeatCharacter = "a"
const repeatCount = 5
func Repeat(character string, count int) string{
	var repeated string
	for i:=0; i<count; i++{
		repeated += character
	}
	return repeated
}
func ExmapleRepeat(){
	repeated := Repeat("a",5)
	fmt.Println(repeated)
	// Output: "aaaaa"
}
func TestRepeat(t *testing.T){
	repeated := Repeat(repeatCharacter, repeatCount)
	var expected string
	for i:= 0; i<repeatCount;i++{
		expected += repeatCharacter
	}
	if expected != repeated{
		t.Errorf("expected '%q', but got '%q'", expected, repeated)
	}
}
func BenchmarkRepeat(b* testing.B){
	for i:=0; i<b.N; i++{
		Repeat(repeatCharacter,repeatCount)
	}
}


package main

import "testing"

func TestfindMaximalSum( t *testing.T){
	var test = struct{
		size int
		expected int
	}{
		10000,9240,
	}

	MaximumSum:= findMaximalSum(test.size)
	if MaximumSum != test.expected{
		t.Error("Func is incorrect, got: " , 	findMaximalSum(test.size)," want: " ,test.expected)
	}
}



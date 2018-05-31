/*Найти натуральное число от 1 до 10 000 с максимальной
суммой делителей.*/

package main

func findMaximalSum( n int) int{
	sum:=0
	max:=0
	result:=0

	for i := 1; i < n; i++ {
		for j:=1; j<=i;j++{
			if i%j==0 {
				sum += j
			}
		}
		if sum>max {
			max = sum
			result = i
		}
		sum = 0
	}
	return result
}

func main(){
}

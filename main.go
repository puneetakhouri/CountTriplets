package main

import "fmt"

func main() {

	//arr := []int{1, 5, 3, 2}
	//arr := []int{1, 2, 3}
	//arr := []int{3, 2, 7}
	arr := []int{1, 3, 5, 7, 9, 4}
	count := CountTriplets(arr)

	if count > 0 {
		fmt.Printf("%d", count)
	} else {
		fmt.Printf("-1")
	}
}

//CountTriplets test
func CountTriplets(arr []int) int {
	var dataMap = make(map[int]int)
	count := 0

	for i := 0; i < len(arr)-1; i++ {
		dataMap[arr[i]] = i
		for j := i + 1; j < len(arr); j++ {
			fmt.Printf("i=%d, arr[%d]=%d, j=%d, arr[%d]=%d\n", i, i, arr[i], j, j, arr[j])

			index, indexDataExists := dataMap[arr[j]]
			if indexDataExists && index == -1 { //Some sum exists.. increment counter
				fmt.Printf("arr[%d] = %d exists, incrementing counter\n", j, arr[j])
				count = count + 1
				dataMap[arr[j]] = -2
			} else {
				dataMap[arr[j]] = j
			}

			sum := arr[j] + arr[i]
			sumIndex, sumDataExists := dataMap[sum]
			if sumDataExists && sumIndex != -1 {
				fmt.Printf("sum = %d exists, incrementing counter\n", sum)
				count = count + 1
				dataMap[sum] = -2
			} else {
				dataMap[sum] = -1
			}

			for key, value := range dataMap {
				fmt.Printf("key=%d, value=%d\n", key, value)
			}
			fmt.Println()
		}
	}
	return count
}

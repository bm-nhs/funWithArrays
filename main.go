package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args[1:]
	//testArray is an array of strings
	testArray := []string{"one", "two", "three", "four", "five"}
	out := findMatch(testArray, arguments)

	//if out length is 0, no match was found

	if len(out) == 0 {
    fmt.Println("No match found")
  } else {
    fmt.Println(out)
  }

}

func findMatch(arr1 []string, arr2 []string) []string {
	matchingValues := []string{}
	for _, i := range arr1 {
		for _, j := range arr2 {
			if i == j {
				matchingValues = append(matchingValues, i)
			}
		}
	}

	return removeDuplicateValues(matchingValues)
}

func removeDuplicateValues(stringSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}

	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.
	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
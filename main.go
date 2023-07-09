package main

import "fmt"

func removestring(stringList []string, stringToRemove string) []string {
	stringListLengs := len(stringList)
	for i, ss := range stringList {
		if ss == stringToRemove {
			stringList[stringListLengs-1], stringList[i] = stringList[i], stringList[stringListLengs-1] //
			return stringList[:stringListLengs-1]
		}
	}
	return stringList
}

func main() {
	str := []string{"a", "b", "c", "d", "e"}

	strr := removestring(str, "c")

	fmt.Print(strr)
}

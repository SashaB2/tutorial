package main

import (
	"fmt"
	"sort"
	"time"
)

type SliceTime []time.Time

func (s SliceTime) Len() int {
	return len(s)
}

func (s SliceTime) Less(i, j int) bool {
	return s[i].Before(s[j])
}

func (s SliceTime) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	// Створення масиву time.Time
	myTimes := []time.Time{
		time.Now(),
		time.Now().AddDate(0, -1, 0),
		time.Now().AddDate(0, 1, 0),
	}

	// Виведення масиву перед сортуванням
	fmt.Println("Before sorting:")
	for _, t := range myTimes {
		fmt.Println(t)
	}

	// Сортування масиву
	sort.Sort(SliceTime(myTimes))

	// Виведення масиву після сортування
	fmt.Println("\nAfter sorting:")
	for _, t := range myTimes {
		fmt.Println(t)
	}
}

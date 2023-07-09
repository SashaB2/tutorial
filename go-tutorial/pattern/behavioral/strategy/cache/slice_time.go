package cache

import "time"

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

type Lru struct {
}

package main

import (
	"fmt"
	"time"
	cache3 "tutorial/go-tutorial/pattern/behavioral/strategy/cache"
	calc2 "tutorial/go-tutorial/pattern/behavioral/strategy/calc"
	duck2 "tutorial/go-tutorial/pattern/behavioral/strategy/duck"
)

func main() {
	duck := duck2.Duck{}

	duckFly := &duck2.DuckSoar{}

	duck.Fly(duckFly)

	duckFlyWings := &duck2.DuckFlapWings{}

	duck.Fly(duckFlyWings)

	//------------------------------------------------------

	calc := calc2.Calc{}
	calc.A = 5
	calc.B = 5

	add := calc2.Add{}

	calc.SetStrategyAction(add)
	fmt.Println(calc.DoAction())

	calc.SetStrategyAction(calc2.Multiply{})
	fmt.Println(calc.DoAction())

	//-------------------------------------------------------

	fifo := &cache3.Fifo{}
	cache := cache3.InitCache(fifo)
	cache.Add("0", map[string]time.Time{"0": time.Now()})
	cache.Add("1", map[string]time.Time{"1": time.Now()})
	cache.Add("2", map[string]time.Time{"2": time.Now()})
	cache.Add("3", map[string]time.Time{"3": time.Now()})

	for key, value := range cache.GetAll() {
		fmt.Printf("key: %s; value: %s", key, value)
		fmt.Println("")
	}

	lru := &cache3.Lru{}
	cache.SetEviction(lru)
	cache.AddDefault("4", "4")
	cache.AddDefault("5", "5")
	cache.AddDefault("6", "6")

}

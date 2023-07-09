package duck

import "fmt"

type Duck struct {
}

func (d *Duck) Fly(fly DuckFly) {
	fly.fly()
}

type DuckFly interface {
	fly()
}

type DuckFlapWings struct {
}

func (d *DuckFlapWings) fly() {
	fmt.Println("Duck flaps wings")
}

type DuckSoar struct {
}

func (s *DuckSoar) fly() {
	fmt.Println("Duck soars")
}

type DuckDive struct {
}

func (d *DuckDive) fly() {
	fmt.Println("Duck dives")
}

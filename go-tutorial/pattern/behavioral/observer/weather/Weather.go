package weather

import (
	"fmt"
	"tutorial/go-tutorial/helpers"
)

type IObservable interface {
	Add(observer IObserver)
	Remove(observer IObserver)
	Notify()
}

type Weather struct {
	observables []IObserver
}

func (w *Weather) Add(observer IObserver) {
	w.observables = append(w.observables, observer)
}

func (w *Weather) Remove(observer IObserver) {
	helpers.DeleteOfSlice(&w.observables, observer)
}

func (w *Weather) Notify() {
	for _, v := range w.observables {
		v.Update()
	}
}

type IObserver interface {
	Update()
}

type Phone struct {
	weather IObservable
}

func (p Phone) Update() {
	fmt.Print("Update")
}

func (p Phone) setWeather(observable IObservable) {
	p.weather = observable
}

type Desktop struct {
	weather IObservable
}

func (d Desktop) Update() {
	fmt.Print("Update")
}

func (d Desktop) setWeather(observable IObservable) {
	d.weather = observable
}

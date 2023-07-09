package item

import "fmt"

type Subject interface {
	Subscribe(observer Observer)
	Unsubscribe(observer Observer)
	NotifyAll()
}

type Item struct {
	Observer []Observer
	Name     string
	InStock  bool
}

func (i *Item) Subscribe(observer Observer) {
	i.Observer = append(i.Observer, observer)
}

func (i *Item) Unsubscribe(observer Observer) {
	i.Observer = removeObserver(i.Observer, observer)
}

func (i *Item) NotifyAll() {
	for _, observ := range i.Observer {
		observ.Update(i.Name)
	}
}

func removeObserver(observerList []Observer, observerToRemove Observer) []Observer {
	observerListLengs := len(observerList)
	for i, observer := range observerList {
		if observer.GetId() == observerToRemove.GetId() {
			observerList[observerListLengs-1], observerList[i] = observerList[i], observerList[observerListLengs-1]
			return observerList[:observerListLengs-1]
		}
	}
	return observerList
}

type Observer interface {
	Update(string)
	GetId() string
}

type Customer struct {
	Id string
}

func (c Customer) Update(s string) {
	fmt.Println("Send email to: " + s)
}

func (c Customer) GetId() string {
	return c.Id
}

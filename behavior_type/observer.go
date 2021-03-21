package behavior_type

import "fmt"

type subject interface {
	register(observer Observer)
	deregister(observer Observer)
	notifyAll()
}

type item struct {
	observerList []Observer
	name         string
	inStock      bool
}

func newItem(name string) *item {
	return &item{
		name: name,
	}
}

func (it *item) register(observer Observer) {
	it.observerList = append(it.observerList, observer)
}

func (it *item) deregister(observer Observer) {
	it.observerList = removeFromslice(it.observerList, observer)
}

func (it *item) notifyAll() {
	for _, ob := range it.observerList {
		ob.update(it.name)
	}
}

func (it *item) updateAvailability() {
	fmt.Printf("Item %s is now in stock\n", it.name)
	it.inStock = true
	it.notifyAll()
}

func removeFromslice(observerList []Observer, removeObserver Observer) []Observer {
	for index, observer := range observerList {
		removeIndex := 0
		if observer.getID() == removeObserver.getID() {
			removeIndex = index
		}
		if removeIndex <= 0 {
			continue
		}
		return append(observerList[:index], observerList[index+1:]...)
	}
	return observerList
}

type Observer interface {
	update(string)
	getID() string
}

type customer struct {
	id string
}

func (c *customer) update(itemName string) {
	fmt.Printf("Sending email to customer %s for item %s\n", c.id, itemName)
}

func (c *customer) getID() string {
	return c.id
}

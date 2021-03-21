package create_type

import "fmt"

type Car struct {
	id    int
	model string
	price float64
}

func (c Car) clone() Car {
	return Car{
		id:    c.id,
		model: c.model,
		price: c.price,
	}
}

func (c Car) printDetails() {
	fmt.Printf("cat id:%d, name:%s, price:%v\n", c.id, c.model, c.price)
}

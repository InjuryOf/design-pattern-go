package create_type

import (
	"fmt"
	"sync"
)

type book struct {
	name string
	price float64
}

var ibook *book
var once sync.Once

// 单例初始化
func GetBookInstance() *book {
	once.Do(func() {
		ibook = &book{}
	})
	return ibook
}


func (b book) printDetails() {
	fmt.Printf("book %p name:%s,  price:%v\n", &b, b.name, b.price)
}

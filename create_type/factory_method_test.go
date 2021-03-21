package create_type

import (
	"fmt"
	"testing"
)

func TestSimpleFactory(t *testing.T) {
	fmt.Printf("=====【简单工厂】=====\n")
	land, _ := getSimpleFactory(1)
	ocean, _ := getSimpleFactory(2)
	land.printDetails()
	ocean.printDetails()
}


func TestFactoryMethod(t *testing.T) {
	fmt.Printf("=====【工厂方法】=====\n")
	// 方式一
	f := new(FactoryMethod)
	land := f.CreateLogistics(1)
	ocean := f.CreateLogistics(2)
	land.printDetails()
	ocean.printDetails()

	// 方式二
	ft := FactoryMethod{ltype: 1}
	landTwo := ft.CreateLogisticsTwo()
	landTwo.printDetails()

}

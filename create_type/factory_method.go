package create_type

import "fmt"

type iLogistics interface {
	setName(name string)
	setType(lType int)
	getName() string
	getType() int
	printDetails()
}

// 物流
type Logistics struct {
	name  string
	lType int
}

func (l *Logistics) setName(name string) {
	l.name = name
}

func (l *Logistics) getName() string {
	return l.name
}

func (l *Logistics) setType(lType int) {
	l.lType = lType
}

func (l *Logistics) getType() int {
	return l.lType
}

func (l *Logistics) printDetails() {
	fmt.Printf("logistics name = %s type = %d\n", l.getName(), l.getType())
}

// 陆运物流
type LandLogistics struct {
	Logistics
}

func newLandLogistics() iLogistics {
	return &LandLogistics{
		Logistics{
			name:  "陆运物流",
			lType: 1,
		},
	}
}

// 海运物流
type OceanLogistics struct {
	Logistics
}

func newOceanLogistics() iLogistics {
	return &OceanLogistics{
		Logistics{
			name:  "海运物流",
			lType: 2,
		},
	}
}

// 获取实例化工厂对象 - 简单工厂
func getSimpleFactory(ltype int) (iLogistics, error) {
	switch ltype {
	case 1:
		return newLandLogistics(), nil
	case 2:
		return newOceanLogistics(), nil
	}
	return nil, fmt.Errorf("no type")
}

// 工厂方法
type FactoryInterface interface {
	CreateLogistics(ltype int) iLogistics
}

type FactoryMethod struct {
	ltype int
}

func (fm FactoryMethod) CreateLogistics(ltype int) iLogistics {
	switch ltype {
	case 1:
		return newLandLogistics()
	case 2:
		return newOceanLogistics()
	}
	return nil
}

func (fm FactoryMethod) CreateLogisticsTwo() iLogistics {
	switch fm.ltype {
	case 1:
		return newLandLogistics()
	case 2:
		return newOceanLogistics()
	}
	return nil
}

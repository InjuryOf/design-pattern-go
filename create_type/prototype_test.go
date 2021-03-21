package create_type

import "testing"

func TestPrototype(t *testing.T) {
	kmr := Car{id: 1101, model: "丰田凯美瑞", price: 21000.50}
	kmr.printDetails()
	copyCar := kmr.clone()
	copyCar.printDetails()
}

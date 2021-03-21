package create_type

import "testing"

func TestSingle(t *testing.T)  {
	book := GetBookInstance()
	book.name  = "黑客与画家"
	book.printDetails()
	bookNew := GetBookInstance()
	bookNew.printDetails()
}
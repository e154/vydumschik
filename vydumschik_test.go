package vydumschik

import (
	"testing"
	"fmt"
)

func TestVydumschik(t *testing.T) {

	fmt.Println("--------------------------------")
	name := new(Name)
	first := name.First_name("male")
	middle := name.Middle_name("male")
	sure := name.Sur_name("male")
	fmt.Printf("%s %s %s\n", first, middle, sure)

	fmt.Println(name.Full_name("male"))

	for i:=0;i<3;i++ {
		fmt.Println(name.Full_name("female"))
	}

	fmt.Println("--------------------------------")

	address := new(Address)
	for i:=0;i<3;i++ {
		fmt.Println(address.Street())
	}

	for i:=0;i<3;i++ {
		fmt.Println(address.Street_address())
	}

	fmt.Println("--------------------------------")
	lorem := new(Lorem)
	fmt.Println(lorem.Worlds(5))
	fmt.Println("--------------------------------")

	for i:=0;i<3;i++ {
		fmt.Println(lorem.Speech(1))
	}
	fmt.Println("--------------------------------")
	fmt.Println(lorem.Paragraphs(2))
	fmt.Println("--------------------------------")
	str := lorem.Bytes(112)
	fmt.Println(str)
	fmt.Println(len(str))
}
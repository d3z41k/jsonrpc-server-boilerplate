package services

import (
	"fmt"
	"strconv"
)

type Name struct {
	Name string
	Age  int
}

type HelloService struct{}

func (hm *HelloService) SayHello(in *Name, out *string) error {
	fmt.Println("call SayHello", in)

	*out = "Hello " + in.Name + ", your age is " + strconv.Itoa(in.Age) + "."
	return nil
}

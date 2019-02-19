package main

import "fmt"

type Obj struct {
	id uint
}

func (obj *Obj) setId(id uint) {
	obj.id = id
}

func (obj *Obj) getId() uint {
	return obj.id
}

func main() {
	obj := Obj{1}
	fmt.Println(obj)
	obj.setId(123)
	fmt.Println(obj)
	fmt.Println(obj.getId())
}

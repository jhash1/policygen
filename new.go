package main

import "fmt"

type Example struct {
	Id   []int
	Name []string
}

func (data *Example) AppendOffer(id int, name string) {
	data.Id = append(data.Id, id)
	data.Name = append(data.Name, name)
}

var MyMap map[string]*Example

func main() {
	obj := &Example{[]int{}, []string{}}
	obj.AppendOffer(1, "SomeText")
	MyMap = make(map[string]*Example)
	MyMap["key1"] = obj
	fmt.Println(MyMap)
}

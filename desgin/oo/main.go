package oo

import "fmt"

type PersonInfo interface {
    Print()
}

type MyData struct {
    Name string
    Age int
}

func NewMyData() *MyData {
    tmp := &MyData{"luv",30}
    return tmp
}
func (this *MyData)Print() {
    fmt.Println("name:",this.Name, " age:",this.Age)
}



func NewMyData2() PersonInfo {
    return &MyData{"liuhui",30}
}

func main() {
    c := NewMyData()
    fmt.Println("name:",c.Name, " age:",c.Age)

    c2 := NewMyData2()
    c2.Print()
}
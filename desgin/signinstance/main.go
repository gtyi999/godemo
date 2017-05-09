package signinstance

import "fmt"

type MyClass struct{
}

var myapp *MyClass

func GetInstance() *MyClass {
    if myapp == nil {
        myapp = &MyClass{}
    }
    return myapp
}



func init() {
    app := GetInstance()
    fmt.Println(&app)
    app = GetInstance()
    fmt.Println(&app)
    appnew := GetInstance()
    fmt.Println(&appnew)
}
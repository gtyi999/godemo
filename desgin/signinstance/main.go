package signinstance

import (
    "fmt"
    "sync"
)

type MyClass struct{
}

var myapp *MyClass
var lock *sync.Mutex = &sync.Mutex{}

func GetInstance() *MyClass {
    if myapp == nil {
        lock.Lock()
        defer lock.Unlock()
        if myapp == nil {
            myapp = &MyClass {}
        }
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
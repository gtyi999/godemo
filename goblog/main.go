package main

import (
    "github.com/julienschmidt/httprouter"
    "github.com/qqqc/godemo/goblog/utils"
    "net/http"
    "fmt"
    //"github.com/qqqc/godemo/goblog/utils"

    _ "github.com/qqqc/godemo/goblog/controller"
)




func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w,"hello")
}

func main() {

    applog := utils.NewAppLog()
    applog.Info("this is a info test")
    applog.Debug("this is a debug test")

    r := httprouter.New()
    r.GET("/",Index)
    http.ListenAndServe(":8888",r)
}

package main

import (
    "github.com/julienschmidt/httprouter"
    "github.com/qqqc/godemo/goblog/utils"
    "net/http"
    C "github.com/qqqc/godemo/goblog/controller"
    _ "github.com/qqqc/godemo/goblog/models"
)






func main() {

    applog := utils.NewAppLog()
    applog.Info("this is a info test")
    applog.Debug("this is a debug test1")


    applog = utils.NewAppLog()
    applog.Debug("this is a debug test2")

    r := httprouter.New()
    r.GET("/adduser",C.UserRegist)
    r.GET("/",C.HomeIndex)
    http.ListenAndServe(":8888",r)
}

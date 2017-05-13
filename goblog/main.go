package main

import (
    "github.com/julienschmidt/httprouter"
    "github.com/qqqc/godemo/goblog/utils"
    "net/http"
    C "github.com/qqqc/godemo/goblog/controller"
)

func main() {


    applog := utils.NewAppLog()
    applog.Info("this is a info test")
    applog.Debug("this is a debug test1")


    applog = utils.NewAppLog()
    applog.Debug("this is a debug test2")



    r := httprouter.New()
    r.GET("/",C.HomeIndex)
    r.GET("/about",C.HomeAbout)
    r.GET("/contact",C.HomeContact)











    r.GET("/adduser",C.UserRegist)

    //设置静态文件路径
    r.ServeFiles("/static/*filepath",http.Dir(C.ThemePath))


    http.ListenAndServe(":8888",r)



}

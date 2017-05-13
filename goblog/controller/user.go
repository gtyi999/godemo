package controller

import (
    "net/http"
    "github.com/julienschmidt/httprouter"
    "github.com/qqqc/godemo/goblog/models"
    "fmt"
    "time"
)

//用户主页页面 http://localhost/u/profile
func UserIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    u := models.GetUserById(3)
    if u==nil {
        fmt.Fprint(w,"hello guest")
    } else {
        fmt.Fprint(w,"hello ", u.UserNicename)
    }
}

//用户注册页面 http://localhost/u/regist
func UserRegist(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

    udata :=models.WpUsers{UserLogin:"luv",UserPass:"331234958", UserNicename:"abner",UserEmail:"331234958@qq.com", UserUrl:"http://www.web012.com", UserRegistered:time.Now(),UserActivationKey:"#aadf@$",UserStatus:0,DisplayName:"小辉"}
    id,e := models.AddUser(udata)
    if e==nil {
        fmt.Fprint(w,"add success id",id)
    } else {
        fmt.Fprint(w,"add failed")
    }
}

//用户登录页面 http://localhost/u/login
func UserLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}

//重置密码页面
func UserRestPass(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}


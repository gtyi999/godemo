package controller

import (
    "net/http"
    "github.com/julienschmidt/httprouter"
    "github.com/qqqc/godemo/goblog/models/xorm"
    "github.com/qqqc/godemo/goblog/models"
    "fmt"
    "time"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    u := models.GetUserById(3)
    if u==nil {
        fmt.Fprint(w,"hello guest")
    } else {
        fmt.Fprint(w,"hello ", u.UserNicename)
    }
}

func Regist(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

    udata :=xorm.WpUsers{UserLogin:"luv",UserPass:"331234958", UserNicename:"abner",UserEmail:"331234958@qq.com", UserUrl:"http://www.web012.com", UserRegistered:time.Now(),UserActivationKey:"#aadf@$",UserStatus:0,DisplayName:"小辉"}
    id,e := models.AddUser(udata)
    if e==nil {
        fmt.Fprint(w,"add success id",id)
    } else {
        fmt.Fprint(w,"add failed")
    }
}


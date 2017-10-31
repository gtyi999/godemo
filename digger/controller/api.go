package controller

import (
    "net/http"
    "github.com/julienschmidt/httprouter"
)

//api接口

//指定手机号码发送SMS
func ApiSendSms(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}

//生成验证码
func ApiGenVercode(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}

//检查账号是否存在
func ApiCheckAcct(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}

//检查手机号是否存在 {"ok":false}
func ApiCheckMobile(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}









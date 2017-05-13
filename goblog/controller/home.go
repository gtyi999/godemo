package controller

import (
    "net/http"
    "html/template"
    "github.com/julienschmidt/httprouter"
    "fmt"
)

//api接口





// 网站首页
func HomeIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    tplpath := ThemePath + "index.tpl"
    fmt.Println(tplpath)
    tmpl, _ := template.ParseFiles(tplpath)
    datamap := make(map[string]interface{})
    datamap["Name"] = "liuhui"
    tmpl.Execute(w, datamap)
}

// 分类列表页面
func HomeCatory(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}

// 文章页面
func HomeArticle(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}

// 编辑页面
func HomeWrite(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}

package controller

import (
    "net/http"
    "html/template"
    "github.com/julienschmidt/httprouter"
    "github.com/qqqc/godemo/goblog/models"
    "fmt"
)

//api接口


func mytest(in string)(out string, e error) {
    return "hello " + in, e
}


// 网站首页
func HomeIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    tplpath := ThemePath + "index.html"
    tmpl, _ := template.ParseFiles(tplpath)
    tmpl = tmpl.Funcs(template.FuncMap{"mytest":mytest})

    datamap := make(map[string]interface{})

    //生成测试数据
    /*
    wppost := models.NewWpPosts()
    wppost.PostAuthor = 2
    wppost.PostContent = "这里是很长的文章内容"
    wppost.PostTitle   = "文章标题测试"
    wppost.PostExcerpt   = "这里是文章摘要"
    wppost.PostStatus    = "publish"
    wppost.CommentStatus = "closed"
    wppost.PostName      = "WhoVimDo"
    wppost.Guid         = "23"
    wppost.PostType     = "context" // context文章 mediea多媒体
    wppost.CommentCount  = 0
    wppost.PostModified = time.Now()
    wppost.PostDate = time.Now()
    postid,e := wppost.Add()
    if e!=nil {
        datamap["retcode"] = "1"
        datamap["retmsg"] = "发布失败:" + e.Error()
        tmpl.Execute(w, datamap)
    } else {
        datamap["retcode"] = "1"
        datamap["retmsg"] = "发布成功" + "影响行数:" +strconv.FormatInt(postid,10) + " !"
        tmpl.Execute(w, datamap)
    }
    */
    posts := models.NewWpPosts()
    postlist, e := posts.GetArticles(4)
    if e!=nil {
        datamap["retcode"] = "1"
        datamap["retmsg"] = "查询失败:" + e.Error()
        tmpl.Execute(w, datamap)
    } else {
        datamap["retcode"] = "0"
        datamap["retmsg"] = "查询成功:"
        datamap["postlist"] = postlist
        fmt.Println(postlist)
        //fmt.Println("datamap:",datamap)
        tmpl.Execute(w, datamap)
    }



    //显示文章数据
}

// 关于页面
func HomeAbout(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    tplpath := ThemePath + "about.html"
    fmt.Println(tplpath)
    tmpl, _ := template.ParseFiles(tplpath)
    datamap := make(map[string]interface{})
    datamap["Name"] = "liuhui"
    tmpl.Execute(w, datamap)
}

// 联系人页面
func HomeContact(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    tplpath := ThemePath + "contact.html"
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

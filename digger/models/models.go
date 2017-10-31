package models

import (
    "fmt"
    "github.com/go-xorm/xorm"
    _ "github.com/go-sql-driver/mysql"
    "log"
)

var APPDB *xorm.Engine = nil
var err error

func init() {
    APPDB, err = xorm.NewEngine("mysql","root:111111@tcp(172.17.0.2:3306)/wpdb?charset=utf8")
    if err!=nil {
        log.Fatal(err.Error())
    } else {
        //创建表的时候下面才需要
        //e := APPDB.Sync(new(M.WpUsers))
        //if e!=nil {
        //	log.Fatal("create table failed...")
        //} else {
        //	fmt.Println("create table success")
        //}
    }

    //newid, err := APPDB.Insert(&M.WpUsers{UserLogin:"luv",UserPass:"331234958", UserNicename:"abner",UserEmail:"331234958@qq.com",
    //	UserUrl:"http://www.web012.com", UserRegistered:time.Now(),UserActivationKey:"#aadf@$",UserStatus:0,DisplayName:"小辉"})
    //if err!=nil {
    //	log.Fatal("insert failed...",err.Error())
    //}
    //fmt.Println("insert success ...",newid)

}

func NewWpPosts() *WpPosts {
    return &WpPosts{}
}
func (this *WpPosts)Add()(irows int64, err error) {
    irows,err = APPDB.Insert(this)
    return
}
func (this *WpPosts)GetArticles(nums int)(articles []*WpPosts, err error) {


    e := APPDB.Limit(nums).Iterate(new(WpPosts), func(idx int, bean interface{}) error {
        articles = append(articles,bean.(*WpPosts) )
        fmt.Println(bean.(*WpPosts).Id)
        //articles[idx] = bean.(*WpPosts)
        return nil
    })

    if e!=nil {
        fmt.Println("Get function")
        return nil, e
    } else {
        return
    }
}








func AddUser(users WpUsers)(id int64, err error) {
    id,err = APPDB.Insert(users)
    return
}

func GetUserById(id int64)*WpUsers {
    var user WpUsers
    if ok,e:=APPDB.ID(id).Get(&user); e!=nil{
        return nil
    } else {
        if ok {
            return &user
        } else {
            fmt.Println("不存在...id:",id)
        }
    }
    return nil
}


package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	M "github.com/qqqc/godemo/goblog/models/xorm"
	"log"
	"fmt"
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

func AddUser(users M.WpUsers)(id int64, err error) {
	id,err = APPDB.Insert(users)
	return
}

func GetUserById(id int64)*M.WpUsers {
	var user M.WpUsers
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







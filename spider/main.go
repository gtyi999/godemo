package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "strings"
    "strconv"
    "container/list"
    "regexp"
    "os"
)

//获取指定URL的html内容
//1start add by guanty 20170324
//2start add by guanty 20170324
func getPageContent(url string) string {
    //检查url合法性
    //获取url
    var resp *http.Response
    var err error
    for{
        resp, err = http.Get(url)
        if err != nil {
            continue
        }
        if(resp.StatusCode!=200) {
            continue
        }
        break
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        // handle error
    }
    //fmt.Println(body)
    return (string(body))
}

//获取内容中 开始为strstart...strend中间的内容,1-只获取一次 0-获取最多次 N-获取指定次数
func splitData(content string, start string, end string, itype int) string {
    if itype==1 {
        firstpos := strings.Index(content,start)
        firstpos = firstpos + len(start)
        datalen := strings.Index(content[firstpos:],end)
        if datalen<0 {
            return ""
        }
        endpos := firstpos + datalen
        //fmt.Println("tmpindex:",firstpos," datalen:", datalen)
        return content[firstpos:endpos]
    }
    return ""
}

func getPagecnt(content string) int {
    return 0
}

func main() {
    rooturl := "http://www.kancloud.cn/explore"
    fmt.Println(rooturl)

    content := getPageContent(rooturl)
    //fmt.Println(content)
    //<a class="end" href="/explore?page=73">73</a>
    fmt.Println(splitData(content,"<a class=\"end\" href=\"/explore?page=","\"",1))
    //pagecnts := splitData(content,"<a class=\"end\" href=\"/explore?page=","\"",1)
    bookurl := "http://www.kancloud.cn/explore?page="
    pagecnts,_ := strconv.Atoi(splitData(content,"<a class=\"end\" href=\"/explore?page=","\"",1))

    itemlist := list.New()
    pagelist := list.New()
    for index:=1; index<=pagecnts; index++ {
        //fmt.Println(bookurl,index)
        listcontent := getPageContent(bookurl+strconv.Itoa(index))
        reg := regexp.MustCompile("<a href=\".*\" target=\"_blank\" class=\"name\"")
        listurl := reg.FindAllString(listcontent, -1)
        //fmt.Println(listcontent)
        //获取每张页面内容中书籍
        for _,val := range listurl {
            tmp := splitData(val,"\"","\"",1)
            //fmt.Println(tmp)
            itemlist.PushBack("http://www.kancloud.cn"+tmp)
        }
    }

    for e := itemlist.Front(); e != nil; e = e.Next() {
        //fmt.Println(e.Value)
        //开始检查是否免费，如果免费找到起始页面
        datacontent := getPageContent( e.Value.(string))

        tmpurl,ok := e.Value.(string)
        if ok{
            if strings.Index(datacontent,"<i class=\"icon icon-assignment\"></i> 试读</b>") >0 {
                fmt.Println("试读："+tmpurl)
                continue
            } else {

                myreg := regexp.MustCompile("<a href=\""+tmpurl+".*\" class=\"e-data data-split w-btn btn-l btn-yellow\"")
                firtpage := myreg.FindAllString(datacontent,-1)

                if firtpage==nil {
                    continue
                }
                url := firtpage[0][len("<a href=\"")+len(tmpurl)+1:]
                endpos := strings.Index(url,"\"")
                //fmt.Println(url[0:endpos])
                fmt.Println(tmpurl+"/"+url[0:endpos])
                pagelist.PushBack(tmpurl+"/"+url[0:endpos])

                //tmp := splitData(datacontent,"<li class=\"navg-item recommend \">","\" class=\"e-data data-split w-btn btn-l btn-yellow\" target=\"_self\">",1)
                //fmt.Println(tmp)

            }
        }
    }

    file1, err := os.Create("url.txt")
    for e := pagelist.Front(); e != nil; e = e.Next() {
        defer file1.Close()
        if err != nil {
            fmt.Println(file1, err)
            return
        }
        data,ok := e.Value.(string)
        if ok {
            file1.WriteString(data+"\r\n")
        }

    }
    file1.Close()


}

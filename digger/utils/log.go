package utils

import (
    log "github.com/cihub/seelog"
    "sync"
    "fmt"
)


var applog log.LoggerInterface
var mutex sync.Mutex

func NewAppLog() log.LoggerInterface{
    mutex.Lock()

    if applog ==nil {
        fmt.Println("applog is nil")

        loghandler, err := log.LoggerFromConfigAsFile("D:/goproject/src/github.com/qqqc/godemo/goblog/conf/seelog.xml")
        if err != nil {
            panic("read log config file failed! error:" + err.Error())
        }
        log.ReplaceLogger(loghandler)
    } else {
        fmt.Println("applog is not nil")
    }
    mutex.Unlock()
    applog = log.Current
    return log.Current
}


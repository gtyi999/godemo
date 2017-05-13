package controller

import (
    log "github.com/cihub/seelog"
)

var ThemePath string = "/home/luv/goproject/src/github.com/qqqc/godemo/goblog/views/static/"

func init() {
    log.Debug("控制器加载...")
    log.Debug("网页模板加载...")

}

type BaseController interface {
}

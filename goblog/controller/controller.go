package controller

import (
    log "github.com/cihub/seelog"
    "path/filepath"
)

var ThemePath string = "/home/luv/goproject/src/github.com/qqqc/godemo/goblog/views/theme/default/"
var TplFiles []string

func init() {
    log.Debug("控制器加载...")
    log.Debug("网页模板加载...")


    TplFiles = []string{
        filepath.Join(ThemePath,"index.html"),
    }

}

type BaseController interface {
}

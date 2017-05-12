package controller

import (
    log "github.com/cihub/seelog"
)

func init() {
    log.Debug("控制器加载...")
}

type BaseController interface {
}

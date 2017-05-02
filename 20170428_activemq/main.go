package main

import (
    "runtime"
    "./publish"
    //"./subscribe"

)



func main() {
    // 单生产者+事务+持久化
    publish.Run(runtime.NumCPU() * 2,"publish.log","192.168.10.166","61613","/queue/liuhui_test","1.0",false, true,1000000,1)
    //subscribe.Run()
}

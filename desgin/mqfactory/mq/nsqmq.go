package mq

import (
    "github.com/nsqio/go-nsq"
    "fmt"
)

type NsqMq struct {
    BaseMq
}


//连接MQ 子类实现
func (this *NsqMq)Connect(host,port string, opts []string)(conn MQCONN, err error) {
    nsqserver := host + ":" + port
    p, err := nsq.NewProducer(nsqserver, nsq.NewConfig())
    fmt.Println("connect to ",p.String(), " success")
    return nil,nil
}

//断开MQ 子类实现
func (this *NsqMq)Disconnect()(err error) {

    return nil
}

//获取一条消息 子类实现
func (this *NsqMq)GetMessage(queue string)(msg MQMSG, err error) {
    return nil,nil
}

//发布一条消息一条消息 子类实现
func (this *NsqMq)PutMessage(queue string, msg MQMSG)(err error) {
    return nil
}

//循环处理消息 子类实现
func (this *NsqMq)OnMessage() {

}
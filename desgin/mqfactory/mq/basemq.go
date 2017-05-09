package mq

type MQCONN interface{}
type MQCLIENT interface{}
type MQMSG  interface{}


//接口
type IBaseMq interface {
    Connect(host,port string, opts []string)(conn MQCONN, err error)
    Disconnect()(err error)
    GetMessage(queue string)(msg MQMSG, err error)
    PutMessage(queue string, msg MQMSG)(err error)
    OnMessage()
}

type BaseMq struct {
    //Host string
    //Port string
    Conn MQCONN
    MqClient MQCLIENT
    Err  error
}

//创建实例 子类实现

//连接MQ 子类实现
func (this *BaseMq)Connect(host,port string, opts []string)(conn MQCONN, err error) {
    return nil,nil
}

//断开MQ 子类实现
func (this *BaseMq)Disconnect()(err error) {
    return nil
}

//获取一条消息 子类实现
func (this *BaseMq)GetMessage(queue string)(msg MQMSG, err error) {
    return nil,nil
}

//发布一条消息一条消息 子类实现
func (this *BaseMq)PutMessage(queue string, msg MQMSG)(err error) {
    return nil
}

//循环处理消息 子类实现
func (this *BaseMq)OnMessage() {

}


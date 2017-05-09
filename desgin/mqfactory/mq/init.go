package mq

type MqFactory struct {
}

func (this *MqFactory)MakeMq(name string)(basemq IBaseMq) {
    if name=="nsq" {
        basemq = new(NsqMq)
    } else if name == "rabbitmq" {
        basemq = new()
    }
    return
}

func init() {


}

package main

var (
	cfg *Cfg
)

var success = map[string]int{"code": 1}
var failure = map[string]int{"code": 0}

type Consumer struct {
	way Way
}

func (consumer *Consumer) setWay(way Way) {
	consumer.way = way
}

func (consumer *Consumer) Send(message Message) interface{} {
	return consumer.way.ConsumeMsg(message)
}

type Way interface {
	ConsumeMsg(message Message) interface{}
}

type Cfg struct {
	HttpPort string
	Scs      []Sc
	QqMails  []QqMail
}

type Message struct {
	Title    string
	Content  string
	Original string
	MailName string
}

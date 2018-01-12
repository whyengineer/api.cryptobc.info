package caculate

import(
	"golang.org/x/sync/syncmap"
	"log"
)

type Mq struct{
	subtopic syncmap.Map
	pubtopic syncmap.Map
}

func NewMq() *Mq{
	m:=new(Mq)
	return m
}
func (m *Mq)Pub(topic string)chan StaInfo{
	
}
func (m *Mq)Sub(topic string)chan StaInfo{
	a:=make(chan StaInfo)
	cl,ok:=m.subtopic.Load(topic)
	acl,ok:=cl.([]chan StaInfo)
	if ok{
		if ok {
			acl=append(acl,a)
			m.subtopic.Store(topic,acl)
		}else{
			b:=[]chan StaInfo{a}
			m.subtopic.Store(topic,b)
		}
	}else{
		log.Panic("assert failed")
		
	}
	return a
}
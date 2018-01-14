package caculate

import (
	"log"
	"testing"
	"time"
)

func Test_mq(t *testing.T) {
	mq := NewMq()
	go func() {
		a := make(chan StaInfo)
		mq.Sub("3", a)
		for {
			b := <-a
			log.Println(b)
			log.Println("rt 3")
		}
	}()
	go func() {
		a := make(chan StaInfo)
		err := mq.Sub("4", a)
		if err != nil {
			log.Println(err)
		}
		for {
			b := <-a
			log.Println(b)
			log.Println("rt 4")
		}
	}()
	c := StaInfo{}
	c.BuyAmount = 300
	log.Println("pub")
	mq.Pub(c)
	time.Sleep(3 * time.Second)
	c.BuyAmount = 400
	mq.Pub(c)
	time.Sleep(3 * time.Second)
}

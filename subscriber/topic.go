package subscriber

import (
	"errors"
	"log"
	"sync"
)

type Topic struct {
	stopCh chan struct{}
	// 可以用链表实现， 减少内存的分配
	// 或者是用游标记录已经消费的下表 定时清理已经消费的元素并进行拷贝
	msg     []Message
	subs    []*Subscriber
	subRwMu sync.RWMutex
	msgRwMu sync.RWMutex
}

func NewTopic() *Topic {
	t := &Topic{
		subRwMu: sync.RWMutex{},
		msgRwMu: sync.RWMutex{},
		stopCh:  make(chan struct{}),
	}

	t.run()

	return t
}

func (t *Topic) removeSubscriber(index int) {
	t.subRwMu.Lock()
	defer t.subRwMu.Unlock()
	t.subs = append(t.subs[:index], t.subs[index+1:]...)
}

func (t *Topic) dispatch(msg Message) {

	t.subRwMu.RLock()
	subscribers := t.subs
	t.subRwMu.Unlock()

	for index, sub := range subscribers {
		go func(ch *Subscriber, index int) {
			// 这个订阅者不再有消费能力
			defer func() {
				if err := recover(); err != nil {
					log.Printf("msg send err Subscriber #%d not working； err: %s \n", index, err)
					t.removeSubscriber(index)
				}
			}()
			*ch <- msg
		}(sub, index)
	}
}

func (t *Topic) run() {
	ch := t.getMessage()

	go func() {
	Loop:
		for {

			select {
			case <-t.stopCh:
				break Loop
			case msg := <-*ch:
				t.dispatch(msg)
			default:
			}
		}
	}()

}

func (t *Topic) getMessage() *chan Message {
	ch := make(chan Message)
	go func() {
	Loop:
		for {
			select {
			case <-t.stopCh:
				break Loop
			default:
			}
			message, err := t.popMessage()
			if err == nil {
				ch <- message
			}

		}
	}()
	return &ch
}

func (t *Topic) popMessage() (Message, error) {
	t.msgRwMu.Lock()
	defer t.msgRwMu.Unlock()
	if len(t.msg) > 0 {
		msg := t.msg[0]
		t.msg = t.msg[1:]
		return msg, nil
	}

	return nil, errors.New("empty")
}

func (t *Topic) Publish(msg Message) {
	t.msgRwMu.Lock()
	defer t.msgRwMu.Unlock()
	t.msg = append(t.msg, msg)
}

func (t *Topic) Subscribe() *Subscriber {
	s := make(Subscriber)
	t.subRwMu.Lock()
	t.subs = append(t.subs, &s)
	t.subRwMu.Unlock()
	return &s
}

package engine

import "sync"

type MessageQueue struct {
	sync.Mutex
	c []Command
}

func (q *MessageQueue) push(command Command) {
	q.Lock()
	q.c = append(q.c, command)
	q.Unlock()
}

func (q *MessageQueue) pull() Command {
	q.Lock()
	cmd := q.c[0]
	q.c[0] = nil
	q.c = q.c[1:]
	q.Unlock()
	return cmd
}

func (q *MessageQueue) isEmpty() bool {
	return !(len(q.c) > 0)
} 

type EventLoop struct {
	queue *MessageQueue
	stopConfirm chan struct{}
	stopRequest bool
}

func (el *EventLoop) Start() {
	el.queue = &MessageQueue{}
	el.stopConfirm = make(chan struct{}, 1)
	el.stopRequest = false
	go func() {
		for {
		  if (el.queue.isEmpty() && el.stopRequest) {
				el.stopConfirm <- struct{}{}
				return
			} else if (!el.queue.isEmpty()) {
				cmd := el.queue.pull()
				cmd.Execute(el)
			}
		}
	}()
}

func (el *EventLoop) Post(command Command) {
	if el.stopRequest {
		panic("queue is already closed.")
	}
	el.queue.push(command)
}

func (el *EventLoop) AwaitFinish() {
	el.stopRequest = true
	<-el.stopConfirm
}

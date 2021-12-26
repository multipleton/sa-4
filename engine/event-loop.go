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

type EventLoop struct {
	queue *MessageQueue
}


func (el *EventLoop) Start() {
	el.queue = &MessageQueue{}
	
	// implement execution of commands
	
}

func (el *EventLoop) Post(command Command) {
	el.queue.push(command)
}

func (el *EventLoop) AwaitFinish() {
	// implement method
}

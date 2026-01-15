package ws

import "sync"

type Subscriber chan any

type PubSub struct {
	subscribers map[string][]Subscriber
	mu          sync.Mutex
}

func NewPubSub() *PubSub {
	return &PubSub{subscribers: make(map[string][]Subscriber)}
}

func (ps *PubSub) Subscribe(topic string) Subscriber {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	ch := make(Subscriber, 1)
	ps.subscribers[topic] = append(ps.subscribers[topic], ch)
	return ch
}

func (ps *PubSub) Publish(topic string, msg any) {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	for _, sub := range ps.subscribers[topic] {
		sub <- msg
	}
}

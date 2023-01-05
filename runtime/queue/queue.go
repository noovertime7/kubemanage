package queue

import (
	"fmt"
)

// Queue integrates all data events to one seqential queue
type Queue interface {
	// Push specified event to local queue
	Push(e *Event)
	// Get event from queue, blocked
	Get() (*Event, error)
	// AGet async get event from queue, not blocked
	AGet() (*Event, error)
	// GetChannel event reading queue
	GetChannel() (<-chan *Event, error)
	// Close close Queue
	Close()
	// IsClosed Check queue IsClosed
	IsClosed() bool
}

// NewQueue create default Queue for local usage
func NewQueue() Queue {
	return &channelQueue{
		localQ: make(chan *Event, 128),
	}
}

// channelQueue default queue using channel
type channelQueue struct {
	localQ chan *Event
}

// Push specified event to local queue
func (cq *channelQueue) Push(e *Event) {
	if e != nil {
		cq.localQ <- e
	}
}

// Get event from queue
func (cq *channelQueue) Get() (*Event, error) {
	e, ok := <-cq.localQ
	if ok {
		return e, nil
	}
	return nil, fmt.Errorf("queue closed")
}

// AGet async get event from queue, not blocked
func (cq *channelQueue) AGet() (*Event, error) {
	select {
	case e, ok := <-cq.localQ:
		if ok {
			return e, nil
		}
		return nil, fmt.Errorf("queue closed")
	default:
		return nil, nil
	}
}

// GetChannel event reading queue
func (cq *channelQueue) GetChannel() (<-chan *Event, error) {
	if cq.localQ == nil {
		return nil, fmt.Errorf("lost event queue")
	}
	return cq.localQ, nil
}

// Close event queue
func (cq *channelQueue) Close() {
	close(cq.localQ)
}

// IsClosed check queue is closed or not
func (cq *channelQueue) IsClosed() bool {
	select {
	case data, received := <-cq.localQ:
		// 使用读判断channel是否关闭，如果未关闭需要将消息再写回去，否则会被判断消费掉这条消息
		if received {
			cq.localQ <- data
		}
		return !received
	default:
	}
	return false
}

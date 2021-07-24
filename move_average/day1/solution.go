package day1

import "container/list"

type MovingAvg struct {
	size  int
	sum   int
	queue *list.List
}

func New(size int) MovingAvg {
	return MovingAvg{
		size:  size,
		sum:   0,
		queue: list.New(),
	}
}

func (m *MovingAvg) Next(n int) float64 {
	if m.queue.Len() == m.size {
		front := m.queue.Front()
		m.sum -= front.Value.(int)
		m.queue.Remove(front)
	}

	m.queue.PushBack(n)
	m.sum += n
	return float64(m.sum / m.queue.Len())
}

package sort

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrBufferFull  = errors.New("buffer is full")
	ErrBufferEmpty = errors.New("buffer is empty")
)

type ringBuffer[T any] struct {
	data []T
	size int
	rp   int
	wp   int
}

func NewRingBuffer[T any](size int) *ringBuffer[T] {
	return &ringBuffer[T]{
		data: make([]T, size),
		size: size,
	}
}

func (rb *ringBuffer[T]) enqueue(value T) (int, error) {
	rb.data[rb.wp] = value
	if rb.wpRound() {
		rb.wp = 0
	} else {
		rb.wp++
	}
	return rb.wp, nil
}

func (rb *ringBuffer[T]) dequeue() (any, error) {
	if rb.isEmpty() {
		return nil, fmt.Errorf("failed to dequeue: %w", ErrBufferFull)
	}
	v := rb.data[rb.rp]
	if rb.rpRound() {
		rb.rp = 0
	} else {
		rb.rp++
	}
	return v, nil
}

func (rb *ringBuffer[T]) isEmpty() bool {
	return len(rb.data) == 0
}

func (rb *ringBuffer[T]) wpRound() bool {
	return rb.wp == rb.size-1
}

func (rb *ringBuffer[T]) rpRound() bool {
	return rb.rp == rb.size-1
}

func (rb *ringBuffer[T]) String() string {
	var buf strings.Builder
	for i, v := range rb.data {
		fmt.Fprintf(&buf, "{id: %d, data: %v}\n", i, v)
	}
	return buf.String()
}

func (rb *ringBuffer[T]) currentPositions() {
	fmt.Printf("{readPos: %d, writePos: %d}\n", rb.rp, rb.wp)
}

package circular

import "errors"

var ErrBufferFull = errors.New("buffer is full")
var ErrBufferEmpty = errors.New("buffer is empty")

type Buffer struct {
	data     []byte
	readPos  int
	writePos int
	size     int
	capacity int
}

func NewBuffer(capacity int) *Buffer {
	return &Buffer{
		data:     make([]byte, capacity),
		capacity: capacity,
	}
}

func (b *Buffer) ReadByte() (byte, error) {
	if b.size == 0 {
		return 0, ErrBufferEmpty
	}

	c := b.data[b.readPos]
	b.readPos = (b.readPos + 1) % b.capacity
	b.size--

	return c, nil
}

func (b *Buffer) WriteByte(c byte) error {
	if b.size == b.capacity {
		return ErrBufferFull
	}

	b.data[b.writePos] = c
	b.writePos = (b.writePos + 1) % b.capacity
	b.size++
	return nil
}

func (b *Buffer) Overwrite(c byte) {
	if b.size == b.capacity {
		b.readPos = (b.readPos + 1) % b.capacity
	} else {
		b.size++
	}

	b.data[b.writePos] = c
	b.writePos = (b.writePos + 1) % b.capacity
}

func (b *Buffer) Reset() {
	b.size = 0
	b.writePos = 0
	b.readPos = 0
}

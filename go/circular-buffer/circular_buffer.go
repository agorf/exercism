package circular

import "errors"

type Buffer struct {
	values             []byte
	head, tail, length int
}

var (
	errBufferEmpty = errors.New("buffer empty")
	errBufferFull  = errors.New("buffer full")
)

func NewBuffer(capacity int) *Buffer {
	b := Buffer{head: 0, tail: 0, length: 0}
	b.values = make([]byte, capacity)
	return &b
}

func (b *Buffer) Capacity() int {
	return cap(b.values)
}

func (b *Buffer) IsEmpty() bool {
	return b.length == 0
}

func (b *Buffer) IsFull() bool {
	return b.length == b.Capacity()
}

func (b *Buffer) AdvanceHead() {
	b.head = (b.head + 1) % b.Capacity()
}

func (b *Buffer) AdvanceTail() {
	b.tail = (b.tail + 1) % b.Capacity()
}

func (b *Buffer) ReadByte() (byte, error) {
	if b.IsEmpty() {
		return 0, errBufferEmpty
	}
	value := b.values[b.head]
	b.length--
	b.AdvanceHead()
	return value, nil
}

func (b *Buffer) WriteByte(c byte) error {
	if b.IsFull() {
		return errBufferFull
	}
	b.values[b.tail] = c
	b.length++
	b.AdvanceTail()
	return nil
}

func (b *Buffer) Overwrite(c byte) {
	if b.IsFull() {
		b.values[b.head] = c
		b.AdvanceHead()
	} else {
		_ = b.WriteByte(c)
	}
}

func (b *Buffer) Reset() {
	b.head = 0
	b.tail = 0
	b.length = 0
}

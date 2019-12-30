package circular

import (
	"errors"
)

type Buffer struct {
	bytes     []byte
	next      int
	available int
	size      int
}

func (buff *Buffer) ReadByte() (byte, error) {
	if buff.available == buff.size {
		return 0, errors.New("no available data")
	}

	pos := (buff.next + buff.available) % buff.size
	buff.available += 1
	return buff.bytes[pos], nil
}

func (buff *Buffer) WriteByte(b byte) error {
	if buff.available == 0 {
		return errors.New("no more space")
	}

	buff.bytes[buff.next] = b
	buff.next = (buff.next + 1) % buff.size
	buff.available -= 1
	return nil
}

func (buff *Buffer) Overwrite(b byte) {
	buff.bytes[buff.next] = b
	buff.next = (buff.next + 1) % buff.size
	if buff.available > 0 {
		buff.available -= 1
	}
}

func NewBuffer(size int) *Buffer {
	buff := new(Buffer)
	buff.bytes = make([]byte, size)
	buff.next = 0
	buff.available = size
	buff.size = size
	return buff
}

func (buff *Buffer) Reset() {
	buff.next = 0
	buff.available = buff.size
}

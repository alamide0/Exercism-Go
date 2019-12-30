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

	pos := buff.next + buff.available
	if pos >= buff.size {
		pos -= buff.size 
	}
	buff.available += 1
	return buff.bytes[pos], nil
}

func (buff *Buffer) WriteByte(b byte) error {
	if buff.available == 0 {
		return errors.New("no more space")
	}

	buff.bytes[buff.next] = b
	buff.next = buff.next + 1
	if buff.next >= buff.size {
		buff.next -= buff.size
	}
	buff.available -= 1
	return nil
}

func (buff *Buffer) Overwrite(b byte) {
	buff.bytes[buff.next] = b
	buff.next = buff.next + 1
	if buff.next >= buff.size {
		buff.next -= buff.size
	}
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
/**
Version 1: 
pos := buff.next + buff.available
if pos >= buff.size {
	pos -= buff.size 
}
BenchmarkOverwrite-8    1000000000               1.93 ns/op     519052853364.44 MB/s           0 B/op          0 allocs/op
BenchmarkWriteRead-8    500000000                3.33 ns/op     150156769132.68 MB/s           0 B/op          0 allocs/op

Version 2:
pos := (buff.next + buff.available) % buff.size
BenchmarkOverwrite-8    100000000               13.7 ns/op      7289753300.06 MB/s             0 B/op          0 allocs/op
BenchmarkWriteRead-8    100000000               18.9 ns/op      5304322677.14 MB/s             0 B/op          0 allocs/op

so, mod will spend more time
**/

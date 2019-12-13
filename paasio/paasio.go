package paasio
import (
	"io"
	"sync"
)

type writeCounter struct {
	w io.Writer
	counter
}

type readCounter struct {
	r io.Reader
	counter
}

type counter struct {
	bytes int64
	nops int
	sync.Mutex
}

type rwCounter struct {
	WriteCounter
	ReadCounter
}

func (wc *writeCounter) WriteCount() (n int64, nops int) {
	wc.Lock()
	defer wc.Unlock()
	return wc.bytes, wc.nops
}

func (wc *writeCounter) Write(p []byte) (n int, err error) {
	n, err = wc.w.Write(p)
	wc.Lock()
	defer wc.Unlock()
	wc.bytes += int64(n)
	wc.nops++
	return
}

func (rc *readCounter) Read(p []byte)(n int, err error) {
	n, err = rc.r.Read(p)
	rc.Lock()
	defer rc.Unlock()
	rc.bytes += int64(n)
	rc.nops++
	return 
}

func (rc *readCounter) ReadCount() (n int64, nops int) {
	rc.Lock()
	defer rc.Unlock()
	return rc.bytes, rc.nops
}

func NewWriteCounter(writer io.Writer) WriteCounter {
	return &writeCounter{
		w: writer,
		counter: counter{0, 0, sync.Mutex{}},
	}
}

func NewReadCounter(reader io.Reader) ReadCounter {
	return &readCounter{
		r: reader,
		counter: counter{0, 0, sync.Mutex{}},
	}
}

func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
	return &rwCounter{
		NewWriteCounter(rw),
		NewReadCounter(rw),
	}
}


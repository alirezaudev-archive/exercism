package paasio

import (
	"io"
	"sync"
)

type readCounter struct {
	r       io.Reader
	readN   int64
	readOps int
	mu      sync.Mutex
}

type writeCounter struct {
	w        io.Writer
	writeN   int64
	writeOps int
	mu       sync.Mutex
}

type readWriteCounter struct {
	*readCounter
	*writeCounter
}

func NewWriteCounter(writer io.Writer) WriteCounter {
	return &writeCounter{w: writer}
}

func NewReadCounter(reader io.Reader) ReadCounter {
	return &readCounter{r: reader}
}

func NewReadWriteCounter(readwriter io.ReadWriter) ReadWriteCounter {
	return &readWriteCounter{
		readCounter:  &readCounter{r: readwriter},
		writeCounter: &writeCounter{w: readwriter},
	}
}

func (rc *readCounter) Read(p []byte) (int, error) {
	rc.mu.Lock()
	defer rc.mu.Unlock()

	n, err := rc.r.Read(p)
	rc.readN += int64(n)
	rc.readOps++
	return n, err
}

func (rc *readCounter) ReadCount() (int64, int) {
	rc.mu.Lock()
	defer rc.mu.Unlock()

	return rc.readN, rc.readOps
}

func (wc *writeCounter) Write(p []byte) (int, error) {
	wc.mu.Lock()
	defer wc.mu.Unlock()

	n, err := wc.w.Write(p)
	wc.writeN += int64(n)
	wc.writeOps++
	return n, err
}

func (wc *writeCounter) WriteCount() (int64, int) {
	wc.mu.Lock()
	defer wc.mu.Unlock()

	return wc.writeN, wc.writeOps
}

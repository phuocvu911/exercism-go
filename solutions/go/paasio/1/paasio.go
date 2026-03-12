package paasio

import (
    "io"
    "sync/atomic"
)

// Define readCounter and writeCounter types here.
type readCounter struct{
    reader io.Reader
    totalByte int64
    totalCall int64 // have to be int64 because atomic ppackage doesnt have int8
}
type writeCounter struct{
    writer io.Writer
    totalByte int64
    totalCall int64
}

// For the return of the function NewReadWriteCounter, you must also define a type that satisfies the ReadWriteCounter interface.
type readwriteCounter struct{
    readCounter
    writeCounter
}

func NewWriteCounter(writer io.Writer) WriteCounter { //return interface
	return &writeCounter{writer:writer}
}

func NewReadCounter(reader io.Reader) ReadCounter {
	return &readCounter{reader:reader}
}

func NewReadWriteCounter(readwriter io.ReadWriter) ReadWriteCounter {
	return &readwriteCounter{
        readCounter{reader:readwriter},
        writeCounter{writer:readwriter},
    }
}

func (rc *readCounter) Read(p []byte) (int, error) {
	n, err:= rc.reader.Read(p)
    atomic.AddInt64(&rc.totalByte, int64(n))
    atomic.AddInt64(&rc.totalCall, 1)
    return n, err
}

func (rc *readCounter) ReadCount() (int64, int) {
	return atomic.LoadInt64(&rc.totalByte), int(atomic.LoadInt64(&rc.totalCall))
}

func (wc *writeCounter) Write(p []byte) (int, error) {
	n, err:= wc.writer.Write(p)
    atomic.AddInt64(&wc.totalByte, int64(n))
    atomic.AddInt64(&wc.totalCall, 1)
    return n, err
}

func (wc *writeCounter) WriteCount() (int64, int) {
	return atomic.LoadInt64(&wc.totalByte), int(atomic.LoadInt64(&wc.totalCall))
}

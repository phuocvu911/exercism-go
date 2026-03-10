package circular
import "errors"

// Define the Buffer type here.
type Buffer struct{
    data []byte
    size int
    head int
    tail int
    count int //capacity, how full is our buffer
}

func NewBuffer(size int) *Buffer {
	return &Buffer{data: make([]byte, size), size: size}
}
//err when empty
func (b *Buffer) ReadByte() (byte, error) {
	if b.count == 0{
        return 0, errors.New("Buffer is empty")
    }
    
    res:= b.data[b.head]
    b.head = (b.head +1)%b.size // circular
    b.count--
    return res, nil
}
//err when full
func (b *Buffer) WriteByte(c byte) error {
	if b.count == b.size {
        return errors.New("Buffer is full")
    }
    b.data[b.tail] = c
    b.tail = (b.tail + 1)%b.size
    b.count++
	return nil
}

func (b *Buffer) Overwrite(c byte) {
    if b.count < b.size{ //can call Overwrite on non-full buffer
        b.WriteByte(c)
        return
    }
	b.data[b.tail] = c
    b.head = (b.head +1)%b.size
    b.tail = (b.tail + 1)%b.size
}

func (b *Buffer) Reset() {
	b.head =0
    b.tail=0
    b.count =0
}

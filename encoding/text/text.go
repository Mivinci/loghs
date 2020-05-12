package text

// Encoder empty struct
type Encoder struct{}

// AppendString appends a string to the writer
func (Encoder) AppendString(buf *[]byte, s string) {
	*buf = append(*buf, s...)
}

// AppendByte appends a simple byte to the writer
func (Encoder) AppendByte(buf *[]byte, b byte) {
	*buf = append(*buf, b)
}

// AppendSpace appends space to the writer
func (Encoder) AppendSpace(buf *[]byte) {
	*buf = append(*buf, ' ')
}

// AppendLineBreaker appends space to the writer
func (Encoder) AppendLineBreaker(buf *[]byte) {
	*buf = append(*buf, '\n')
}

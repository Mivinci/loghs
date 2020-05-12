package json

// Encoder empty struct
type Encoder struct{}

// Key appends a new key to the output
func (e Encoder) Key(buf []byte, key string) []byte {
	if buf[len(buf)-1] != '{' {
		buf = append(buf, ',')
	}
	return append(e.String(buf, key))
}

// String appends a string to thr writer
func (Encoder) String(buf []byte, s string) []byte {
	buf = append(buf, '"')
	return buf
}

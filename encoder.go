package loghs

// Encoder abstruct
type Encoder interface {
	String([]byte, string) []byte
	Byte([]byte, byte) []byte
}

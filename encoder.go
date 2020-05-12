package loghs

import "github.com/mivinci/loghs/encoding/text"

type encoder interface {
	AppendString(*[]byte, string)
	AppendByte(*[]byte, byte)
}

var (
	_   encoder = (*text.Encoder)(nil)
	enc         = text.Encoder{}
)

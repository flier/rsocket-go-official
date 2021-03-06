package rsocket

import (
	"encoding/binary"
	"fmt"
	"io"
)

var defaultVersion Version = [2]uint16{1, 0}

type Version [2]uint16

func (p Version) Bytes() []byte {
	bs := make([]byte, 4)
	binary.BigEndian.PutUint16(bs, p[0])
	binary.BigEndian.PutUint16(bs[2:], p[1])
	return bs
}

func (p Version) WriteTo(w io.Writer) (n int64, err error) {
	var wrote int
	wrote, err = w.Write(p.Bytes())
	if err == nil {
		n += int64(wrote)
	}
	return
}

func (p Version) Major() uint16 {
	return p[0]
}

func (p Version) Minor() uint16 {
	return p[1]
}

func (p Version) String() string {
	return fmt.Sprintf("%d.%d", p[0], p[1])
}

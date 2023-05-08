package mandd

import (
	"fmt"
	"io"
	"os"
)

func Copy(out *os.File, in *os.File, off int, lim int) error {
	_, err := in.Seek(int64(off), 0)
	if err != nil {
		return err
	}
	var b []byte
	if lim == 0 {
		b = make([]byte, 1024*1024)
		lim = 1e+10
	} else {
		b = make([]byte, lim)
	}
	var offset int
	for offset < int(lim) {
		n, err := in.Read(b[offset:])
		offset += n
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
	}
	_, err = out.Write(b[:offset])
	if err != nil {
		return fmt.Errorf("%s, cannot write file: %s", err.Error(), in.Name())
	}
	return nil
}

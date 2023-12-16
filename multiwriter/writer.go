package multiwriter

import (
	"io"
)

type MultiWriter struct {
	w []io.Writer
}

func New(w ...io.Writer) io.Writer {
	if len(w) == 0 {
		return nil
	}
	return &MultiWriter{
		w: w,
	}
}

func (s *MultiWriter) Write(p []byte) (int, error) {
	for _, w := range s.w {
		if n, err := w.Write(p); err != nil {
			return n, err
		}
	}
	return len(p), nil
}

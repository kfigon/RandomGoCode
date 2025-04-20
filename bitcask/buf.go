package main

import (
	"bytes"
	"fmt"
	"io"
)

type MemBuffer struct {
	buf  *bytes.Buffer
	pos  int
}

// NewMemBuffer initializes a new MemBuffer
func NewMemBuffer() *MemBuffer {
	return &MemBuffer{
		buf: bytes.NewBuffer(nil),
	}
}

// Write writes data to the buffer at the current position
func (m *MemBuffer) Write(p []byte) (n int, err error) {
	if m.pos > m.buf.Len() {
		// Extend buffer if seeking beyond current length
		padding := make([]byte, m.pos-m.buf.Len())
		m.buf.Write(padding)
	}

	data := m.buf.Bytes()
	if m.pos+len(p) > len(data) {
		m.buf.Write(p)
		m.pos += len(p)
	} else {
		copy(data[m.pos:], p)
		m.pos += len(p)
	}
	return len(p), nil
}

// Read reads from the buffer at the current position
func (m *MemBuffer) Read(p []byte) (n int, err error) {
	if m.pos >= m.buf.Len() {
		return 0, io.EOF
	}

	n = copy(p, m.buf.Bytes()[m.pos:])
	m.pos += n
	return n, nil
}

// Seek sets the current read/write position
func (m *MemBuffer) Seek(offset int64, whence int) (int64, error) {
	var newPos int
	switch whence {
	case io.SeekStart:
		newPos = int(offset)
	case io.SeekCurrent:
		newPos = m.pos + int(offset)
	case io.SeekEnd:
		newPos = m.buf.Len() + int(offset)
	default:
		return 0, fmt.Errorf("invalid whence")
	}

	m.pos = newPos
	return int64(newPos), nil
}
// Copyright 2018 The go-hep Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rootio

import(
	"io"
)
type wbuff struct {
	p []byte // buffer of data to write on
	c int    // current position in buffer of data
}

func (w *wbuff) Write(p []byte) (int, error) {
	if w.c >= len(w.p) {
		return 0, io.EOF
	}
	n := copy(w.p[w.c:], p)
	w.c += n
	return n, nil
}

func (w *wbuff) WriteByte(p byte) error {
	if w.c >= len(w.p) {
		return io.EOF
	}
	w.p[w.c] = p
	w.c++
	return nil
}




// WBuffer is a write-only ROOT buffer for streaming.
type WBuffer struct {
	w      *wbuff
	err    error
	offset uint32
	refs   map[int64]interface{}
	sictx  StreamerInfoContext
}






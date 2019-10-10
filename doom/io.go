package doom

import (
	"context"
	"io"
)

type ioreturn struct {
	n   int
	err error
}

type CtxReader struct {
	ctx context.Context
	r   io.Reader
}

func NewReader(ctx context.Context, r io.Reader) *CtxReader {
	return &CtxReader{ctx: ctx, r: r}
}

func (c *CtxReader) Read(buf []byte) (int, error) {
	buf2 := make([]byte, len(buf))
	ch := make(chan ioreturn, 1)

	go func() {
		n, err := c.r.Read(buf2)
		ch <- ioreturn{n: n, err: err}

		close(ch)
	}()

	select {
	case ret := <-ch:
		copy(buf, buf2)
		return ret.n, ret.err
	case <-c.ctx.Done():
		return 0, c.ctx.Err()
	}
}

type CtxWriter struct {
	ctx context.Context
	w   io.Writer
}

func NewWriter(ctx context.Context, w io.Writer) *CtxWriter {
	return &CtxWriter{ctx: ctx, w: w}
}

func (c *CtxWriter) Write(buf []byte) (int, error) {
	buf2 := make([]byte, len(buf))
	copy(buf2, buf)
	ch := make(chan ioreturn, 1)

	go func() {
		n, err := c.w.Write(buf2)
		ch <- ioreturn{n: n, err: err}
		close(ch)
	}()

	select {
	case r := <-ch:
		return r.n, r.err
	case <-c.ctx.Done():
		return 0, c.ctx.Err()
	}
}

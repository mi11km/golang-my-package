package trace

import (
	"fmt"
	"io"
	"log"
)

// Tracer コード内での出来事を記録できるオブジェクトを表すインターフェイス
type Tracer interface {
	Trace(...interface{})
}

func New(buf io.Writer) Tracer {
	return &tracer{out: buf}
}

type tracer struct {
	out io.Writer
}

func (t *tracer) Trace(a ...interface{}) {
	_, err := t.out.Write([]byte(fmt.Sprint(a...)))
	if err != nil {
		log.Println(err)
	}
	_, err = t.out.Write([]byte("\n"))
	if err != nil {
		log.Println(err)
	}
}

type nilTracer struct{}

func (n *nilTracer) Trace(_ ...interface{}) {}
func Off() Tracer {
	return &nilTracer{}
}

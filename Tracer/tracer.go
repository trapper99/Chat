package tracer

import "io"

// Tracer is the interface that describes an object capable of tracing events throughout code.
type Tracer interface {
	Tracer(...interface{})
}

type tracerImpl struct {
	out io.Writer
}

func New(out io.Writer) *tracerImpl {
	return &tracerImpl{out: out}
}
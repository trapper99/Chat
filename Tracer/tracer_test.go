package tracer

import (
	"bytes"
	"fmt"
	"testing"
)
func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Error("Return value should not be nil.")
	} else {
		tracer.Tracer("Hello tracer package.")
		if buf.String() != "Hello tracer package.\n" {
			t.Errorf("Tracer should not write '%s'.", buf.String())
		}
	}
}

func (t *tracerImpl) Tracer(args ...interface{}) {
	fmt.Fprintln(t.out, args...)
}
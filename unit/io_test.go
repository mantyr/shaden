package unit

import (
	"testing"

	"github.com/brettbuddin/shaden/dsp"
	"github.com/stretchr/testify/require"
)

func TestExposeIn(t *testing.T) {
	io := NewIO("example", frameSize)
	in := io.NewIn("x", dsp.Float64(1))
	require.Equal(t, in, io.In["x"])
}

func TestExposeOut(t *testing.T) {
	io := NewIO("example", frameSize)
	out := io.NewOut("x")
	require.Equal(t, out, io.Out["x"])
}

type output struct {
	out  *Out
	proc func(int)
}

func (o output) Out() *Out {
	return o.out
}

func (o output) ProcessSample(i int) {
	o.proc(i)
}

func (o output) ProcessFrame(n int) {
	for i := 0; i < n; i++ {
		o.ProcessSample(i)
	}
}

func TestExposeOutProcessor(t *testing.T) {
	io := NewIO("example", frameSize)

	var called bool
	io.ExposeOutputProcessor(output{
		out: NewOut("x", make([]float64, frameSize)),
		proc: func(n int) {
			called = true
		},
	})
	out := io.Out["x"]
	require.NotNil(t, out)

	out.(SampleProcessor).ProcessSample(1)
	require.True(t, called)
}

func TestExposeProp(t *testing.T) {
	io := NewIO("example", frameSize)
	p := io.NewProp("x", 1, nil)
	require.Equal(t, p, io.Prop["x"])
}

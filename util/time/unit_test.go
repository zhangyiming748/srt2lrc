package time

import "testing"

func TestSrt2Lrc(t *testing.T) {
	ti := "00:00:56,120 --> 00:00:56,840"
	ret := Srt2Lrc(ti)
	t.Logf("final %s\n", ret)
}

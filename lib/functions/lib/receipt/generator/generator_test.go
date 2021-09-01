package generator

import "testing"

func TestSerialToPng(t *testing.T) {
	if toSerialString(1) != "000001" {
		t.Error("got", toSerialString(1))
	}
}

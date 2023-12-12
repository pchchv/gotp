package otp_test

import (
	"testing"

	"github.com/pchchv/otp"
)

func TestITob(t *testing.T) {
	var i int64 = 1524486261
	expect := []byte{0, 0, 0, 0, 90, 221, 208, 117}

	if string(expect) != string(otp.Itob(i)) {
		t.Error("ITob error")
	}
}

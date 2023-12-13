package otp_test

import (
	"testing"
	"time"

	"github.com/pchchv/otp"
)

var totp = otp.NewDefaultTOTP("4S62BZNFXXSZLCRO")

func TestTOTP_AtTime(t *testing.T) {
	if totp.Now() != totp.AtTime(time.Now()) {
		t.Error("TOTP at time generate otp error!")
	}
}

func TestTOTP_Verify(t *testing.T) {
	if !totp.Verify("179394", 1524485781) {
		t.Error("verify faild")
	}
}

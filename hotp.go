package otp

// HMAC-based OTP counters.
type HOTP struct {
	OTP
}

func NewHOTP(secret string, digits int, hasher *Hasher) *HOTP {
	otp := NewOTP(secret, digits, hasher)
	return &HOTP{OTP: otp}

}

func NewDefaultHOTP(secret string) *HOTP {
	return NewHOTP(secret, 6, nil)
}

// Generates the OTP for the given count.
func (h *HOTP) At(count int) string {
	return h.generateOTP(int64(count))
}

/*
Verify OTP.

params:

	otp:   the OTP to check against
	count: the OTP HMAC counter
*/
func (h *HOTP) Verify(otp string, count int) bool {
	return otp == h.At(count)
}

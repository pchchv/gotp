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

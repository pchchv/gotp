package otp

// time-based OTP counters.
type TOTP struct {
	OTP
	interval int
}

package otp

import "time"

// time-based OTP counters.
type TOTP struct {
	OTP
	interval int
}

func NewTOTP(secret string, digits, interval int, hasher *Hasher) *TOTP {
	otp := NewOTP(secret, digits, hasher)
	return &TOTP{OTP: otp, interval: interval}
}

func NewDefaultTOTP(secret string) *TOTP {
	return NewTOTP(secret, 6, 30, nil)
}

func (t *TOTP) timecode(timestamp int64) int64 {
	return timestamp / int64(t.interval)
}

// Generate time OTP of given timestamp
func (t *TOTP) At(timestamp int64) string {
	return t.generateOTP(t.timecode(timestamp))
}

func (t *TOTP) AtTime(timestamp time.Time) string {
	return t.At(timestamp.Unix())
}

// Generate the current time OTP
func (t *TOTP) Now() string {
	return t.At(currentTimestamp())
}

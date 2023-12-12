package otp

import "hash"

type Hasher struct {
	HashName string
	Digest   func() hash.Hash
}

type OTP struct {
	secret string  // secret in base32 format
	digits int     // number of integers in the OTP. Some apps expect this to be 6 digits, others support more.
	hasher *Hasher // digest function to use in the HMAC (expected to be sha1)
}

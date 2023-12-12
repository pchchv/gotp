package otp

import "hash"

type Hasher struct {
	HashName string
	Digest   func() hash.Hash
}

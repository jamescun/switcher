package main

import (
	"bytes"
)

type SSH string

// address to proxy to
func (s SSH) Address() string {
	return string(s)
}

// identify header as one of SSH
func (s SSH) Identify(header []byte) bool {
	// first 3 bytes of 1.0/2.0 is literal `SSH`
	if bytes.Compare(header, []byte("SSH")) == 0 {
		return true
	}

	return false
}

package main

import (
	"io"
)

// proxy between two sockets
func Shovel(local, remote io.ReadWriteCloser) error {
	errch := make(chan error, 1)

	go chanCopy(errch, local, remote)
	go chanCopy(errch, remote, local)

	for i := 0; i < 2; i++ {
		if err := <-errch; err != nil {
			// If this returns early the second func will push into the
			// buffer, and the GC will clean up
			return err
		}
	}
	return nil
}

// copy between pipes, sending errors to channel
func chanCopy(e chan error, dst, src io.ReadWriter) {
	_, err := io.Copy(dst, src)
	e <- err
}

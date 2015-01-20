package main

import (
	"io"
)

// proxy between two sockets
func Shovel(local, remote io.ReadWriteCloser) error {
	err := make(chan error)

	go chanCopy(err, local, remote)
	go chanCopy(err, remote, local)

	return <-err
}

// copy between pipes, sending errors to channel
func chanCopy(e chan error, dst, src io.ReadWriter) {
	_, err := io.Copy(dst, src)
	if err != nil {
		e <- err
	}
}

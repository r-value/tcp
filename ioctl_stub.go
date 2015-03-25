// Copyright 2014 Mikio Hara. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !darwin,!dragonfly,!freebsd,!linux,!netbsd,!openbsd

package tcp

func (c *Conn) readBufferLen() int {
	return -1
}

func (c *Conn) writeBufferSpace() int {
	return -1
}
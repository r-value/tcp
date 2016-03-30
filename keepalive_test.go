// Copyright 2014 Mikio Hara. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tcp_test

import (
	"net"
	"runtime"
	"testing"
	"time"

	"github.com/mikioh/tcp"
)

func TestKeepAliveOptions(t *testing.T) {
	switch runtime.GOOS {
	case "darwin", "dragonfly", "freebsd", "linux", "netbsd", "solaris", "windows":
	default:
		t.Skipf("%s/%s", runtime.GOOS, runtime.GOARCH)
	}

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer ln.Close()

	go func() {
		for {
			c, err := ln.Accept()
			if err != nil {
				break
			}
			defer c.Close()
		}
	}()

	c, err := net.Dial(ln.Addr().Network(), ln.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()
	tc, err := tcp.NewConn(c)
	if err != nil {
		t.Fatal(err)
	}
	defer tc.Close()

	opt := tcp.KeepAliveOptions{
		IdleInterval:  10 * time.Second, // solaris requires 10 seconds as the lowest value
		ProbeInterval: time.Second,
		ProbeCount:    1,
	}
	if err := tc.Conn.(*net.TCPConn).SetKeepAlive(true); err != nil {
		t.Error(err)
	}
	if err := tc.SetKeepAliveOptions(&opt); err != nil {
		t.Error(err)
	}
}

// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !amd64

package runtime

// This version is used on Linux systems on which we don't use cgo to
// call the C version of sigaction.

//go:nosplit
//go:nowritebarrierrec
func sigaction(sig uint32, new, old *sigactiont) {
	sysSigaction(sig, new, old)
}

// Copyright (c) 2025 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package v2transport

import (
	"github.com/btcsuite/btclog"
)

const Subsystem = "V2TR"

// log is a logger that is initialized with no output filters.  This
// means the package will not perform any logging by default until the caller
// requests it.
var log btclog.Logger

// The default amount of logging is none.
func init() {
	DisableLog()
}

// DisableLog disables all library log output.  Logging output is disabled
// by default until UseLogger is called.
func DisableLog() {
	log = btclog.Disabled
}

// UseLogger uses a specified Logger to output package logging info.
func UseLogger(logger btclog.Logger) {
	log = logger
}

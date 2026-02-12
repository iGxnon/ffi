//go:build !ffi_no_embed

package ffi

import _ "embed"

const libname = "libffi.so.8"

//go:embed assets/libffi/linux_arm64/libffi.so.8
var embeddedLib []byte

//go:build !ffi_no_embed

package ffi

import _ "embed"

const libname = "libffi.so.8"

//go:embed assets/libffi/linux_amd64/libffi.so.8
var embeddedLib []byte

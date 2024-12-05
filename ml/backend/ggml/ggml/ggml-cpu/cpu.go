package cpu

// #cgo CFLAGS: -std=c17
// #cgo CXXFLAGS: -std=c++17
// #cgo CPPFLAGS: -I${SRCDIR}/.. -I${SRCDIR}/../include -I${SRCDIR}/amx
// #cgo amd64,avx CPPFLAGS: -mavx
// #cgo amd64,avx2 CPPFLAGS: -mavx2 -mfma
// #cgo amd64,f16c CPPFLAGS: -mf16c
// #include "ggml-cpu.h"
// #include "ggml-backend.h"
import "C"

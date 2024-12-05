package ggml

// #cgo avx2 CPPFLAGS: -DGGML_USE_ACCELERATE -DACCELERATE_USE_LAPACK -DACCELERATE_LAPACK_ILP64
// #cgo avx2 LDFLAGS: -framework Accelerate
import "C"

func newBackend() *C.struct_ggml_backend {
	return newCPUBackend()
}

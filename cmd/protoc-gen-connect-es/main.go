package main

import (
	"os"

	"github.com/wasilibs/go-protoc-gen-connect-es/internal/runner"
	"github.com/wasilibs/go-protoc-gen-connect-es/internal/wasm"
)

func main() {
	os.Exit(runner.Run("protoc-gen-connect-es", os.Args[1:], wasm.ProtocGenConnectES, os.Stdin, os.Stdout, os.Stderr, "."))
}

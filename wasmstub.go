// Our empty version of the httpServer for usage with the wasm target
// this way we will not include any of the related code
//go:build wasm

package main

func AppServer() {
}

package main

import "github.com/opa-oz/emoji-lsp/pkg/server"

const lsName = "Emoji LSP"
const version string = "0.0.4"

func main() {
	server.SpinUp(lsName, version, 0, false)
}

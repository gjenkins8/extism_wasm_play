package main

import (
	"context"
	_ "embed"
	"fmt"
	"os"

	extism "github.com/extism/go-sdk"
)

//go:embed count_vowels_go.wasm
var countVowelsWasm []byte

func main() {
	fmt.Printf("len: %d\n", len(countVowelsWasm))
	manifest := extism.Manifest{
		Wasm: []extism.Wasm{
			extism.WasmData{
				Data: countVowelsWasm,
			},
		},
	}

	ctx := context.Background()
	config := extism.PluginConfig{
		EnableWasi: true,
	}
	plugin, err := extism.NewPlugin(ctx, manifest, config, []extism.HostFunction{})

	if err != nil {
		fmt.Printf("Failed to initialize plugin: %v\n", err)
		os.Exit(1)
	}

	data := []byte("Hello, World!")
	exit, out, err := plugin.Call("count_vowels", data)
	if err != nil {
		fmt.Println(err)
		os.Exit(int(exit))
	}

	response := string(out)
	fmt.Println(response)
}

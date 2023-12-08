package main

import (
	"os"

	"github.com/mrbns/assetLoader/helper"
	"github.com/mrbns/assetLoader/internal/generator"
)

func init() {

	args := os.Args

	for _, arg := range args {
		helper.ProcessArg(arg)
	}
}

func main() {

	generator.GenerateAsset("")

}

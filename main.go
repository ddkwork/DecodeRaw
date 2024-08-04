package main

import (
	"DecodeRaw"

	"github.com/ddkwork/unison"
	"github.com/ddkwork/unison/app"
)

func main() {
	app.Run("DecodeRaw", func(w *unison.Window) {
		w.Content().AddChild(decodeRaw.New().Layout())
	})
}

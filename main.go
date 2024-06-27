package main

import (
	"DecodeRaw"

	"github.com/ddkwork/app"
	"github.com/richardwilkes/unison"
)

func main() {
	app.Run("DecodeRaw", func(w *unison.Window) {
		w.Content().AddChild(decodeRaw.New().Layout())
	})
}

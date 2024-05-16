package main

import (
	"DecodeRaw"

	"github.com/ddkwork/app"
	"github.com/richardwilkes/unison"
)

func main() {
	app.Run("DecodeRaw", func(w *unison.Window) {
		DecodeRaw.New().Layout(w.Content())
	})
}

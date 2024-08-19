package dongming

import (
	"github.com/gookit/color"
)

func Run() {
	err := auth()
	if err != nil {
		color.Redln(err)
	}
}

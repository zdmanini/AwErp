package routers

import (
	"Awesome/core"
	"Awesome/generator/routers/gen"
)

var InitRouters = []*core.GroupBase{
	// gen
	gen.GenGroup,
}

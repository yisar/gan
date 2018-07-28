package gan

import (
	"github.com/julienschmidt/httprouter"
)

type Gay struct {
	Router     *httprouter.Router
}

func New() *Gay {
	r := httprouter.New()
	return &Gay{
		Router: r,
	}
}

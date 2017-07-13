package handlers

import (
	macaron "gopkg.in/macaron.v1"

	"github.com/mbier/aceleratomo/acelerato"
	"github.com/mbier/aceleratomo/projeto"
)

func Quadro(ctx *macaron.Context) {

	projeto, err := projeto.GetProjeto(ctx.Params(":nome"))

	if err != nil {
		panic(err)
	}

	ctx.Resp.Header().Set("Content-Type", "text/html")
	ctx.Resp.Write([]byte(acelerato.GerarQuadro(projeto)))
}

func QuadroGeral(ctx *macaron.Context) {
	ctx.Resp.Header().Set("Content-Type", "text/html")
	ctx.Resp.Write([]byte(acelerato.GerarQuadroGeral()))
}

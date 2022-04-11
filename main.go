package main

import (
	_ "go-crm/internal/packed"
	"github.com/gogf/gf/v2/os/gctx"
	"go-crm/internal/cmd"

)

func main() {
	cmd.Main.Run(gctx.New())
}

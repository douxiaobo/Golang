package main

import (
	"fmt"

	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	var (
		ctx    = gctx.New()
		name   = "sum"
		amount = 1
	)
	i18n := gi18n.New()

	i18n.SetLanguage("zh-CN")
	fmt.Println(i18n.Tf(ctx, `{#paySuccess}`, name, amount))
}

// sum用户支付成功1元

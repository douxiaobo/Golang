package main

import (
	"fmt"

	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	var (
		ctx  = gctx.New()
		i18n = gi18n.New()
	)
	i18n.SetLanguage("en")
	fmt.Println(i18n.Translate(ctx, `hello`))
	fmt.Println(i18n.Translate(ctx, `GF says: {#hello}{#world}!`))
	i18n.SetLanguage("ja")
	fmt.Println(i18n.Translate(ctx, `hello`))
	fmt.Println(i18n.Translate(ctx, `GF says: {#hello}{#world}!`))
	i18n.SetLanguage("ru")
	fmt.Println(i18n.Translate(ctx, `hello`))
	fmt.Println(i18n.Translate(ctx, `GF says: {#hello}{#world}!`))
	ctx = gi18n.WithLanguage(ctx, "zh-CN")
	fmt.Println(i18n.Translate(ctx, `hello`))
	fmt.Println(i18n.Translate(ctx, `GF says: {#hello}{#world}!`))
}

// douxiaobo@192 i18n % go mod init i18n
// go: creating new go.mod: module i18n
// go: to add module requirements and sums:
// 	go mod tidy
// douxiaobo@192 i18n % go get github.com/gogf/gf/v2/
// go: downloading github.com/gogf/gf v1.16.9
// go: downloading github.com/gogf/gf/v2 v2.7.2
// go: added github.com/gogf/gf/v2 v2.7.2
// douxiaobo@192 i18n % go run
// go: no go files listed
// douxiaobo@192 i18n % go run main.go
// ../../../../go/pkg/mod/github.com/gogf/gf/v2@v2.7.2/internal/intlog/intlog.go:17:2: missing go.sum entry for module providing package go.opentelemetry.io/otel/trace (imported by github.com/gogf/gf/v2/internal/intlog); to add:
// 	go get github.com/gogf/gf/v2/internal/intlog@v2.7.2
// ../../../../go/pkg/mod/github.com/gogf/gf/v2@v2.7.2/container/gtree/gtree_btree.go:12:2: missing go.sum entry for module providing package github.com/emirpasic/gods/trees/btree (imported by github.com/gogf/gf/v2/container/gtree); to add:
// 	go get github.com/gogf/gf/v2/container/gtree@v2.7.2
// ../../../../go/pkg/mod/github.com/gogf/gf/v2@v2.7.2/encoding/gproperties/gproperties.go:15:2: missing go.sum entry for module providing package github.com/magiconair/properties (imported by github.com/gogf/gf/v2/encoding/gproperties); to add:
// 	go get github.com/gogf/gf/v2/encoding/gproperties@v2.7.2
// ../../../../go/pkg/mod/github.com/gogf/gf/v2@v2.7.2/encoding/gtoml/gtoml.go:13:2: missing go.sum entry for module providing package github.com/BurntSushi/toml (imported by github.com/gogf/gf/v2/encoding/gtoml); to add:
// 	go get github.com/gogf/gf/v2/encoding/gtoml@v2.7.2
// ../../../../go/pkg/mod/github.com/gogf/gf/v2@v2.7.2/encoding/gxml/gxml.go:13:2: missing go.sum entry for module providing package github.com/clbanning/mxj/v2 (imported by github.com/gogf/gf/v2/encoding/gxml); to add:
// 	go get github.com/gogf/gf/v2/encoding/gxml@v2.7.2
// ../../../../go/pkg/mod/github.com/gogf/gf/v2@v2.7.2/encoding/gcharset/gcharset.go:27:2: missing go.sum entry for module providing package golang.org/x/text/encoding (imported by github.com/gogf/gf/v2/encoding/gcharset); to add:
// 	go get github.com/gogf/gf/v2/encoding/gcharset@v2.7.2
// ../../../../go/pkg/mod/github.com/gogf/gf/v2@v2.7.2/encoding/gcharset/gcharset.go:28:2: missing go.sum entry for module providing package golang.org/x/text/encoding/ianaindex (imported by github.com/gogf/gf/v2/encoding/gcharset); to add:
// 	go get github.com/gogf/gf/v2/encoding/gcharset@v2.7.2
// ../../../../go/pkg/mod/github.com/gogf/gf/v2@v2.7.2/encoding/gcharset/gcharset.go:29:2: missing go.sum entry for module providing package golang.org/x/text/transform (imported by github.com/gogf/gf/v2/encoding/gcharset); to add:
// 	go get github.com/gogf/gf/v2/encoding/gcharset@v2.7.2
// ../../../../go/pkg/mod/github.com/gogf/gf/v2@v2.7.2/encoding/gyaml/gyaml.go:14:2: missing go.sum entry for module providing package gopkg.in/yaml.v3 (imported by github.com/gogf/gf/v2/encoding/gyaml); to add:
// 	go get github.com/gogf/gf/v2/encoding/gyaml@v2.7.2
// ../../../../go/pkg/mod/github.com/gogf/gf/v2@v2.7.2/os/gfsnotify/gfsnotify.go:15:2: missing go.sum entry for module providing package github.com/fsnotify/fsnotify (imported by github.com/gogf/gf/v2/os/gfsnotify); to add:
// 	go get github.com/gogf/gf/v2/os/gfsnotify@v2.7.2
// ../../../../go/pkg/mod/github.com/gogf/gf/v2@v2.7.2/net/gtrace/internal/provider/provider.go:10:2: missing go.sum entry for module providing package go.opentelemetry.io/otel/sdk/trace (imported by github.com/gogf/gf/v2/net/gtrace/internal/provider); to add:
// 	go get github.com/gogf/gf/v2/net/gtrace/internal/provider@v2.7.2
// ../../../../go/pkg/mod/github.com/gogf/gf/v2@v2.7.2/net/gtrace/gtrace.go:15:2: missing go.sum entry for module providing package go.opentelemetry.io/otel (imported by github.com/gogf/gf/v2/os/gctx); to add:
// 	go get github.com/gogf/gf/v2/os/gctx@v2.7.2
// ../../../../go/pkg/mod/github.com/gogf/gf/v2@v2.7.2/net/gtrace/gtrace.go:16:2: missing go.sum entry for module providing package go.opentelemetry.io/otel/attribute (imported by github.com/gogf/gf/v2/net/gtrace); to add:
// 	go get github.com/gogf/gf/v2/net/gtrace@v2.7.2
// ../../../../go/pkg/mod/github.com/gogf/gf/v2@v2.7.2/net/gtrace/gtrace_baggage.go:12:2: missing go.sum entry for module providing package go.opentelemetry.io/otel/baggage (imported by github.com/gogf/gf/v2/net/gtrace); to add:
// 	go get github.com/gogf/gf/v2/net/gtrace@v2.7.2
// ../../../../go/pkg/mod/github.com/gogf/gf/v2@v2.7.2/net/gtrace/gtrace.go:17:2: missing go.sum entry for module providing package go.opentelemetry.io/otel/propagation (imported by github.com/gogf/gf/v2/os/gctx); to add:
// 	go get github.com/gogf/gf/v2/os/gctx@v2.7.2
// ../../../../go/pkg/mod/github.com/gogf/gf/v2@v2.7.2/net/gtrace/gtrace.go:18:2: missing go.sum entry for module providing package go.opentelemetry.io/otel/semconv/v1.18.0 (imported by github.com/gogf/gf/v2/net/gtrace); to add:
// 	go get github.com/gogf/gf/v2/net/gtrace@v2.7.2
// douxiaobo@192 i18n % go get gibhub.com/gogf/gf/v2/i18n/gi18n
// go: unrecognized import path "gibhub.com/gogf/gf/v2/i18n/gi18n": https fetch: Get "https://gibhub.com/gogf/gf/v2/i18n/gi18n?go-get=1": tls: failed to verify certificate: x509: certificate is valid for expiereddnsmanager.com, www.expiereddnsmanager.com, not gibhub.com
// douxiaobo@192 i18n % go get github.com/gogf/gf/v2/i18n/gi18n
// go: downloading github.com/emirpasic/gods v1.18.1
// go: downloading go.opentelemetry.io/otel v1.14.0
// go: downloading github.com/fsnotify/fsnotify v1.7.0
// go: downloading go.opentelemetry.io/otel/trace v1.14.0
// go: downloading go.opentelemetry.io/otel/sdk v1.14.0
// go: downloading github.com/magiconair/properties v1.8.7
// go: downloading github.com/BurntSushi/toml v1.3.2
// go: downloading github.com/clbanning/mxj/v2 v2.7.0
// go: downloading golang.org/x/sys v0.19.0
// go: downloading github.com/go-logr/logr v1.2.3
// go: downloading github.com/go-logr/stdr v1.2.2
// douxiaobo@192 i18n % go run main.go
// hello
// GF says: {#hello}{#world}!
// hello
// GF says: {#hello}{#world}!
// hello
// GF says: {#hello}{#world}!
// hello
// GF says: {#hello}{#world}!
// douxiaobo@192 i18n % go run main.go
// hello
// GF says: {#hello}{#world}!
// hello
// GF says: {#hello}{#world}!
// hello
// GF says: {#hello}{#world}!
// hello
// GF says: {#hello}{#world}!
// douxiaobo@192 i18n % go run main.go
// hello
// GF says: helloworld!
// hello
// GF says: {#hello}{#world}!
// hello
// GF says: {#hello}{#world}!
// hello
// GF says: {#hello}{#world}!
// douxiaobo@192 i18n % go run main.go
// hello
// GF says: helloworld!
// hello
// GF says: {#hello}{#world}!
// hello
// GF says: {#hello}{#world}!
// 你好
// GF says: 你好世界!
// douxiaobo@192 i18n %

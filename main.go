package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/xenoPot/json-api-sandbox/handler"
)

func main() {
	// Echo使ってみる
	// インスタンス生成
	e := echo.New()

	// ミドルウェア設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルーティング
	e.GET("/test", handler.Test())

	// サーバ起動
	e.Start(":8080")
}

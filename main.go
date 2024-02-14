package main

import (
	"log"

	"github.com/7oh2020/simple-bluesky-webhook/config"
	"github.com/7oh2020/simple-bluesky-webhook/webhook"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	// 環境変数から設定値を取得する
	cfg, err := config.GetConfig()
	if err != nil {
		return err
	}

	// Webhookサーバーをセットアップする
	ws, err := webhook.NewWebhookServer(cfg.Handle, cfg.Pass, cfg.Port)
	if err != nil {
		return err
	}

	// サーバーを起動する
	return ws.ListenAndServe()
}

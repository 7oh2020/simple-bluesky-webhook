package config

import (
	"github.com/caarlos0/env/v10"
)

// サーバー設定
type Config struct {
	// サーバーのポート番号
	Port string `env:"PORT" envDefault:"8080"`
	// アカウントのハンドル(DNS名)
	Handle string `env:"HANDLE,required"`
	// アカウントのパスワード
	Pass string `env:"PASS,required"`
}

// 環境変数をバインドして設定値を取得します。
func GetConfig() (*Config, error) {
	var value Config
	if err := env.Parse(&value); err != nil {
		return nil, err
	}
	return &value, nil
}

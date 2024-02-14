# Simple Bluesky Webhook

これは Bluesky へテキストを投稿するためのシンプルな Webhook です。
セッションはサーバーを終了するまでメモリ上で保持されます。

## 使い方

以下のコマンドを実行するとサーバーが起動します。
環境変数による接続先の指定が必要です。

- HANDLE は`example.bsky.social`のようなあなたの Bluesky でのドメイン名です。
- PASS は Bluesky でのあなたのパスワードです。

```bash
HANDLE="your handle" PASS="your password" go run ./main.go
```

さらに任意のポート番号も指定できます。(デフォルトは 8080 です。)

```bash
PORT="8000" HANDLE="your handle" PASS="your password" go run ./main.go
```

サーバーの起動時に WebhookURL が表示されます。
WebhookURL はサーバーが起動するたびに変わります。

```txt
server is running on port 8080
webhook url is accessible at: POST http://localhost:8080/webhook/...
```

WebhookURL に POST アクセスするとテキストが Bluesky へ投稿されます。

```bash
curl -d "text=Hello World" http://localhost:8080/webhook/...
```

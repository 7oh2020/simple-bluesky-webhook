package webhook

import (
	"fmt"
	"net/http"

	"github.com/7oh2020/simple-bluesky-webhook/api"
	"github.com/7oh2020/simple-bluesky-webhook/validation"
)

// URLに使用するランダム文字列の長さ
const TokenSize = 16

// Webhookサーバー
type WebhookServer struct {
	// サーバーのポート番号
	port string
	// WebhookURLに使用されるURLセーフなランダムな文字列
	token string
	// BlueskyAPIのラッパー
	blueskyAPI *api.BlueskyAPI
}

func NewWebhookServer(handle, pass, port string) (*WebhookServer, error) {
	ba := api.NewBlueskyAPI()

	// セッションを作成する
	if err := ba.CreateSession(handle, pass); err != nil {
		return nil, err
	}

	// ランダム文字列を生成する
	token, err := GenerateToken(TokenSize)
	if err != nil {
		return nil, err
	}

	return &WebhookServer{
		port:       port,
		token:      token,
		blueskyAPI: ba,
	}, nil
}

// ハンドラを定義してサーバーを起動します。
func (ws *WebhookServer) ListenAndServe() error {
	// 各種ハンドラを定義する
	mux := http.NewServeMux()
	mux.HandleFunc("/", ws.handleIndex)
	mux.HandleFunc("/webhook/"+ws.token, ws.handleWebhook)

	fmt.Printf("server is running on port %s\n", ws.port)
	fmt.Printf("webhook url is accessible at: POST http://localhost:%s/webhook/%s\n", ws.port, ws.token)

	// サーバーを起動する
	return http.ListenAndServe(fmt.Sprintf(":%s", ws.port), mux)
}

// トップページ(GET /)のハンドラ
func (ws *WebhookServer) handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Simple Bluesky Webhook!")
}

// WebhookURL(POST /webhook/ + token)のハンドラ
func (ws *WebhookServer) handleWebhook(w http.ResponseWriter, r *http.Request) {
	// POSTメソッドのみを許可する
	if r.Method != http.MethodPost {
		http.Error(w, "invalid http method", http.StatusMethodNotAllowed)
		return
	}

	// POSTパラメータを取得する
	if err := r.ParseForm(); err != nil {
		http.Error(w, "failed to parse form", http.StatusBadRequest)
		return
	}
	text := r.PostFormValue("text")

	// パラメータをバリデーションする
	if err := validation.ValidateText(text); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// セッションをリフレッシュする
	if err := ws.blueskyAPI.RefreshSession(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// テキストを投稿する
	if err := ws.blueskyAPI.CreatePost(text); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "post created successfully!")

}

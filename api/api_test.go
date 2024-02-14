package api

import (
	"testing"
	"time"

	"github.com/7oh2020/simple-bluesky-webhook/config"

	"github.com/stretchr/testify/require"
)

func TestApi(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	cfg, err := config.GetConfig()
	require.NoError(t, err, "エラーが発生しないこと")

	duration := time.Second
	ba := NewBlueskyAPI()

	// セッションを作成する
	err1 := ba.CreateSession(cfg.Handle, cfg.Pass)
	require.NoError(t, err1, "エラーが発生しないこと")

	time.Sleep(duration)

	// セッション作成後のトークンを使用する
	p1, err2 := ba.GetProfile(ba.client.Auth.Handle)
	require.NoError(t, err2, "エラーが発生しないこと")
	require.Equal(t, ba.client.Auth.Handle, p1.Handle, "プロフィールが取得できること")

	AccessJwt := ba.client.Auth.AccessJwt
	RefreshJwt := ba.client.Auth.RefreshJwt

	time.Sleep(duration)

	// セッションをリフレッシュする
	err3 := ba.RefreshSession()
	require.NoError(t, err3, "エラーが発生しないこと")
	require.NotEqual(t, AccessJwt, ba.client.Auth.AccessJwt, "アクセストークンが更新されていること")
	require.NotEqual(t, RefreshJwt, ba.client.Auth.RefreshJwt, "リフレッシュトークンが更新されていること")

	time.Sleep(duration)

	// セッションリフレッシュ後のトークンを使用する
	p2, err4 := ba.GetProfile(ba.client.Auth.Handle)
	require.NoError(t, err4, "エラーが発生しないこと")
	require.Equal(t, ba.client.Auth.Handle, p2.Handle, "プロフィールが取得できること")

	time.Sleep(duration)

	// セッションを削除する
	err5 := ba.DeleteSession()
	require.NoError(t, err5, "エラーが発生しないこと")

	time.Sleep(duration)

	// セッション削除後のトークンを使用する
	_, err6 := ba.GetProfile(ba.client.Auth.Handle)
	require.EqualError(t, err6, "XRPC ERROR 400: InvalidToken: Bad token scope", "認証エラーになること")
}

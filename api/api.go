package api

import (
	"context"
	"fmt"
	"time"

	"github.com/bluesky-social/indigo/api/atproto"
	"github.com/bluesky-social/indigo/api/bsky"
	lexutil "github.com/bluesky-social/indigo/lex/util"
	"github.com/bluesky-social/indigo/util"
	"github.com/bluesky-social/indigo/util/cliutil"
	"github.com/bluesky-social/indigo/xrpc"
)

// 接続先ホスト
const DefaultHost = "https://bsky.social"

// BlueskyAPIのラッパー
type BlueskyAPI struct {
	// XRPCクライアント
	client *xrpc.Client
}

func NewBlueskyAPI() *BlueskyAPI {
	return &BlueskyAPI{}
}

// セッションを作成し、クライアントを返します。
func (ba *BlueskyAPI) CreateSession(handle, pass string) error {
	client := &xrpc.Client{
		Client: cliutil.NewHttpClient(),
		Host:   DefaultHost,
	}

	// 入力パラメータを作成する
	params := &atproto.ServerCreateSession_Input{
		Identifier: handle,
		Password:   pass,
	}

	// セッションを作成する
	ses, err := atproto.ServerCreateSession(context.TODO(), client, params)
	if err != nil {
		return err
	}

	// 取得したセッション情報をクライアントにセットする
	client.Auth = &xrpc.AuthInfo{
		AccessJwt:  ses.AccessJwt,
		RefreshJwt: ses.RefreshJwt,
		Handle:     ses.Handle,
		Did:        ses.Did,
	}
	ba.client = client

	return nil
}

// 現在のセッションをリフレッシュします。
func (ba *BlueskyAPI) RefreshSession() error {
	if ba.client == nil {
		return fmt.Errorf("client not created")
	}

	// リフレッシュトークンをセットする
	ba.client.Auth.AccessJwt = ba.client.Auth.RefreshJwt

	// 新しいセッションを取得する
	ses, err := atproto.ServerRefreshSession(context.TODO(), ba.client)
	if err != nil {
		return err
	}

	// 取得したセッション情報をクライアントにセットする
	ba.client.Auth = &xrpc.AuthInfo{
		AccessJwt:  ses.AccessJwt,
		RefreshJwt: ses.RefreshJwt,
		Handle:     ses.Handle,
		Did:        ses.Did,
	}

	return nil
}

// 現在のセッションを削除します。
func (ba *BlueskyAPI) DeleteSession() error {
	if ba.client == nil {
		return fmt.Errorf("client not created")
	}

	// リフレッシュトークンをセットする
	ba.client.Auth.AccessJwt = ba.client.Auth.RefreshJwt

	// セッションを削除する
	if err := atproto.ServerDeleteSession(context.TODO(), ba.client); err != nil {
		return err
	}
	return nil
}

// 指定アカウントのプロフィールを取得します。
func (ba *BlueskyAPI) GetProfile(handle string) (*bsky.ActorDefs_ProfileViewDetailed, error) {
	if ba.client == nil {
		return nil, fmt.Errorf("client not created")
	}

	// プロフィールを取得する
	profile, err := bsky.ActorGetProfile(context.TODO(), ba.client, handle)
	if err != nil {
		return nil, err
	}
	return profile, nil
}

// 現在のセッションでtextを投稿します。
func (ba *BlueskyAPI) CreatePost(text string) error {
	if ba.client == nil {
		return fmt.Errorf("client not created")
	}

	// 入力パラメータを作成する
	params := &atproto.RepoCreateRecord_Input{
		// コレクションのNSID(名前空間識別子)
		Collection: "app.bsky.feed.post",
		// リポジトリのDIDもしくはハンドル
		Repo: ba.client.Auth.Did,
		Record: &lexutil.LexiconTypeDecoder{
			Val: &bsky.FeedPost{
				// 本文テキスト
				Text: text,
				// 作成日時
				CreatedAt: time.Now().Format(util.ISO8601),
				// 言語(複数指定可能)
				Langs: []string{"ja"},
			},
		},
	}

	// テキストを投稿する
	_, err := atproto.RepoCreateRecord(context.TODO(), ba.client, params)
	if err != nil {
		return err
	}
	return nil
}

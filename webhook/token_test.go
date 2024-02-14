package webhook

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWebhook_GenerateToken(t *testing.T) {
	n := 16

	ret1, err1 := GenerateToken(n)
	require.NoError(t, err1, "エラーが発生しないこと")

	ret2, err2 := GenerateToken(n)
	require.NoError(t, err2, "エラーが発生しないこと")

	require.NotEqual(t, ret1, ret2, "生成される文字列が毎回異なること")
}

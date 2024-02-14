package validation

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidation_ValidateText(tt *testing.T) {
	testcases := []struct {
		title string
		arg   string
		err   error
	}{
		{"正常な入力の場合", strings.Repeat("*", 300), nil},
		{"空文字の場合", "", errors.New("can not be blank")},
		{"半角で最大値を超える場合", strings.Repeat("*", 301), errors.New("length must be less than or equal to 300")},
		{"全角で最大値を超える場合", strings.Repeat("あ", 301), errors.New("length must be less than or equal to 300")},
	}

	for _, v := range testcases {
		tt.Run(v.title, func(t *testing.T) {
			err := ValidateText(v.arg)
			if err == nil {
				require.NoError(t, err, "エラーが発生しないこと")
			} else {
				require.EqualError(t, err, v.err.Error(), "エラーが一致すること")
			}

		})
	}
}

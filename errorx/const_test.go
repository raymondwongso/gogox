package errorx_test

import (
	"testing"

	"github.com/raymondwongso/gogox/errorx"
	"github.com/stretchr/testify/assert"
)

func Test_ErrInternal(t *testing.T) {
	err := errorx.ErrInternal("some error")
	assert.Equal(t, "some error", err.Error())
	assert.Equal(t, errorx.CodeInternal, err.Code)
}

func Test_ErrNotFound(t *testing.T) {
	err := errorx.ErrNotFound("some error")
	assert.Equal(t, "some error", err.Error())
	assert.Equal(t, errorx.CodeNotFound, err.Code)
}

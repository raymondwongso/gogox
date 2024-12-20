package errorx_test

import (
	"fmt"
	"testing"

	"github.com/raymondwongso/gogox/errorx"
	"github.com/stretchr/testify/assert"
)

func Test_ErrInternal(t *testing.T) {
	err := errorx.ErrInternal("some error")
	assert.Equal(t, "some error", err.Error())
	assert.Equal(t, errorx.CodeInternal, err.Code)
	assert.Equal(t, fmt.Sprintf("[%s] %s", err.Code, err.Error()), err.LogError())
}

func Test_ErrNotFound(t *testing.T) {
	err := errorx.ErrNotFound("some error")
	assert.Equal(t, "some error", err.Error())
	assert.Equal(t, errorx.CodeNotFound, err.Code)
	assert.Equal(t, fmt.Sprintf("[%s] %s", err.Code, err.Error()), err.LogError())
}

func Test_ErrUnauthorized(t *testing.T) {
	err := errorx.ErrUnauthorized("some error")
	assert.Equal(t, "some error", err.Error())
	assert.Equal(t, errorx.CodeUnauthorized, err.Code)
	assert.Equal(t, fmt.Sprintf("[%s] %s", err.Code, err.Error()), err.LogError())
}

func Test_ErrInvalidParameter(t *testing.T) {
	err := errorx.ErrInvalidParameter("some error")
	assert.Equal(t, "some error", err.Error())
	assert.Equal(t, errorx.CodeInvalidParameter, err.Code)
	assert.Equal(t, fmt.Sprintf("[%s] %s", err.Code, err.Error()), err.LogError())
}

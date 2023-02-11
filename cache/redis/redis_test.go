package redis_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/go-redis/redismock/v8"
	"github.com/raymondwongso/gogox/cache/redis"
	"github.com/raymondwongso/gogox/errorx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	goredis "github.com/go-redis/redis/v8"
)

type RedisTestSuite struct {
	suite.Suite
	client *redis.Redis
	mock   redismock.ClientMock
}

func TestRedisSuite(t *testing.T) {
	suite.Run(t, new(RedisTestSuite))
}

func (ts *RedisTestSuite) SetupTest() {
	rds, mock := redismock.NewClientMock()

	ts.client = redis.New(rds)
	ts.mock = mock
}

type testStruct struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

const (
	tempKey = "tempkey"
)

func (ts *RedisTestSuite) Test_Get() {
	testRes := testStruct{ID: 1, Name: "John"}

	ts.Run("success", func() {
		dest := testStruct{}

		expStr := ts.mock.ExpectGet(tempKey)
		resBytes, _ := json.Marshal(&testRes)
		expStr.SetVal(string(resBytes))

		err := ts.client.Get(context.Background(), tempKey, &dest)

		assert.Equal(ts.T(), &testRes, &dest)
		assert.NoError(ts.T(), err)
	})

	ts.Run("error - not found", func() {
		dest := testStruct{}

		expStr := ts.mock.ExpectGet(tempKey)
		resBytes, _ := json.Marshal(&testStruct{})
		expStr.SetVal(string(resBytes))
		expStr.SetErr(goredis.Nil)

		err := ts.client.Get(context.Background(), tempKey, &dest)

		assert.Equal(ts.T(), &testStruct{}, &dest)
		gogoxErr, _ := errorx.Parse(err)
		assert.Equal(ts.T(), errorx.CodeNotFound, gogoxErr.Code)
	})

	ts.Run("error - unexpected", func() {
		dest := testStruct{}

		expStr := ts.mock.ExpectGet(tempKey)
		resBytes, _ := json.Marshal(&testStruct{})
		expStr.SetVal(string(resBytes))
		expStr.SetErr(fmt.Errorf("some error"))

		err := ts.client.Get(context.Background(), tempKey, &dest)

		assert.Equal(ts.T(), &testStruct{}, &dest)
		assert.Equal(ts.T(), fmt.Errorf("some error"), err)
	})
}

func (ts *RedisTestSuite) Test_Set() {
	type args struct {
		ctx        context.Context
		value      *testStruct
		expiration time.Duration
	}

	type mock struct {
		err error
	}

	type exp struct {
		err error
	}

	testCases := map[string]struct {
		args args
		mock mock
		exp  exp
	}{
		"success": {
			args: args{
				ctx:        context.Background(),
				value:      &testStruct{ID: 1},
				expiration: 1 * time.Hour,
			},
		},
		"error - unexpected": {
			args: args{
				ctx:   context.Background(),
				value: &testStruct{ID: 1},
			},
			mock: mock{
				err: errors.New("some error"),
			},
			exp: exp{
				err: errors.New("some error"),
			},
		},
	}

	for name, tc := range testCases {
		ts.Run(name, func() {
			setBytes, _ := json.Marshal(tc.args.value)

			expStr := ts.mock.ExpectSet(tempKey, setBytes, tc.args.expiration)
			if tc.mock.err != nil {
				expStr.SetErr(tc.mock.err)
			}

			expStr.SetVal("1")

			err := ts.client.Set(tc.args.ctx, tempKey, tc.args.value, tc.args.expiration)

			assert.Equal(ts.T(), tc.exp.err, err)
		})
	}
}

func (ts *RedisTestSuite) Test_Del() {
	type args struct {
		ctx context.Context
	}

	type mock struct {
		err error
	}

	type exp struct {
		err error
	}

	testCases := map[string]struct {
		args args
		mock mock
		exp  exp
	}{
		"success": {
			args: args{
				ctx: context.Background(),
			},
		},
		"error - unexpected": {
			args: args{
				ctx: context.Background(),
			},
			mock: mock{
				err: errors.New("some error"),
			},
			exp: exp{
				err: errors.New("some error"),
			},
		},
	}

	for name, tc := range testCases {
		ts.Run(name, func() {
			expStr := ts.mock.ExpectDel(tempKey)
			if tc.mock.err != nil {
				expStr.SetErr(tc.mock.err)
			}

			expStr.SetVal(1)

			err := ts.client.Del(tc.args.ctx, tempKey)

			assert.Equal(ts.T(), tc.exp.err, err)
		})
	}
}

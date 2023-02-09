package redis_test

import (
	"context"
	"encoding/json"
	"errors"
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
	type args struct {
		ctx context.Context
	}

	type mock struct {
		res interface{}
		err error
	}

	type exp struct {
		res interface{}
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
			mock: mock{
				res: &testStruct{ID: 1, Name: "John"},
			},
			exp: exp{
				res: &testStruct{ID: 1, Name: "John"},
			},
		},
		"error - key not found": {
			args: args{
				ctx: context.Background(),
			},
			mock: mock{
				err: goredis.Nil,
			},
			exp: exp{
				res: &testStruct{},
				err: errorx.ErrNotFound("redis key is not found"),
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
				res: &testStruct{},
				err: errors.New("some error"),
			},
		},
	}

	for name, tc := range testCases {
		ts.Run(name, func() {
			dest := testStruct{}

			expStr := ts.mock.ExpectGet(tempKey)
			if tc.mock.err != nil {
				expStr.SetErr(tc.mock.err)
			}

			if tc.mock.res != nil {
				resBytes, _ := json.Marshal(tc.mock.res)
				expStr.SetVal(string(resBytes))
			}

			err := ts.client.Get(tc.args.ctx, tempKey, &dest)

			assert.Equal(ts.T(), tc.exp.res, &dest)
			assert.Equal(ts.T(), tc.exp.err, err)
		})
	}
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

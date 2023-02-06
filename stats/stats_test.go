package stats_test

import (
	"testing"

	"github.com/raymondwongso/gogox/stats"
	"github.com/stretchr/testify/assert"
)

func Test_MergeTags(t *testing.T) {
	expected := stats.Tags{
		"a": "1",
		"b": "2",
	}

	assert.Equal(t, expected, stats.MergeTags(nil, stats.Tags{"a": "1", "b": "2"}))
	assert.Equal(t, expected, stats.MergeTags(stats.Tags{"a": "1", "b": "2"}, nil))
	assert.Equal(t, expected, stats.MergeTags(stats.Tags{"b": "2"}, stats.Tags{"a": "1"}))
}

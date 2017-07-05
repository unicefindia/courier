package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapAsJSON(t *testing.T) {
	assert.Equal(t, "{}", string(MapAsJSON(map[string]string{})))
	assert.Equal(t, "{\"foo\":\"bar\"}", string(MapAsJSON(map[string]string{"foo": "bar"})))
}

func TestJoinNonEmpty(t *testing.T) {
	assert.Equal(t, "", JoinNonEmpty(" "))
	assert.Equal(t, "hello world", JoinNonEmpty(" ", "", "hello", "", "world"))
}

func TestStringArrayContains(t *testing.T) {
	assert.False(t, StringArrayContains([]string{}, "x"))
	assert.False(t, StringArrayContains([]string{"a", "b"}, "x"))
	assert.True(t, StringArrayContains([]string{"a", "b", "x", "y"}, "x"))
}
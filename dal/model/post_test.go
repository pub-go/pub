package model_test

import (
	"testing"

	"code.gopub.tech/pub/dal/model"
	"code.gopub.tech/pub/util/assert"
)

func TestAllow(t *testing.T) {
	p := &model.Post{}
	assert.True(t, p.AllowPing())
	assert.True(t, p.AllowComment())
	p.SetAllowPing(false)
	p.SetAllowComment(false)
	assert.False(t, p.AllowPing())
	assert.False(t, p.AllowComment())
	p.SetAllowPing(true)
	p.SetAllowComment(true)
	assert.True(t, p.AllowPing())
	assert.True(t, p.AllowComment())
}

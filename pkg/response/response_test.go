package response

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResponse(t *testing.T) {
	t.Run("normal init", func(t *testing.T) {
		response1 := New(
			Code(100), Message("message"), Detail("detail"),
			NameCode(100), NodeCode(100),
		)
		assert.Equal(t, int32(100100100), response1.Code())
	})
}

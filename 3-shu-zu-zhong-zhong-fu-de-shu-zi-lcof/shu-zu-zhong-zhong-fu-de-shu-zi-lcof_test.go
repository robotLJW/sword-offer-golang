package __shu_zu_zhong_zhong_fu_de_shu_zi_lcof

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindRepeatNumber(t *testing.T) {
	t.Run("2, 4, 1, 0, 2, 5, 3", func(t *testing.T) {
		data := []int{2, 3, 1, 0, 2, 5, 3}
		ans := findRepeatNumber(data)
		assert.Equal(t, ans, 2)
	})
}

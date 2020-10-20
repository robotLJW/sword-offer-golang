package _4_er_cha_shu_zhong_he_wei_mou_yi_zhi_de_lu_jing_lcof

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_pathSum(t *testing.T) {

	t.Run("", func(t *testing.T) {
		node := &TreeNode{
			Val: 5,
		}
		data := pathSum(node, 5)
		expect := [][]int{
			[]int{5},
		}
		assert.Equal(t, expect, data)

	})

}

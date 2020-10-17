package _5_fu_za_lian_biao_de_fu_zhi_lcof

import (
	"testing"
)

func Test_copyRandomList(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		node1 := &Node{
			Val: 7,
		}
		node2 := &Node{
			Val: 13,
		}
		node3 := &Node{
			Val: 11,
		}
		node4 := &Node{
			Val: 10,
		}
		node5 := &Node{
			Val: 1,
		}
		node1.Next = node2
		node2.Next = node3
		node3.Next = node4
		node4.Next = node5
		node2.Random = node1
		node3.Random = node5
		node4.Random = node3
		node5.Random = node1
		copyRandomList(node1)
	})
}

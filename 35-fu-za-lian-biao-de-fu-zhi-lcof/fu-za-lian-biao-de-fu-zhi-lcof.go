package _5_fu_za_lian_biao_de_fu_zhi_lcof

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// 先拷贝节点处理好next，再指定random,通过遍历random在哪个位置
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	tempHeadNode := head
	copyHead := &Node{
		Val: tempHeadNode.Val,
	}
	currentNode, currentCopyNode := tempHeadNode.Next, copyHead
	for currentNode != nil {
		tempNode := &Node{
			Val: currentNode.Val,
		}
		currentCopyNode.Next = tempNode
		currentCopyNode = tempNode
		currentNode = currentNode.Next
	}
	tempCopyHead := copyHead
	for tempHeadNode != nil {
		randomNode := tempHeadNode.Random
		if randomNode != nil {
			count := 0
			for tempNode := head; tempNode != randomNode; count, tempNode = count+1, tempNode.Next {

			}
			tempNode := copyHead
			for count > 0 {
				tempNode = tempNode.Next
				count--
			}
			tempCopyHead.Random = tempNode
		}
		tempHeadNode = tempHeadNode.Next
		tempCopyHead = tempCopyHead.Next
	}
	return copyHead
}

func copyRandomListTwo(head *Node) *Node {
	maps := make(map[*Node]*Node)
	cur := head
	for cur != nil {
		temp := &Node{
			Val: cur.Val,
		}
		maps[cur] = temp
		cur = cur.Next
	}
	cur = head
	for cur != nil {
		maps[cur].Next = maps[cur.Next]
		maps[cur].Random = maps[cur.Random]
		cur = cur.Next
	}
	return maps[head]
}

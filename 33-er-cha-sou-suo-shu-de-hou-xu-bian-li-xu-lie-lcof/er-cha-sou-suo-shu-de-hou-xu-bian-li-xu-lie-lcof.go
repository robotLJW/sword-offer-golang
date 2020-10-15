package _3_er_cha_sou_suo_shu_de_hou_xu_bian_li_xu_lie_lcof

func verifyPostorder(postorder []int) bool {
	return verify(postorder, 0, len(postorder)-1)
}

func verify(postorder []int, i, j int) bool {
	if i >= j {
		return true
	}
	p := i
	for postorder[p] < postorder[j] {
		p++
	}
	m := p
	for postorder[p] > postorder[j] {
		p++
	}
	return p == j && verify(postorder, i, m-1) && verify(postorder, m, j-1)
}

package lesson8

//树可以看成是一个连通且 无环 的 无向 图。
//给定往一棵 n 个节点 (节点值 1～n) 的树中添加一条边后的图。添加的边的两个顶点包含在 1 到 n 中间，
//且这条附加的边不属于树中已存在的边。图的信息记录于长度为 n 的二维数组 edges ，edges[i] = [ai, bi] 表示图中在 ai 和 bi 之间存在一条边。
//请找出一条可以删去的边，删除后可使得剩余部分是一个有着 n 个节点的树。如果有多个答案，则返回数组 edges 中最后出现的边。
/*parameter
edges: int型二维数组
return: 多余的边
*/
func FindRedundantConnection(edges [][]int) []int {

	father := make([]int, len(edges)+1)
	for i := range father {

		father[i] = i
	}

	var find func(int) int
	find = func(item int) int {

		if father[item] != item {

			father[item] = find(father[item])
		}

		return father[item]
	}

	union := func(a, b int) {

		fa, fb := find(a), find(b)
		father[fa] = fb
	}

	for i := range edges {

		fa, fb := find(edges[i][0]), find(edges[i][1])
		if fa != fb {

			union(edges[i][0], edges[i][1])
		} else {

			return edges[i]
		}
	}

	return nil
}

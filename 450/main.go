package main

import (
	"computerBase/code"
)

func main() {
	//root :=
	//	code.TreeNode{5,
	//		&code.TreeNode{3,
	//			&code.TreeNode{2,nil,nil},
	//			&code.TreeNode{4,nil,nil}},
	//		&code.TreeNode{6,nil,
	//			&code.TreeNode{7,nil,nil}}}
	root :=
		code.TreeNode{5,
			&code.TreeNode{3,
				nil,
				&code.TreeNode{4,nil,nil}},
			&code.TreeNode{7,
				&code.TreeNode{6,nil, nil},
				&code.TreeNode{8,nil,nil}}}
	//root :=
	//	code.TreeNode{8,nil,
	//	&code.TreeNode{15,
	//		&code.TreeNode{13,
	//			&code.TreeNode{12,nil,nil},
	//			&code.TreeNode{14,nil,nil}},
	//	&code.TreeNode{19,
	//		&code.TreeNode{17,nil,
	//			&code.TreeNode{18,nil,nil}},nil}}}
	node := code.DeleteNode(&root, 5)
	println(node)


	newRoot :=
		code.BalanceTreeNode{5,0,
			&code.BalanceTreeNode{2,-1,
				&code.BalanceTreeNode{1,0,nil,nil},
				&code.BalanceTreeNode{3,-1,nil,&code.BalanceTreeNode{4,0,nil,nil}}},
			&code.BalanceTreeNode{6,0,nil,nil}}

	//newRoot = code.LeftRotate(newRoot)

	balance := code.LeftBalance(&newRoot)
	println(balance)

	newRoot =
		code.BalanceTreeNode{2,0,
			&code.BalanceTreeNode{1,0,nil,nil},
			&code.BalanceTreeNode{5,1,
				&code.BalanceTreeNode{4,1,
					&code.BalanceTreeNode{3,0,nil,nil},
					nil},
			&code.BalanceTreeNode{6,0,nil,nil}}}
	balance = code.RightBalance(&newRoot)
	println(balance)
	aTaller := false
	balance = &code.BalanceTreeNode{1,0,nil,nil}
	code.InsertNode(balance, 1, &aTaller)
	code.InsertNode(balance, 2, &aTaller)
	code.InsertNode(balance, 3, &aTaller)
	println(balance)
	println(aTaller)


	//var a int = 10
	//fmt.Printf("变量的地址: %x\n", &a  )
	//fmt.Printf("变量的值: %d\n", a  )
	//var ip *int        /* 声明指针变量 */
	//ip = &a
	//fmt.Printf("a 变量的地址是: %x\n", &a  )
	///* 指针变量的存储地址 */
	//fmt.Printf("ip 变量储存的指针地址: %x\n", ip )
	//fmt.Println(ip)
	///* 使用指针访问值 */
	//fmt.Printf("*ip 变量的值: %d\n", *ip)
	//fmt.Println(*ip == a)
	//fmt.Println(ip == &a)
	//*ip = 20
	//println(a)
	//a = 30
	//println(*ip)
	//
	//var s = "123"
	//var sp *string
	//sp = &s
	//s = "345"
	//print(*sp)
}
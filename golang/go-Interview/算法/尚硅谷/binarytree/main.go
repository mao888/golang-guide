package main
import (
	"fmt"
)

type Hero struct {
	No int
	Name string
	Left *Hero
	Right *Hero
}

//前序遍历[先输root结点，然后再输出左子树，然后再输出右子树]
func PreOrder(node *Hero) {
	if node != nil {
		fmt.Printf("no=%d name=%s\n", node.No, node.Name)
		PreOrder(node.Left)
		PreOrder(node.Right)
	} 
}

//中序遍历[先输出root的左子树，再输root结点，最后输出root的右子树]
func InfixOrder(node *Hero) {
	if node != nil {
		
		InfixOrder(node.Left)
		fmt.Printf("no=%d name=%s\n", node.No, node.Name)
		InfixOrder(node.Right)
	} 
}

func PostOrder(node *Hero) {
	if node != nil {
		
		PostOrder(node.Left)
		PostOrder(node.Right)
		fmt.Printf("no=%d name=%s\n", node.No, node.Name)
	} 
}


func main() {

	//构建一个二叉树
	root := &Hero{
		No : 1,
		Name : "宋江",
	}

	left1 := &Hero{
		No : 2,
		Name : "吴用",
	}

	node10 := &Hero{
		No : 10,
		Name : "tom",
	}

	node12 := &Hero{
		No : 12,
		Name : "jack",
	}
	left1.Left = node10
	left1.Right = node12

	right1 := &Hero{
		No : 3,
		Name : "卢俊义",
	}

	root.Left = left1
	root.Right = right1

	right2 := &Hero{
		No : 4,
		Name : "林冲",
	}
	right1.Right = right2

	//PreOrder(root)
	//InfixOrder(root) //
	PostOrder(root) //后序遍历
}
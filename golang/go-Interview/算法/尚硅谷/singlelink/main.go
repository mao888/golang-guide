package main
import (
	"fmt"
)

//定义一个HeroNode
type HeroNode struct {
	no 		 int
	name 	 string
	nickname string
	next     *HeroNode //这个表示指向下一个结点
}

//给链表插入一个结点
//编写第一种插入方法，在单链表的最后加入.[简单]
func InsertHeroNode(head *HeroNode, newHeroNode *HeroNode) {
	//思路
	//1. 先找到该链表的最后这个结点
	//2. 创建一个辅助结点[跑龙套, 帮忙]
	temp := head
	for {
		if temp.next == nil { //表示找到最后
			break
		}
		temp = temp.next // 让temp不断的指向下一个结点
	}

	//3. 将newHeroNode加入到链表的最后
	temp.next = newHeroNode
}

//给链表插入一个结点
//编写第2种插入方法，根据no 的编号从小到大插入..【实用】
func InsertHeroNode2(head *HeroNode, newHeroNode *HeroNode) {
	//思路
	//1. 找到适当的结点
	//2. 创建一个辅助结点[跑龙套, 帮忙]
	temp := head
	flag := true
	//让插入的结点的no，和temp的下一个结点的no比较
	for {
		if temp.next == nil {//说明到链表的最后
			break
		} else if temp.next.no >= newHeroNode.no {
			//说明newHeroNode 就应该插入到temp后面
			break 
		} else if temp.next.no == newHeroNode.no {
			//说明我们额链表中已经有这个no,就不然插入.
			flag = false
			break
		}
		temp = temp.next
	}

	if !flag {
		fmt.Println("对不起，已经存在no=", newHeroNode.no)
		return
	} else {
		newHeroNode.next = temp.next
		temp.next = newHeroNode
	}


}
//删除一个结点
func DelHerNode(head *HeroNode, id int) {
	temp := head
	flag := false
	//找到要删除结点的no，和temp的下一个结点的no比较
	for {
		if temp.next == nil {//说明到链表的最后
			break
		} else if temp.next.no == id {
			//说明我们找到了.
			flag = true
			break
		}
		temp = temp.next
	}
	if flag {//找到, 删除

		temp.next = temp.next.next
	} else {
		fmt.Println("sorry, 要删除的id不存在")
	}
}

//显示链表的所有结点信息
func ListHeroNode(head *HeroNode) {

	//1. 创建一个辅助结点[跑龙套, 帮忙]
	temp := head

	// 先判断该链表是不是一个空的链表
	if temp.next == nil {
		fmt.Println("空空如也。。。。")
		return 
	}
	//2. 遍历这个链表
	for {
		fmt.Printf("[%d , %s , %s]==>", temp.next.no, 
			temp.next.name, temp.next.nickname)
		//判断是否链表后
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
}




func main() {

	//1. 先创建一个头结点, 
	head := &HeroNode{}

	//2. 创建一个新的HeroNode
	hero1 := &HeroNode{
		no : 1,
		name : "宋江",
		nickname : "及时雨",
	}

	hero2 := &HeroNode{
		no : 2,
		name : "卢俊义",
		nickname : "玉麒麟",
	}

	hero3 := &HeroNode{
		no : 3,
		name : "林冲",
		nickname : "豹子头",
	}

	// hero4 := &HeroNode{
	// 	no : 3,
	// 	name : "吴用",
	// 	nickname : "智多星",
	// }

	//3. 加入
	InsertHeroNode2(head, hero3)
	InsertHeroNode2(head, hero1)
	InsertHeroNode2(head, hero2)
	
	// 4. 显示
	ListHeroNode(head)

	// 5 删除
	fmt.Println()
	DelHerNode(head, 1)
	DelHerNode(head, 3)
	ListHeroNode(head)
}
	
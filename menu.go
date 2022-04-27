package main

import "fmt"

//定义链表结点
type LinkNode struct {
	data interface{}
	next *LinkNode
}

//定义链表
type LinkTable struct {
	length int
	head   *LinkNode
	tail   *LinkNode
}

//初始化LinkNode
func InitLinkNode(data interface{}) *LinkNode {
	return &LinkNode{
		data: data,
		next: nil,
	}
}

//初始化LinkTable
func InitLinkTable() *LinkTable {
	return &LinkTable{
		length: 0,
		head:   nil,
		tail:   nil,
	}
}

//判断链表是否为空
func IsEmpty(pLink *LinkTable) bool {
	return pLink.length == 0
}

//链表头部添加结点
func AddToHead(pLink *LinkTable, node *LinkNode) bool {
	if pLink == nil || node == nil {
		return false
	}
	if IsEmpty(pLink) {
		pLink.tail = node
	} else {
		node.next = pLink.head
	}
	pLink.head = node
	pLink.length++
	return true
}

//链表尾部添加结点
func AddToTail(pLink *LinkTable, node *LinkNode) bool {
	if pLink == nil || node == nil {
		return false
	}
	if IsEmpty(pLink) {
		pLink.head = node
	} else {
		pLink.tail.next = node
	}
	pLink.tail = node
	pLink.length++
	return true
}

//删除对应值的第一个出现的结点
func DeleteFirstValueNode(pLink *LinkTable, data interface{}) bool {
	if pLink == nil {
		return false
	}
	if pLink.head.data == data {
		if pLink.length == 1 {
			pLink.head = nil
			pLink.tail = nil
		} else {
			pLink.head = pLink.head.next
		}
		pLink.length--
		return true
	}
	var p *LinkNode = pLink.head
	for p != pLink.tail {
		if p.next.data == data {
			p.next = p.next.next
			pLink.length--
			if p.next == nil {
				pLink.tail = p
			}
			return true
		}
		p = p.next
	}
	return false
}

//查找对应值的结点
func SelectLinkNode(pLink *LinkTable, data interface{}) *LinkNode {
	if pLink == nil {
		return nil
	}
	var p *LinkNode = pLink.head
	for p != nil {
		if p.data == data {
			return p
		}
		p = p.next
	}
	return nil
}

//对链表遍历，并显示结点对应值
func ShowLinkTableValues(pLink *LinkTable) {
	var p *LinkNode = pLink.head
	if p == nil {
		fmt.Println("length=", pLink.length, " 链表为空")
		return
	}
	fmt.Print("length=", pLink.length, " ")
	for p != nil {
		fmt.Print(p.data, " ")
		p = p.next
	}
	fmt.Println()
}

func main() {
	pLinkTable := InitLinkTable()
	ShowLinkTableValues(pLinkTable)
	node1 := InitLinkNode("abc")
	AddToHead(pLinkTable, node1)
	AddToHead(pLinkTable, InitLinkNode("xas"))
	AddToTail(pLinkTable, InitLinkNode("lso"))
	AddToTail(pLinkTable, InitLinkNode("2021"))
	AddToHead(pLinkTable, InitLinkNode("1"))
	ShowLinkTableValues(pLinkTable)

	DeleteFirstValueNode(pLinkTable, "2")
	DeleteFirstValueNode(pLinkTable, "1")
	ShowLinkTableValues(pLinkTable)

	fmt.Println("查找值为2020的结点:", SelectLinkNode(pLinkTable, "2020"))
	fmt.Println("查找值为2021的结点:", SelectLinkNode(pLinkTable, "2021"))
	fmt.Println("查找值为xas的结点:", SelectLinkNode(pLinkTable, "xas"))
	DeleteFirstValueNode(pLinkTable, "xas")
	DeleteFirstValueNode(pLinkTable, "2021")
	ShowLinkTableValues(pLinkTable)
}

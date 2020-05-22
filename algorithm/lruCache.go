package leetcodes

type LRUCache struct {
	mapKV    map[int]*Node
	MaxSize  int
	realSize int
	head     *Node
}
type Node struct {
	key   int
	value int
	count int
	right *Node
}

//func Constructor(capacity int) LRUCache {
//	return LRUCache{MaxSize:capacity,realSize:0,mapKV:make(map[int]*Node,capacity),head:nil,tail:nil}
//}
//
//
//func (this *LRUCache) Get(key int) int {
//	if node,ok := this.mapKV[key];ok{
//		node.count++
//		return node.value
//	}
//	return 0
//}
//
//
//func (this *LRUCache) Put(key int, value int)  {
//	if this.realSize == 0 {
//		node := &Node{key:key,value:value,count:1}
//		this.head = node
//		this.realSize = 1
//		this.mapKV[key] = node
//		return
//	}
//	if node,ok := this.mapKV[key];ok{
//		node.value = value
//		node.count++
//		if node.left == nil {
//			return
//		}
//		tmpNode := node.left
//		for tmpNode != nil && tmpNode.count < node.count {
//			tmpNode = tmpNode.left
//		}
//		if tmpNode == nil {
//			node.right.left,node.left.right = node.left,node.right
//			node.left,node.right,this.head.left = nil,this.head,node
//			this.head = node
//		}else{
//			node.left,tmpNode.left = tmpNode.left,node.left
//			node.right,tmpNode.right = tmpNode.right,node.right
//		}
//		//排序
//	}else {
//		node := &Node{key:key,value:value,count:1}
//		if this.realSize+1 > this.MaxSize {
//			tmp := this.tail.left
//			this.tail.left.right = node
//			node.left = tmp
//			this.tail = node
//		}else {
//			node.left = this.tail
//			this.tail.right=node
//			this.tail = node
//			this.realSize++
//		}
//		this.mapKV[key] = node
//	}
//}
//
//func (this *LRUCache)Print()  {
//	println("++++++++++++")
//	node := this.head
//	for node != nil {
//		println(node.key,node.value,node.count)
//		if node.left != nil {
//			node = node.left
//		}else {
//			break
//		}
//	}
//	println("++++++++++++")
//}

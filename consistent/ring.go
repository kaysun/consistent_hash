package consistent

import (
	"sort"
)

// HashRing 哈希环
type HashRing struct {
	// virtualNodeList 虚拟节点列表
	virtualNodeList VirtualNodeSort
	// physicalNodeList 物理节点列表
	physicalNodeList []*PhysicalNode
}

// NewRing 创建哈希环
func NewRing(ipList []string, virtualNodeCountPerPhysicalNode int) *HashRing {
	ring := &HashRing{}
	for _, ip := range ipList {
		ring.AddPhysicalNode(ip, virtualNodeCountPerPhysicalNode)
	}
	//将哈希环上的虚拟节点按照hash值排序
	sort.Sort(ring.virtualNodeList)
	return ring
}

// AddPhysicalNode 向hash环中添加物理机
func (ring *HashRing) AddPhysicalNode(ip string, virtualNodeCount int) {
	// 创建物理节点
	physicalNode := &PhysicalNode{IP: ip}
	// 设置虚拟节点个数
	physicalNode.SetVirtualNodeCount(virtualNodeCount)
	for j := 0; j < virtualNodeCount; j++ {
		// 创建虚拟节点
		virtualNode := NewVirtualNode(ip, j)
		// 设置虚拟节点的物理节点
		virtualNode.PhysicalNode = physicalNode
		// 物理节点添加虚拟节点
		physicalNode.AddVirtualNode(virtualNode)
		// 哈希环维护虚拟节点列表
		ring.virtualNodeList = append(ring.virtualNodeList, virtualNode)
	}
	// 哈希环维护物理节点列表
	ring.physicalNodeList = append(ring.physicalNodeList, physicalNode)
}

// PutKeyVal 插入kv数据
func (ring *HashRing) PutKeyVal(key string, value interface{}) {
	// 取key的hash值
	hashCode := hash(key)
	// 顺时针查找第一个大于等于key的哈希值的虚拟节点
	virtualNode := ring.FindVirtualNodeByClockwise(hashCode)
	// 将key和value写入该虚拟节点
	virtualNode.PutKeyValue(key, value)
}

// FindVirtualNodeByClockwise 顺时针查找第一个大于等于key的哈希值的虚拟节点
func (ring *HashRing) FindVirtualNodeByClockwise(keyHash uint32) *VirtualNode {
	var last *VirtualNode
	// 遍历hash环的虚拟节点
	for i, node := range ring.virtualNodeList {
		// 若没有比keyHash大的节点，则说明这个key会对应到环上的第0个节点
		if i == 0 {
			last = node
		}
		// 当找到第一个哈希值大于keyHash的节点，则返回上一个节点
		if node.HashCode > keyHash {
			break
		}
		last = node
	}
	return last
}

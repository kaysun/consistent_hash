package consistent

import (
	"fmt"
	"sync"
)

// VirtualNode 虚拟节点
type VirtualNode struct {
	// PhysicalNode 虚拟节点对应的物理节点
	PhysicalNode *PhysicalNode
	// HashCode 哈希值
	HashCode uint32
	// 虚拟节点的ip:port
	VirtualIPPort string

	// data kv数据，不可导出
	data map[string]interface{}
	// sync.RWMutex 读写锁
	sync.RWMutex
}

// NewVirtualNode 自己造构造函数，用于初始化map。创建虚拟节点必须使用NewVirtualNode，否则调用方法panic，后果自负哦～
func NewVirtualNode(ip string, seq int) *VirtualNode {
	ipport := fmt.Sprintf("%s:%s", ip, getNonce())
	return &VirtualNode{
		data: make(map[string]interface{}),
		// 初始化时设置随机hash值
		//HashCode:randHashValue(),
		// 模拟一个虚拟节点的ip
		VirtualIPPort: ipport,
		// 设置虚拟节点的哈希值
		HashCode: hash(ipport),
	}
}

// PutKeyValue 设置kv
func (node *VirtualNode) PutKeyValue(key string, value interface{}) {
	node.Lock()
	defer node.Unlock()
	// 无论是否存在该key，直接设值即可
	node.data[key] = value
}

// GetKeyCount 获取虚拟节点的key的个数
func (node *VirtualNode) GetKeyCount() int {
	return len(node.data)
}

// VirtualNodeSort 按照hash值排序的VirtualNode
type VirtualNodeSort []*VirtualNode

// Len 实现sort接口
func (sortNode VirtualNodeSort) Len() int {
	return len(sortNode)
}

// Less 实现sort接口
func (sortNode VirtualNodeSort) Less(i, j int) bool {
	return sortNode[i].HashCode <= sortNode[j].HashCode
}

// Swap 实现sort接口
func (sortNode VirtualNodeSort) Swap(i, j int) {
	sortNode[i], sortNode[j] = sortNode[j], sortNode[i]
}

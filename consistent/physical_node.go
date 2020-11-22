package consistent

// PhysicalNode 物理节点
type PhysicalNode struct {
	// VirtualNodeList 物理节点对应的虚拟节点列表
	VirtualNodeList []*VirtualNode
	// virtualNodeCount 每个物理节点对应的虚拟节点的个数
	virtualNodeCount int
	// IP ip地址，若存数据库，则应该采用int类型，若只是内存中使用，string类型也可
	IP string
}

// AddVirtualNode 向物理节点中添加虚拟节点
func (node *PhysicalNode) AddVirtualNode(virtualNode *VirtualNode) {
	node.VirtualNodeList = append(node.VirtualNodeList, virtualNode)
}

// BatchAddVirtualNodes 向物理节点中批量添加虚拟节点
func (node *PhysicalNode) BatchAddVirtualNodes(virtualNodes []*VirtualNode) {
	node.VirtualNodeList = append(node.VirtualNodeList, virtualNodes...)
}

// GetAllKeyCount 获取物理节点下所有的key的个数
func (node *PhysicalNode) GetAllKeyCount() int {
	sum := 0
	// 遍历物理节点上所有的虚拟节点，获取每个虚拟节点里的key的个数，加和即为结果
	for _, virtualNode := range node.VirtualNodeList {
		sum += virtualNode.GetKeyCount()
	}
	return sum
}

// SetVirtualNodeCount 设置每个物理节点对应的虚拟节点的个数
func (node *PhysicalNode) SetVirtualNodeCount(virtualNodeCount int) {
	node.virtualNodeCount = virtualNodeCount
}

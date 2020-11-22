package consistent

import (
	"fmt"
	"math"
)

//标准差公式意义 所有数减去其平均值的平方和，
// 所得结果除以该组数之个数（或个数减一，即变异数），
// 再把所得值开根号，所得之数就是这组数据的标准差。
// 标准差越小，离散度越好

// StandardDeviation 标准差
type StandardDeviation struct {
	Ring *HashRing
}

//编写测试用例测试这个算法，测试 100 万 KV 数据，10 个服务器节点的情况下，
// 计算这些 KV 数据在服务器上分布数量的标准差，以评估算法的存储负载不均衡性。

// CalculationStandardDeviation 计算标准差
func (sd StandardDeviation) CalculationStandardDeviation() float64 {
	len := sd.AllKeyCount()
	if len <= 1 {
		return 0
	}
	avg := sd.AvgKeyCountOfPhysicalNode()
	var sum float64
	for _, node := range sd.Ring.physicalNodeList {
		// 所有数减去其平均值的平方和
		sum += math.Pow(float64(node.GetAllKeyCount())-avg, float64(2))
	}
	// 样本标准差=方差的算术平方根=s=sqrt(((x1-x)^2 +(x2-x)^2 +......(xn-x)^2) /（n-1）)
	standardDeviation := math.Sqrt(sum / float64(len-1))
	return standardDeviation
}

// AllKeyCount hash环上所有的key的个数
func (sd StandardDeviation) AllKeyCount() int {
	sum := 0
	for _, node := range sd.Ring.physicalNodeList {
		sum += node.GetAllKeyCount()
	}
	return sum
}

// AvgKeyCountOfPhysicalNode hash环上每个物理节点平均的key的个数
func (sd StandardDeviation) AvgKeyCountOfPhysicalNode() float64 {
	sum := sd.AllKeyCount()
	len := len(sd.Ring.physicalNodeList)
	if len == 0 {
		return 0
	}
	return float64(sum / len)
}

// PrintKeyDistributedDetail 打印细节
func (sd StandardDeviation) PrintKeyDistributedDetail() {
	all := sd.AllKeyCount()
	for _, node := range sd.Ring.physicalNodeList {
		allPerPhysicalNode := node.GetAllKeyCount()
		fmt.Println(fmt.Sprintf("ip=%s, key count=%d, rate=%f, all = %d",
			node.IP, allPerPhysicalNode, float64(allPerPhysicalNode)/float64(all), all))
	}
}

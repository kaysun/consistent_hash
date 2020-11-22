package consistent

import (
	"fmt"
	"math/rand"
)

var IPList = []string{
	"198.168.0.200",
	"198.168.0.201",
	"198.168.0.202",
	"198.168.0.203",
	"198.168.0.204",
	"198.168.0.205",
	"198.168.0.206",
	"198.168.0.207",
	"198.168.0.208",
	"198.168.0.209",
	"198.168.0.210",
	"198.168.0.211",
	"198.168.0.212",
	"198.168.0.213",
	"198.168.0.214",
	"198.168.0.215",
	"198.168.0.216",
	"198.168.0.217",
	"198.168.0.218",
	"198.168.0.219",
}

// 计算标准差
func CalculationStandardDeviation(kvCount, virtualNodeCountPerPhysicalNode int, ipList []string) float64 {
	// 创建哈希环对象
	ring := NewRing(ipList, virtualNodeCountPerPhysicalNode)
	// 添加kvCount个数据
	for i := 0; i < kvCount; i++ {
		// 使用空结构体，减少内存占用
		ring.PutKeyVal(fmt.Sprintf("%s-%d", getNonce(), i), struct{}{})
	}
	// 创建标准差对象
	sd := StandardDeviation{Ring: ring}
	// 打印详情
	sd.PrintKeyDistributedDetail()
	// 计算标准差
	return sd.CalculationStandardDeviation()
}

// getNonce 获取随机字符串
func getNonce() string {
	return func(len int) string {
		// 种子需要使用随机值，如果只是时间戳，因为每毫秒会调用很多次，所以导致随机字符串是一样的。
		rand.Seed(rand.Int63())
		data := make([]byte, len)
		for i := 0; i < len; i++ {
			num := rand.Intn(57) + 65
			for {
				if num > 90 && num < 97 {
					num = rand.Intn(57) + 65
				} else {
					break
				}
			}
			data[i] = byte(num)
		}
		return string(data)
	}(6)
}

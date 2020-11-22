package consistent

import (
	"hash/fnv"
)

// hash值，返回0 - 2的32次幂-1（Range: 0 through 4294967295）
func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

//// randHashValue 获取一个随机哈希值
//func randHashValue() uint32 {
//	// 采用时间戳作为种子
//	rand.Seed(time.Now().Unix())
//	// 生成-2的32次幂 ～ 2的32次幂-1（Range: -4294967296 through 4294967295）的随机数
//	num := rand.Int63n(math.MaxUint32)
//	// 取绝对值
//	abs := uint32(math.Abs(float64(num)))
//	fmt.Println(fmt.Sprintf("虚拟节点的hash code=%d", abs))
//	return abs
//}

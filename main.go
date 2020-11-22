package main

import (
	"fmt"

	. "git.code.oa.com/kaysun/consistent_hash/consistent"
)

func main() {
	//ip=198.168.0.200, key count=108322, rate=0.108322, all = 1000000
	//ip=198.168.0.201, key count=100114, rate=0.100114, all = 1000000
	//ip=198.168.0.202, key count=109305, rate=0.109305, all = 1000000
	//ip=198.168.0.203, key count=98366, rate=0.098366, all = 1000000
	//ip=198.168.0.204, key count=95997, rate=0.095997, all = 1000000
	//ip=198.168.0.205, key count=97605, rate=0.097605, all = 1000000
	//ip=198.168.0.206, key count=103851, rate=0.103851, all = 1000000
	//ip=198.168.0.207, key count=99153, rate=0.099153, all = 1000000
	//ip=198.168.0.208, key count=96105, rate=0.096105, all = 1000000
	//ip=198.168.0.209, key count=91182, rate=0.091182, all = 1000000
	//consistent10_150 StandardDeviation=16.992873

	fmt.Println(fmt.Sprintf("consistent15_150 StandardDeviation=%f",
		CalculationStandardDeviation(1000000, 150, IPList[0:10])))
}

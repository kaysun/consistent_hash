package consistent

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	t.Run("15-150: ", consistent15150)
	//t.Run("20-150: ", consistent20150)
	//t.Run("20-200: ", consistent20200)
	//t.Run("10-200: ", consistent10200)
	//t.Run("10-150: ", consistent10150)
}

//ip=198.168.0.200, key count=67262, rate=0.067262, all = 1000000
//ip=198.168.0.201, key count=68769, rate=0.068769, all = 1000000
//ip=198.168.0.202, key count=72714, rate=0.072714, all = 1000000
//ip=198.168.0.203, key count=68305, rate=0.068305, all = 1000000
//ip=198.168.0.204, key count=64810, rate=0.064810, all = 1000000
//ip=198.168.0.205, key count=72714, rate=0.072714, all = 1000000
//ip=198.168.0.206, key count=67588, rate=0.067588, all = 1000000
//ip=198.168.0.207, key count=69259, rate=0.069259, all = 1000000
//ip=198.168.0.208, key count=57175, rate=0.057175, all = 1000000
//ip=198.168.0.209, key count=57740, rate=0.057740, all = 1000000
//ip=198.168.0.210, key count=65829, rate=0.065829, all = 1000000
//ip=198.168.0.211, key count=64333, rate=0.064333, all = 1000000
//ip=198.168.0.212, key count=75972, rate=0.075972, all = 1000000
//ip=198.168.0.213, key count=67457, rate=0.067457, all = 1000000
//ip=198.168.0.214, key count=60073, rate=0.060073, all = 1000000
//consistent15_150 StandardDeviation=19.955716
func consistent15150(t *testing.T) {
	fmt.Println(fmt.Sprintf("consistent15_150 StandardDeviation=%f",
		CalculationStandardDeviation(1000000, 150, IPList[0:15])))
}

//ip=198.168.0.200, key count=52555, rate=0.052555, all = 1000000
//ip=198.168.0.201, key count=49782, rate=0.049782, all = 1000000
//ip=198.168.0.202, key count=56026, rate=0.056026, all = 1000000
//ip=198.168.0.203, key count=55435, rate=0.055435, all = 1000000
//ip=198.168.0.204, key count=49494, rate=0.049494, all = 1000000
//ip=198.168.0.205, key count=52723, rate=0.052723, all = 1000000
//ip=198.168.0.206, key count=50220, rate=0.050220, all = 1000000
//ip=198.168.0.207, key count=47975, rate=0.047975, all = 1000000
//ip=198.168.0.208, key count=44843, rate=0.044843, all = 1000000
//ip=198.168.0.209, key count=43520, rate=0.043520, all = 1000000
//ip=198.168.0.210, key count=47008, rate=0.047008, all = 1000000
//ip=198.168.0.211, key count=48353, rate=0.048353, all = 1000000
//ip=198.168.0.212, key count=56167, rate=0.056167, all = 1000000
//ip=198.168.0.213, key count=50336, rate=0.050336, all = 1000000
//ip=198.168.0.214, key count=49621, rate=0.049621, all = 1000000
//ip=198.168.0.215, key count=49463, rate=0.049463, all = 1000000
//ip=198.168.0.216, key count=50234, rate=0.050234, all = 1000000
//ip=198.168.0.217, key count=47128, rate=0.047128, all = 1000000
//ip=198.168.0.218, key count=50227, rate=0.050227, all = 1000000
//ip=198.168.0.219, key count=48890, rate=0.048890, all = 1000000
//consistent15_150 StandardDeviation=14.582892
func consistent20150(t *testing.T) {
	fmt.Println(fmt.Sprintf("consistent15_150 StandardDeviation=%f",
		CalculationStandardDeviation(1000000, 150, IPList)))
}

//ip=198.168.0.200, key count=51959, rate=0.051959, all = 1000000
//ip=198.168.0.201, key count=47697, rate=0.047697, all = 1000000
//ip=198.168.0.202, key count=50713, rate=0.050713, all = 1000000
//ip=198.168.0.203, key count=49588, rate=0.049588, all = 1000000
//ip=198.168.0.204, key count=52011, rate=0.052011, all = 1000000
//ip=198.168.0.205, key count=49301, rate=0.049301, all = 1000000
//ip=198.168.0.206, key count=48135, rate=0.048135, all = 1000000
//ip=198.168.0.207, key count=51111, rate=0.051111, all = 1000000
//ip=198.168.0.208, key count=50059, rate=0.050059, all = 1000000
//ip=198.168.0.209, key count=50678, rate=0.050678, all = 1000000
//ip=198.168.0.210, key count=55181, rate=0.055181, all = 1000000
//ip=198.168.0.211, key count=48093, rate=0.048093, all = 1000000
//ip=198.168.0.212, key count=47573, rate=0.047573, all = 1000000
//ip=198.168.0.213, key count=48060, rate=0.048060, all = 1000000
//ip=198.168.0.214, key count=55887, rate=0.055887, all = 1000000
//ip=198.168.0.215, key count=52057, rate=0.052057, all = 1000000
//ip=198.168.0.216, key count=45683, rate=0.045683, all = 1000000
//ip=198.168.0.217, key count=49110, rate=0.049110, all = 1000000
//ip=198.168.0.218, key count=50273, rate=0.050273, all = 1000000
//ip=198.168.0.219, key count=46831, rate=0.046831, all = 1000000
//consistent20_200 StandardDeviation=11.317915
func consistent20200(t *testing.T) {
	fmt.Println(fmt.Sprintf("consistent20_200 StandardDeviation=%f",
		CalculationStandardDeviation(1000000, 200, IPList)))
}

//ip=198.168.0.200, key count=91650, rate=0.091650, all = 1000000
//ip=198.168.0.201, key count=99018, rate=0.099018, all = 1000000
//ip=198.168.0.202, key count=101388, rate=0.101388, all = 1000000
//ip=198.168.0.203, key count=99150, rate=0.099150, all = 1000000
//ip=198.168.0.204, key count=104893, rate=0.104893, all = 1000000
//ip=198.168.0.205, key count=94156, rate=0.094156, all = 1000000
//ip=198.168.0.206, key count=105702, rate=0.105702, all = 1000000
//ip=198.168.0.207, key count=108369, rate=0.108369, all = 1000000
//ip=198.168.0.208, key count=99606, rate=0.099606, all = 1000000
//ip=198.168.0.209, key count=96068, rate=0.096068, all = 1000000
//consistent10_200 StandardDeviation=15.798694
func consistent10200(t *testing.T) {
	fmt.Println(fmt.Sprintf("consistent10_200 StandardDeviation=%f",
		CalculationStandardDeviation(1000000, 200, IPList[0:10])))
}

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
func consistent10150(t *testing.T) {
	fmt.Println(fmt.Sprintf("consistent10_150 StandardDeviation=%f",
		CalculationStandardDeviation(1000000, 150, IPList[0:10])))
}

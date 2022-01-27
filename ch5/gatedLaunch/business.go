package gatedLaunch

import (
	"math/rand"
	"time"
)

//var cityID2Open = [12000]bool{}
func readConfig() {}

func init() {
	readConfig()
	//for i := 0; i < len(cityID2Open); i++ {
	//if city i is opened in configs {
	//	cityID2Open[i] = true
	//}
	//}
}

//
//func isPassed(cityID int) bool {
//	return cityID2Open[cityID]
//}

var cityID2Open = map[int]struct{}{}

func init() {
	readConfig()
	//for _, city := range openCities {
	//	cityID2Open[city] = struct{}{}
	//}
}

//func isPassed(cityID int) bool {
//	if _, ok := cityID2Open[cityID]; ok {
//		return true
//	}
//
//	return false
//}

// 概率

func init() {
	rand.Seed(time.Now().UnixNano())
}

// rate 为 0~100
func isPassed(rate int) bool {
	if rate >= 100 {
		return true
	}

	if rate > 0 && rand.Int() > rate {
		return true
	}

	return false
}

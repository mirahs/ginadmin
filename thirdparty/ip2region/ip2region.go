package ip2region

import (
	"fmt"
	"github.com/lionsoul2014/ip2region/binding/golang/ip2region"
)


var Ip *ip2region.Ip2Region


func RegionInit(dbFile string)  {
	var err error
	Ip, err = ip2region.New(dbFile)
	if err != nil {
		panic("ip2region.RegionInit New err:" + err.Error())
	}
}

func GetAddress(ip string) string {
	ipInfo, _ := Ip.MemorySearch(ip)
	if ipInfo.Country == "" {
		return ""
	}
	if ipInfo.Province == "0" {
		return ipInfo.Country
	}
	if ipInfo.City == "0" {
		return fmt.Sprintf("%s-%s", ipInfo.Country, ipInfo.Province)
	}
	if ipInfo.ISP == "0" {
		return fmt.Sprintf("%s-%s-%s", ipInfo.Country, ipInfo.Province, ipInfo.City)
	}
	return fmt.Sprintf("%s-%s-%s-%s", ipInfo.Country, ipInfo.Province, ipInfo.City, ipInfo.ISP)
}

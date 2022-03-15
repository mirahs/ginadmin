package util

import (
	"fmt"
	"github.com/lionsoul2014/ip2region/binding/golang/ip2region"
)


var region *ip2region.Ip2Region


// 初始化 Ip2Region
func IpInit(dbFile string)  {
	var err error
	region, err = ip2region.New(dbFile)
	if err != nil {
		panic("ip2region.IpInit New err:" + err.Error())
	}
}

// 关闭资源
func IpClose() {
	region.Close()
}

// 根据 ip 获取区域地址信息
func IpAddress(ip string) string {
	ipInfo, _ := region.MemorySearch(ip)

	if ipInfo.Country == "0" && ipInfo.ISP != "" {
		return ipInfo.ISP
	}
	if ipInfo.Country == "" {
		return "内网地址或ipv6"
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

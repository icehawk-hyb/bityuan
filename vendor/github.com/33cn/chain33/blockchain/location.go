package blockchain

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

/*
{
    "code":0,
    "data":{
        "ip":"114.232.69.81",
        "country":"中国",
        "area":"",
        "region":"江苏",
        "city":"南通",
        "county":"XX",
        "isp":"电信",
        "country_id":"CN",
        "area_id":"",
        "region_id":"320000",
        "city_id":"320600",
        "county_id":"xx",
        "isp_id":"100017"
    }
}
*/
type IPInfo struct {
	Code int         `json:"code"`
	Data interface{} `json:"data`
}

type IP struct {
	Country   string `json:"country"`
	CountryId string `json:"country_id"`
	Area      string `json:"area"`
	AreaId    string `json:"area_id"`
	Region    string `json:"region"`
	RegionId  string `json:"region_id"`
	City      string `json:"city"`
	CityId    string `json:"city_id"`
	Isp       string `json:"isp"`
	Point     point  `json:"point"`
}

type BaiDuIp struct {
	Address string  `json:"address"`
	Content content `json:"content"`
}

type content struct {
	AddressDetail interface{} `json:"address_detail"`
	Address       string      `json:"address"`
	Point         point       `json:"point"`
}

type point struct {
	X string `json:"x"`
	Y string `json:"y"`
}

func LocationAPI(ip string) *IPInfo {
	url := "http://ip.taobao.com/service/getIpInfo.php?ip="
	url += ip
	//log.Info("LocatoinAPi", "url", url)
	resp, err := http.Get(url)
	if err != nil {
		chainlog.Error("http get", "err", err.Error())
		return nil
	}
	defer resp.Body.Close()

	out, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		chainlog.Error("io read", "errr", err.Error())
		return nil
	}
	//log.Info("read", "out", string(out))
	var result IPInfo
	if err := json.Unmarshal(out, &result); err != nil {
		chainlog.Error("json umarshal", "err", err.Error())
		return nil
	}

	return &result
}

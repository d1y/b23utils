// 参考: https://github.com/mrhso/IshisashiWebsite
// `bilibili` 视频 `avid` | `bvid` 小工具

package b23utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

var tableCtx = "fZodR9XQDSUm21yCkr6zBqiveYah8bt4xsWpHnJE7jL5VG3guMTKNPAwcF"
var table = []string{}
var resultCtx = "bv1  4 1 7  "
var result = []string{}

var s = []int64{11, 10, 3, 8, 4, 6}

var xor int64 = 177451812
var add int64 = 8728348608

var offset int64 = 58

var tempArr = [6]int64{}

// BilibiliJSON 自动生成: http://json2struct.mervine.net/
type BilibiliJSON struct {
	Code int64 `json:"code"`
	Data struct {
		Aid        int64  `json:"aid"`
		ArgueMsg   string `json:"argue_msg"`
		Bvid       string `json:"bvid"`
		Coin       int64  `json:"coin"`
		Copyright  int64  `json:"copyright"`
		Danmaku    int64  `json:"danmaku"`
		Evaluation string `json:"evaluation"`
		Favorite   int64  `json:"favorite"`
		HisRank    int64  `json:"his_rank"`
		Like       int64  `json:"like"`
		NoReprint  int64  `json:"no_reprint"`
		NowRank    int64  `json:"now_rank"`
		Reply      int64  `json:"reply"`
		Share      int64  `json:"share"`
		View       int64  `json:"view"`
	} `json:"data"`
	Message string `json:"message"`
	TTL     int64  `json:"ttl"`
}

// Bv2avByAPI `bvid` 转为 `avid`
// (不推荐使用该方法, 因为不知道官方的这个接口何时挂掉..)
func Bv2avByAPI(bvid string) string {
	var api = fmt.Sprintf("https://api.bilibili.com/x/web-interface/archive/stat?bvid=%v", bvid)
	resp, err := http.Get(api)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var respBody BilibiliJSON
	json.Unmarshal(body, &respBody)
	if respBody.Code != 0 {
		return ""
	}
	id := respBody.Data.Aid
	var r = fmt.Sprintf("av%v", id)
	return r
}

// Bv2av `bvid` 转为 `avid`
func Bv2av(bvid string) string {
	var bvidLen = len(bvid)
	id := ""
	if bvidLen == 12 {
		id = bvid
	} else if bvidLen == 10 {
		id = fmt.Sprintf("BV%v", bvid)
	} else if bvidLen == 9 {
		id = fmt.Sprintf("BV1%v", bvid)
	} else {
		return ""
	}
	var r int64 = 0
	for i := 0; i < len(tempArr); i++ {
		var _i = s[i]
		var _s = id[_i]
		var _c = indexOf(string(_s), table)
		var _ii = float64(i)
		var _ss = float64(offset)
		var _u = math.Pow(_ss, _ii)
		var u = int64(_c) * int64(_u)
		r += u
	}
	var next = r - add ^ xor
	var _next = fmt.Sprintf("av%v", next)
	return _next
}

// Av2bv `avid` 转为 `bvid`
func Av2bv(avid string) string {
	n := avid
	var isAvWord = strings.Contains(avid, "av")
	if isAvWord {
		n = avid[2:]
	}
	avn, err := strconv.ParseInt(n, 10, 64)
	if err != nil {
		return ""
	}
	var u = (avn ^ xor) + add
	var _result = result
	for i := 0; i < len(tempArr); i++ {
		var _a = float64(i)
		var _b = float64(offset)
		var _i = math.Pow(_b, _a)
		var _c = u / int64(_i)
		var _d = _c % offset
		var _e = table[_d]
		_result[s[i]] = _e
	}
	var R = strings.Join(_result, "")
	return R
}

// FullURL 生成完整的 `url`
func FullURL(id string) string {
	var r = fmt.Sprintf("https://www.bilibili.com/video/%v", id)
	return r
}

// https://github.com/heapwolf/go-indexof
func indexOf(params ...interface{}) int {
	v := reflect.ValueOf(params[0])
	arr := reflect.ValueOf(params[1])

	var t = reflect.TypeOf(params[1]).Kind()

	if t != reflect.Slice && t != reflect.Array {
		panic("Type Error! Second argument must be an array or a slice.")
	}

	for i := 0; i < arr.Len()-1; i++ {
		if arr.Index(i).Interface() == v.Interface() {
			return i
		}
	}
	return -1
}

func init() {
	table = strings.Split(tableCtx, "")
	result = strings.Split(resultCtx, "")
}

package main

import (
	"fmt"

	"github.com/d1y/b23utils"
)

func main() {
	var bvid = "BV1kZ4y1u7Mg"
	var avid = "av15121849"
	var _avid = b23utils.Bv2av(bvid)
	var _bvid = b23utils.Av2bv(avid)
	var fullURL = b23utils.FullURL(_bvid)
	fmt.Println(fullURL)
	fmt.Println(_avid)
	fmt.Println(_bvid)
}

package main

import (
	"fmt"
)

type IPAddr [4]byte

func ConvertIPAdress(_map map[string]IPAddr) string {
	var src string
	for i, _ := range _map {
		if src == "" {
			src = fmt.Sprintf("%v%v: %v.%v.%v.%v", src, i, _map[i][0], _map[i][1], _map[i][2], _map[i][3])
		} else {
			src = fmt.Sprintf("%v\n%v: %v.%v.%v.%v", src, i, _map[i][0], _map[i][1], _map[i][2], _map[i][3])
		}

	}
	return src
}

// func main() {
// 	hosts := map[string]IPAddr{
// 		"loopback":  {127, 0, 0, 1},
// 		"googleDNS": {8, 8, 8, 8},
// 	}
// 	fmt.Println(ConvertIPAdress(hosts))
// 	// for i, _ := range hosts {
// 	// 	fmt.Printf("%v: %v\n", i, hosts[i])
// 	// }

// }

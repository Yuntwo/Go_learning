package basic

import "strconv"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (iPAddr IPAddr) String() string {
	ss := ""
	for _, addr := range iPAddr {
		// Should use this method instead of string() that will take input as ASCII number
		ss = ss + strconv.Itoa(int(addr)) + "."
	}
	return ss[:len(ss)-1]
}

//func main() {
//	hosts := map[string]IPAddr{
//		"loopback":  {127, 0, 0, 1},
//		"googleDNS": {8, 8, 8, 8},
//	}
//	for name, ip := range hosts {
//		fmt.Printf("%v: %v\n", name, ip)
//	}
//}

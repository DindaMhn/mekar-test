package config

import "mekar/master/tools/getEnv"

//AuthSwitch app
func AuthSwitch() bool {
	isTrue := getEnv.ViperGetEnv("Auth", "YES")
	if isTrue == "YES" {
		return true
	}
	return false
}

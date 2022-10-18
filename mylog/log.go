package mylog

import "log"

/*
LogInfo print logingo
*/
func LogInfo(m string) {
	log.Printf("INFO - %v", m)
}

/*
LogWarn print logingo
*/
func LogWarn(m string) {
	log.Printf("WARN - %v", m)
}

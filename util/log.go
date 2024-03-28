package util

import (
	"fmt"
	"log"
)

type Log struct{}

func (s *Log) LogError(errString string, params interface{}) {
	log.Println(errString, "：", params)
	fmt.Println(errString, "：", params)
}

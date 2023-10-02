package util

import (
	"encoding/json"
	"fmt"
)

type GormErr struct {
	Number  int    `json:"Number"`
	Message string `json:"Message"`
}

func CheckDup(err error) bool {
	byteErr, _ := json.Marshal(err)
	var newError GormErr
	json.Unmarshal((byteErr), &newError)
	switch newError.Number {
	case 1062:
		fmt.Println("Duplicate Key !")
		return true // 重复
	}
	return false // 不重复
}

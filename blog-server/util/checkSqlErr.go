package util

import (
	"encoding/json"
	"fmt"
)

type GormErr struct {
	Number  int    `json:"Number"`
	Message string `json:"Message"`
}

func CheckSqlErr(err error) int {
	byteErr, _ := json.Marshal(err)
	var newError GormErr
	json.Unmarshal((byteErr), &newError)
	switch newError.Number {
	case 1062:
		fmt.Println("Duplicate Key !")
		return 1062 // 重复
	case 1451:
		fmt.Println("Cannot delete or update a parent row: a foreign key constraint fails")
		return 1451 // 存在外键约束
	}
	return 0
}

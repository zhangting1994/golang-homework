package main

import (
	"errors"
	"fmt"
)

func TestSql() (string, error) {
	return "", errors.New("table is not exist")
}

func dto() (string, error) {
	// 模拟数据获取
	data, sqlError := TestSql()
	// 出现 异常
	if sqlError != nil {
		// 封装错误,让上层处理
		sqlError := fmt.Errorf("dto异常,%w", sqlError)
		// 记录日志
		return "", sqlError
	}
	// normal action
	return data, nil
}

func main() {
	s, err2 := dto()
	fmt.Println(s)
	fmt.Println(err2)
	// 出现错误,处理错误,记录日志
	// 解开wrap
	err2 = errors.Unwrap(err2)
	fmt.Println(err2)
	// TODO 根据错误处理
}

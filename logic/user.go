package logic

import (
	"sim/util"
)

type UserImpl struct{}

// CheckUserNameLogic 检查用户名重复，获取 user 表内是否存在用户名
func (s *UserImpl) CheckUserNameLogic(userName string) (bool, error) {
	db := new(util.DB).ConnDB()
	var count int64

	err := db.QueryRow("SELECT count(*) FROM list WHERE userName=?", userName).Scan(&count)

	return count == 0, err
}

// RegisterLogic 创建用户，写入到 user 表
func (s *UserImpl) RegisterLogic(userName, password string) error {
	db := new(util.DB).ConnDB()
	_, err := db.Exec("INSERT INTO user (userName, password) VALUES (?, ?)", userName, password)

	return err
}

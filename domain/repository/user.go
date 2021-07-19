package repository

import (
	"github.com/yino/nlp-controller/domain/po"
)

// UserRepository 聚合工厂
type UserRepository interface {
	Add(user *po.User) error
	Edit(user *po.User) error
	GetUserList(search map[string]interface{}) ([]po.User, error)
	GetUserPage(search map[string]interface{}, page uint, pageSize uint) (datList []po.User, total uint, err error)
	UserInfo(uint64) (*po.User, error)
	FindUserInfo(search map[string]interface{}) (*po.User, error)
	FindUserByToken(token string) (*po.User, error)
	CreateAk(keyPo *po.UserAppKeyPo) error
	GetAkPage(search map[string]interface{}, page, pageSize uint) (datList []po.UserAppKeyPo, total uint, err error)
	FindUserAk(uid uint64, ak string, as string) (po.UserAppKeyPo, error)
}

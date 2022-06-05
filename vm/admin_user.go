package vm

import (
	"ginadmin/model"
	"ginadmin/util"
)


type AdminUser struct {
	Id uint32
	Account string
	Password string
	Type     uint8
	Remark string
	IsLocked uint8
}


func (au *AdminUser) Get() (*model.AdmUser, error) {
	return model.AUGet(au.Id)
}

func (au *AdminUser) GetByAccount() (*model.AdmUser, error) {
	return model.AUGetByAccount(au.Account)
}

func (au *AdminUser) Add() error {
	data := map[string]interface{}{
		"account": au.Account,
		"password": util.Md5(au.Password),
		"type": au.Type,

		"remark": au.Remark,
	}
	return model.AUAdd(data)
}

func (au *AdminUser) Edit() error {
	data := map[string]interface{}{
		"type": au.Type,
		"remark": au.Remark,
	}
	return model.AUEdit(au.Id, data)
}

func (au *AdminUser) Delete() error {
	return model.AUDelete(au.Id)
}

func (au *AdminUser) Lock() error {
	return model.AULock(au.Id, au.IsLocked)
}

func (au *AdminUser) PasswordFun() error {
	return model.AUPassword(au.Id, util.Md5(au.Password))
}

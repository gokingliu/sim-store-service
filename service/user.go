package service

import (
	"net/http"
	"sim/entity"
	"sim/logic"
	"sim/util"
)

type UserImpl struct{}

func (s *UserImpl) Register(res http.ResponseWriter, req *http.Request) {
	// 解析请求参数
	data, err := new(util.Request).JSONRequest(req)
	if err != nil {
		// 解析请求参数出错时，直接返回请求错误
		new(util.Response).JSONResponse(res, entity.ParamErr, false)
		new(util.Log).LogError("Register "+entity.ParamErr.Msg, req.Body)
		return
	}
	// 赋值
	username := data["username"].(string)
	password := data["password"].(string)
	// 正则校验用户名和密码
	userNameResult := new(util.Regexp).CheckUserName(username)
	passwordResult := new(util.Regexp).CheckPassword(password)
	// 正则校验不通过时返回错误
	if !(userNameResult && passwordResult) {
		new(util.Response).JSONResponse(res, entity.ClientUPInvalid, false)
		new(util.Log).LogError("Register "+entity.ClientUPInvalid.Msg, req.Body)
		return
	}
	// 创建用户前判断用户名是否存在
	result, err := new(logic.UserImpl).CheckUserNameLogic(username)
	// 用户名重复时返回错误
	if !result || err != nil {
		new(util.Response).JSONResponse(res, entity.InnerQueryDBError, false)
		new(util.Log).LogError("Register 判断用户名是否存在 "+entity.InnerQueryDBError.Msg, req.Body)
		return
	}
	// 创建用户
	err = new(logic.UserImpl).RegisterLogic(username, password)
	// 创建用户出错时返回错误
	if err != nil {
		new(util.Response).JSONResponse(res, entity.InnerAddDBError, false)
		new(util.Log).LogError("Register "+entity.InnerAddDBError.Msg, req.Body)
		return
	}
	// 正常返回
	new(util.Response).JSONResponse(res, entity.ResOk, true)
}

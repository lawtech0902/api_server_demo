package user

import "go_projects/api_server/model"

/*
__author__ = 'lawtech'
__date__ = '2018/8/10 下午5:34'
*/

type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateResponse struct {
	Username string `json:"username"`
}

type ListRequest struct {
	Username string `json:"username"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
}

type ListResponse struct {
	TotalCount uint64            `json:"total_count"`
	UserList   []*model.UserInfo `json:"user_list"`
}

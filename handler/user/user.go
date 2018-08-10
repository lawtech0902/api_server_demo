package user

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

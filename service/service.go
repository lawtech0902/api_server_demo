package service

import (
	"fmt"
	"go_projects/api_server/model"
	"go_projects/api_server/util"
	"sync"
)

/*
__author__ = 'lawtech'
__date__ = '2018/8/12 下午4:33'
service存放代码量大的业务逻辑处理
*/

func ListUser(username string, limit, offset int) ([]*model.UserInfo, uint64, error) {
	infos := make([]*model.UserInfo, 0)

	users, count, err := model.ListUser(username, offset, limit)
	if err != nil {
		return infos, count, err
	}

	ids := []uint64{}
	for _, user := range users {
		ids = append(ids, user.Id)
	}

	wg := sync.WaitGroup{}
	userList := model.UserList{
		Lock:  new(sync.Mutex),                              // 保证数据一致性
		IdMap: make(map[uint64]*model.UserInfo, len(users)), // 保证数据顺序
	}

	errChan := make(chan error, 1)
	finished := make(chan bool, 1)

	// 提高并行查询的效率
	for _, u := range users {
		wg.Add(1)
		go func(u *model.UserModel) {
			defer wg.Done()

			shortId, err := util.GenShortId()
			if err != nil {
				errChan <- err
			}

			userList.Lock.Lock()
			defer userList.Lock.Unlock()
			userList.IdMap[u.Id] = &model.UserInfo{
				Id:        u.Id,
				Username:  u.Username,
				SayHello:  fmt.Sprintf("Hello %s", shortId),
				Password:  u.Password,
				CreatedAt: u.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt: u.UpdatedAt.Format("2006-01-02 15:04:05"),
			}
		}(u)
	}

	go func() {
		wg.Wait()
		close(finished)
	}()

	select {
	case <-finished:
	case err := <-errChan:
		return infos, count, err
	}

	for _, id := range ids {
		infos = append(infos, userList.IdMap[id])
	}

	return infos, count, err
}

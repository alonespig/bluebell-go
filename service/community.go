package service

import (
	"bluebell/model"
	"bluebell/repository"
)

// GetCommunityList 获取社区列表
func GetCommunityList() ([]*model.Community, error) {
	return repository.GetCommunityList()
}

// GetCommunityDetail 获取社区详情
func GetCommunityDetail(id int) (*model.Community, error) {
	return repository.GetCommunityDetail(id)
}

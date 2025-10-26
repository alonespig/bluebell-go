package repository

import (
	"bluebell/global"
	"bluebell/model"
)

// GetCommunityList 获取社区列表
func GetCommunityList() ([]*model.Community, error) {
	var communities []*model.Community
	if err := global.DB.Model(&model.Community{}).Find(&communities).Error; err != nil {
		return nil, err
	}
	return communities, nil
}

// GetCommunityDetail 获取社区详情	
func GetCommunityDetail(id int) (*model.Community, error) {
	var community *model.Community
	if err := global.DB.Model(&model.Community{}).Where("id = ?", id).First(&community).Error; err != nil {
		return nil, err
	}
	return community, nil
}

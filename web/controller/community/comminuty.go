package community

import (
	"bluebell/froms"
	"bluebell/pkg/code"
	"bluebell/pkg/response"
	"bluebell/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary 获取社区列表
// @Description 获取社区列表
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /api/v1/community [get]
func GetCommunityList(c *gin.Context) {
	communities, err := service.GetCommunityList()
	if err != nil {
		response.JSON(c, http.StatusInternalServerError, code.InternalServerError, err.Error())
	}

	rsp := []froms.CommunityListResponse{}
	for _, community := range communities {
		rsp = append(rsp, froms.CommunityListResponse{
			ID:   community.ID,
			Name: community.CommunityName,
		})
	}

	response.Success(c, rsp)
}

// @Summary 获取社区详情
// @Description 获取社区详情
// @Accept json
// @Produce json
// @Param id path int true "社区ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /api/v1/community/{id} [get]
func GetCommunityDetail(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.JSON(c, http.StatusBadRequest, code.InvalidParams, err.Error())
	}
	community, err := service.GetCommunityDetail(id)
	if err != nil {
		response.JSON(c, http.StatusInternalServerError, code.InternalServerError, err.Error())
	}
	rsp := froms.CommunityDetailResponse{
		ID:           community.ID,
		Name:         community.CommunityName,
		Introduction: community.Introduction,
		CreatedAt:    community.CreatedAt,
	}
	response.Success(c, rsp)
}

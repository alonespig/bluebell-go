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

package froms

import "time"

type CommunityListResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CommunityDetailResponse struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Introduction string    `json:"introduction"`
	MemberCount  int       `json:"member_count"`
	CreatedAt    time.Time `json:"created_at"`
}

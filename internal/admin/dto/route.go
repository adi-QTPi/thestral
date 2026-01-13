package dto

type CreateRouteInput struct {
	Host    string   `json:"host" binding:"required"`
	Targets []string `json:"targets" binding:"required"`
}

type DeleteRouteInput struct {
	Host string `json:"host" binding:"required"`
}

// [TODO] finalise dto struct for api getters' filter

// type RouteFilter struct {
// 	Host     string `json:"host"`
// 	IsActive *bool  `json:"is_active"`
// }

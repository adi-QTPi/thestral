package dto

import "time"

type CreateRouteInput struct {
	Host    string   `json:"host" binding:"required"`
	Targets []string `json:"targets" binding:"required"`
}

type DeleteRouteInput struct {
	Host string `json:"host" binding:"required"`
}

type RouteDisplay struct {
	Host      string    `json:"host"`
	Targets   []string  `json:"targets"`
	IsActive  *bool     `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// [TODO] finalise dto struct for api getters' filter

// type RouteFilter struct {
// 	Host     string `json:"host"`
// 	IsActive *bool  `json:"is_active"`
// }

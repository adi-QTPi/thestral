package types

type AddRouteRequest struct {
	Source      string `json:"source" validate:"required,hostname_rfc1123"`
	Destination string `json:"destination" validate:"required,url"`
}

type HostName struct {
	Host string `json:"host" validate:"required,hostname_rfc1123"`
}

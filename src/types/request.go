package types

type AddRouteRequest struct {
	Source      string `json:"source" valid:"required,hostname_rfc1123"`
	Destination string `json:"destination" valid:"required,url"`
}

type HostName struct {
	Host string `json:"host" valid:"required,hostname_rfc1123"`
}

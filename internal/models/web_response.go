package models

type Info struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type WebResponse struct {
	Info Info `json:"info"`
	Data any  `json:"data"`
}

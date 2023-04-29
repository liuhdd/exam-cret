package controllers

type Resp struct {
    Code int         `json:"code"`
    Msg  string      `json:"message"`
    Data interface{} `json:"data,omitempty"`
}
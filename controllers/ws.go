/*
处理websocket相关业务
*/


package controllers

import (
)

type WS struct {
	*Controller
}

func NewWS() *WS {
	return &WS{
		Controller: &Controller{},
	}
}
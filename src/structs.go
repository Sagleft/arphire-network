package main

import "github.com/gin-gonic/gin"

type solution struct {
	Config userConfig
	Front  front
}

type userConfig struct {
	Front frontConfig `json:"front"`
}

type network struct{}

type front struct {
	Gin    *gin.Engine
	Config frontConfig
}

type frontConfig struct {
	Port    string `json:"port"`
	Version string `json:"version"`
	IsDebug bool   `json:"debug"`
}

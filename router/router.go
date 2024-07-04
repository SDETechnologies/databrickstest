package router

import (
    "main/service"
)

var basePasefiles []string = []string{
    "templates/header.html",
    "templates/base.html",
}

type Router struct {
    service service.IService
}

func NewRouter(service service.IService) *Router {
    return &Router{
        service: service,
    }
}

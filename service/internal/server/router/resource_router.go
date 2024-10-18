package router

import (
	"github.com/go-chi/chi"
	"github.com/tanveerprottoy/jenkins-pipeline/service/internal/server/resource"
	"github.com/tanveerprottoy/jenkins-pipeline/service/pkg/constant"
)

// RegisterResourceRoutes registers resource routes
func RegisterResourceRoutes(router *Router, version string, handler *resource.Handler) {
	router.Mux.Route(
		constant.ApiPattern+version+constant.ResourcesPattern,
		func(r chi.Router) {
			r.Get("/", handler.GetData)
		},
	)
}

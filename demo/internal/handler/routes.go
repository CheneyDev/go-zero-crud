// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"demo/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/users",
				Handler: CreateHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/users/:id",
				Handler: GetHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/users/:id",
				Handler: UpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/users/:id",
				Handler: DeleteHandler(serverCtx),
			},
		},
	)
}

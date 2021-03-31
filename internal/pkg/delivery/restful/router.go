// @Version 0.0.1
// @Title PTCG Trader API v1

package restful

import (
	"pokemon/configuration"
	"pokemon/internal/pkg/iface"
	"pokemon/pkg/claims"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// setRoutes ...
func setRoutes(e *echo.Echo, cfg *configuration.App, h iface.IRestfulHandler) {
	var jwtMiddleware = middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(cfg.JwtSecrets),
		Claims:     &claims.Claims{},
	})

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	rootV1 := e.Group("/apis/v1")

	auth := rootV1.Group("/auth")
	auth.POST("/register", h.Register)
	auth.POST("/login", h.Login)

	card := rootV1.Group("/cards")
	card.GET("", h.ListCards)
	card.GET("/:id", h.GetCard)
	card.POST("", h.CreateCard)
	card.PUT("/:id", h.UpdateCard)

	setSpotOrderRoutes(rootV1, jwtMiddleware, h)
}

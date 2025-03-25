package route

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

	"github.com/dwiprastyoisworo/go-grpc-hexagonal/lib/config"
)

type Route struct {
	db  *sqlx.DB
	ctx *gin.Engine
	cfg *config.App
}

// NewRoute will create an object that represent the Route struct
func NewRoute(db *sqlx.DB, ctx *gin.Engine, cfg *config.App) *Route {
	return &Route{db: db, ctx: ctx, cfg: cfg}
}

// RouteInit will initiate all routes
func (r *Route) RouteInit() {
	r.RouteProduct()
}

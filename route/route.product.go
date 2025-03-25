package route

import (
	"github.com/dwiprastyoisworo/go-grpc-hexagonal/internal/product/adapter/handler"
	"github.com/dwiprastyoisworo/go-grpc-hexagonal/internal/product/adapter/repository"
	"github.com/dwiprastyoisworo/go-grpc-hexagonal/internal/product/service"
)

func (r *Route) RouteProduct() {
	// initiate handler, service and repository
	productRepo := repository.NewProductRepository(r.db)
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	// initiate route product group
	group := r.ctx.Group("/product")
	group.GET("/", productHandler.GetAll)
	group.POST("/", productHandler.Create)
}

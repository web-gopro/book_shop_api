package handlers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/web-gopro/book_shop_api/genproto/book_shop"
)

func (h *Handler) CreateOrder(ctx *gin.Context) {
	var req book_shop.OrderCreateReq

	ctx.BindJSON(&req)

	resp, err := h.service.GetProductSevice().CreateOrder(context.Background(), &req)

	if err != nil {
		h.log.Error(err.Error())
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(201, resp)

}

func (h *Handler) GetOrderById(ctx *gin.Context) {

	var req book_shop.GetByIdReq

	req.Id = ctx.Param("id")

	resp, err := h.service.GetProductSevice().GetOrder(context.Background(), &req)

	if err != nil {
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(201, resp)

}

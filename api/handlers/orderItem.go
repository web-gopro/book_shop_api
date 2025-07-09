package handlers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/web-gopro/book_shop_api/genproto/book_shop"
)

func (h *Handler) CreateOrderItem(ctx *gin.Context) {
	var req book_shop.OrderItemCreateReq

	ctx.BindJSON(&req)

	resp, err := h.service.GetProductSevice().CreateOrdered_Item(context.Background(), &req)

	if err != nil {
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(201, resp)

}

func (h *Handler) GetOrderItemById(ctx *gin.Context) {

	var req book_shop.GetByIdReq

	req.Id = ctx.Param("id")

	resp, err := h.service.GetProductSevice().GetOrdered_Item(context.Background(), &req)

	if err != nil {
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(201, resp)

}

func (h *Handler) GetOrderItemByOrderId(ctx *gin.Context) {

	var req book_shop.GetByIdReq

	req.Id = ctx.Param("id")

	resp, err := h.service.GetProductSevice().GetOrdered_ItemByOrderId(context.Background(), &req)

	if err != nil {
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(201, resp)

}

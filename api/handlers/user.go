package handlers

import (
	"context"
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/web-gopro/book_shop_api/genproto/book_shop"
	"github.com/web-gopro/book_shop_api/mail"
	"github.com/web-gopro/book_shop_api/pkg/helpers"
	"github.com/web-gopro/book_shop_api/token"

	"github.com/saidamir98/udevs_pkg/logger"
)

func (h *Handler) GetUserById(ctx *gin.Context) {

	var req book_shop.GetByIdReq

	req.Id = ctx.Param("id")

	resp, err := h.service.GetUserSevice().GetUser(context.Background(), &req)

	if err != nil {
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(201, resp)

}

func (h *Handler) CheckUser(ctx *gin.Context) {

	var reqBody book_shop.CheckUser

	err := ctx.BindJSON(&reqBody)

	if err != nil {

		ctx.JSON(400, err.Error())
		return
	}

	isExists, err := h.service.GetUserSevice().CheckExists(context.Background(), &book_shop.Common{
		TableName:  "users",
		ColumnName: "email",
		Expvalue:   reqBody.Email,
	})

	if err != nil {

		ctx.JSON(500, err)
		return
	}

	if isExists.IsExists {
		ctx.JSON(201, book_shop.CheckExists{
			IsExists: isExists.IsExists,
			Status:   "sign-in",
		})
		return
	}

	otp := book_shop.OtpData{
		Otp:   mail.GenerateOtp(6),
		Email: reqBody.Email,
	}

	otpdataB, err := json.Marshal(otp)

	if err != nil {

		ctx.JSON(500, err)
		return
	}

	err = h.cache.Set(ctx, reqBody.Email, string(otpdataB), 120)

	err = mail.SendMail([]string{reqBody.Email}, otp.Otp)

	if err != nil {
		h.log.Error("errrr on Send mail", logger.Error(err))
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(201, book_shop.CheckExists{
		IsExists: isExists.IsExists,
		Status:   "registr",
	})

	ctx.JSON(201, "we sent otp")
}

func (h *Handler) SignUp(ctx *gin.Context) {

	var otpData book_shop.OtpData

	var reqBody book_shop.UserCreateReq

	err := ctx.BindJSON(&reqBody)

	if err != nil {
		h.log.Error("errrr on ShouldBindJSON", logger.Error(err))
		ctx.JSON(500, err.Error())
		return
	}

	otpSData, err := h.cache.GetDell(ctx, reqBody.Email)

	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}

	if otpSData == "" {
		ctx.JSON(201, "otp is expired")
		return
	}
	err = json.Unmarshal([]byte(otpSData), &otpData)

	if otpData.Otp != reqBody.Otp {

		ctx.JSON(405, "incorrect otp")
		return
	}

	reqBody.Password, err = helpers.HashPassword(reqBody.Password)

	claim, err := h.service.GetUserSevice().CreateUser(context.Background(), &reqBody)

	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}

	accessToken, err := token.GenerateJWT(*&book_shop.Clamis{UserId: claim.UserId, UserRole: claim.UserRole})

	if err != nil {
		ctx.JSON(201, "registreted")
		return
	}

	ctx.JSON(201, accessToken)

}

func (h *Handler) SigIn(ctx *gin.Context) {

	var reqBody book_shop.UserLogIn

	err := ctx.BindJSON(&reqBody)

	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}

	claim, err := h.service.GetUserSevice().UserLogin(ctx, &reqBody)

	if err != nil {
		if err.Error() == "password is incorrect" {
			ctx.JSON(405, err.Error())
			return
		}
		ctx.JSON(500, err.Error())
		return
	}

	accessToken, err := token.GenerateJWT(*claim)

	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}

	ctx.JSON(201, accessToken)

}

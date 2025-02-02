package handler

import (
	"fmt"
	"net/http"

	"github.com/child6yo/forum-sample"
	"github.com/gin-gonic/gin"
)

// @Summary SignUp
// @Tags Authentication
// @Description create account
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body forum.User true "account info"
// @Success 200 {object} Response
// @Failure 400,404 {object} Response
// @Failure 500 {object} Response
// @Failure default {object} Response
// @Router /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var input forum.User

	if err := c.BindJSON(&input); err != nil {
		err = fmt.Errorf("invalid request data")
		errorResponse(c, "sign up", http.StatusBadRequest, err)
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		err = fmt.Errorf("server error")
		errorResponse(c, "sign up", http.StatusInternalServerError, err)
		return
	}

	successResponse(c, "sign up", map[string]interface{}{
		"user_id": id,
	})
}


// @Summary SignIn
// @Tags Authentication
// @Description login
// @ID login
// @Accept  json
// @Produce  json
// @Param input body forum.SignIn true "credentials"
// @Success 200 {object} Response
// @Failure 400,404 {object} Response
// @Failure 500 {object} Response
// @Failure default {object} Response
// @Router /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var input forum.SignIn

	if err := c.BindJSON(&input); err != nil {
		errorResponse(c, "sign in", http.StatusBadRequest, err)
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		errorResponse(c, "sign in", http.StatusInternalServerError, err)
		return
	}

	successResponse(c, "sign in", map[string]interface{}{
		"token": token,
	})
}
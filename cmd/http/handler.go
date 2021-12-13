package http

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/abdabariassalam/majoo/constants"
	"github.com/abdabariassalam/majoo/internal/service"
	"github.com/gin-gonic/gin"
)

type httpHandler struct {
	svc service.Service
}

func NewHandler(service service.Service) *httpHandler {
	return &httpHandler{
		svc: service,
	}
}

func (h *httpHandler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, "ping pong")
}

func (h *httpHandler) Login(c *gin.Context) {
	req := struct {
		UserName string `json:"user_name"`
		Password string `json:"password"`
	}{}
	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(&req)
	if err != nil {
		if err == io.EOF {
			c.JSON(http.StatusBadGateway, "All input is required")
			return
		}
		c.JSON(http.StatusBadGateway, err.Error())
		return
	}
	//compare the user from the request, with the one we defined:
	if req.UserName == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, "All input is required")
		return
	}
	token, err := h.svc.Login(req.UserName, req.Password)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *httpHandler) MerchantOmzetReport(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.JSON(http.StatusForbidden, constants.ErrTokenRequired)
		return
	}
	user, err := h.svc.VerifyToken(strings.Split(token, "Bearer ")[1])
	if err != nil {
		c.JSON(http.StatusUnauthorized, constants.ErrBadToken)
		return
	}
	req := struct {
		Month   int `json:"month"`
		Year    int `json:"year"`
		Page    int `json:"page"`
		PerPage int `json:"per_page"`
	}{}
	decoder := json.NewDecoder(c.Request.Body)
	err = decoder.Decode(&req)
	if err != nil {
		if err == io.EOF {
			c.JSON(http.StatusBadGateway, "All input is required")
			return
		}
		c.JSON(http.StatusBadGateway, err.Error())
		return
	}
	//compare the user from the request, with the one we defined:
	if req.Month == 0 || req.Year == 0 || req.Page == 0 || req.PerPage == 0 {
		c.JSON(http.StatusBadRequest, "All input is required")
		return
	}
	report, err := h.svc.MerchantOmzetReport(user.ID, req.Month, req.Year, req.Page, req.PerPage)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}
	c.JSON(http.StatusOK, report)
}

func (h *httpHandler) OutletOmzetReport(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.JSON(http.StatusForbidden, constants.ErrTokenRequired)
		return
	}
	user, err := h.svc.VerifyToken(strings.Split(token, "Bearer ")[1])
	if err != nil {
		c.JSON(http.StatusUnauthorized, constants.ErrBadToken)
		return
	}
	req := struct {
		Month   int `json:"month"`
		Year    int `json:"year"`
		Page    int `json:"page"`
		PerPage int `json:"per_page"`
	}{}
	decoder := json.NewDecoder(c.Request.Body)
	err = decoder.Decode(&req)
	if err != nil {
		if err == io.EOF {
			c.JSON(http.StatusBadGateway, "All input is required")
			return
		}
		c.JSON(http.StatusBadGateway, err.Error())
		return
	}
	//compare the user from the request, with the one we defined:
	if req.Month == 0 || req.Year == 0 || req.Page == 0 || req.PerPage == 0 {
		c.JSON(http.StatusBadRequest, "All input is required")
		return
	}
	report, err := h.svc.OutletOmzetReport(user.ID, req.Month, req.Year, req.Page, req.PerPage)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}
	c.JSON(http.StatusOK, report)
}

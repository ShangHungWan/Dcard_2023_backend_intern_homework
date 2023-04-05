package controllers

import (
	"key-value-system/enums"
	"key-value-system/helper"
	"key-value-system/requests"
	"key-value-system/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StoreHead(c *gin.Context) {
	var json requests.CreateHeadRequest

	err := c.ShouldBind(&json)
	if err != nil {
		helper.APIResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = services.StoreHead(json)
	if err != nil {
		helper.APIResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.Status(http.StatusCreated)
}

func StoreNode(c *gin.Context) {
	var json requests.CreateNodeRequest
	if err := c.ShouldBind(&json); err != nil {
		helper.APIResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := services.StoreNode(json, c)
	if err != nil {
		helper.APIResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.Status(http.StatusCreated)
}

func ShowHead(c *gin.Context) {
	key := c.Param("key")

	head, err := services.GetHead(key)
	if err != nil {
		helper.APIResponse(c, http.StatusNotFound, enums.NOT_FOUND)
		return
	}

	helper.APIResponse(c, http.StatusOK, head)
}

func ShowNode(c *gin.Context) {
	key := c.Param("key")

	node, err := services.GetNode(key)
	if err != nil {
		helper.APIResponse(c, http.StatusNotFound, enums.NOT_FOUND)
		return
	}
	helper.APIResponse(c, http.StatusOK, node)
}

func RemoveHead(c *gin.Context) {
	key := c.Param("key")

	err := services.RemoveHead(key)
	if err != nil {
		helper.APIResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}

func RemoveNode(c *gin.Context) {
	key := c.Param("key")

	err := services.RemoveNode(key)
	if err != nil {
		helper.APIResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}

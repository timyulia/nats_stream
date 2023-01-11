package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) get(c *gin.Context) {
	id := (c.PostForm("uid"))

	obj, err := h.services.Read(id)

	if err != nil {
		if err.Error() == "cache miss: map" {
			c.HTML(http.StatusOK, "not_found.html", gin.H{"id": id})
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	order, err := json.MarshalIndent(obj, "", "    ")
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.HTML(http.StatusOK, "found.html", gin.H{"id": id, "order": string(order)})
}

func (h *Handler) mainPage(c *gin.Context) {
	c.HTML(http.StatusOK, "main.html", gin.H{})
}

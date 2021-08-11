package handler

import (
	"github.com/Hargeon/todo"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type dataAllListResponse struct {
	Data []*todo.TodoList `json:"data"`
}

func (h *Handler) createList(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil  {
		return
	}
	var input todo.TodoList
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	listId, err := h.service.TodoList.Create(id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"list_id": listId,
	})
}

func (h *Handler) getAllLists(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil  {
		return
	}

	lists, err := h.service.TodoList.GetLists(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dataAllListResponse{Data: lists})
}

func (h *Handler) getListById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	listId := c.Param("id")
	listIdInt , err := strconv.Atoi(listId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	list, err := h.service.TodoList.GetList(listIdInt, userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, list)
}

func (h *Handler) updateList(c *gin.Context) {

}

func (h *Handler) deleteList(c *gin.Context) {

}

package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"itx-wabizz/models"
	"itx-wabizz/repositories"
)

func HandleChatlist(c *gin.Context){
	var chatlist *models.ChatList
	chatlist, err := repositories.ChatlistRepo.GetChatList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to retrieve chatlist information"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ChatList": chatlist})	
}
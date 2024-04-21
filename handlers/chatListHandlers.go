package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"itx-wabizz/models"
	"itx-wabizz/repositories"
)

func HandleChatlist(c *gin.Context) {
	var chatlists []models.ChatList
	chatlists, err := repositories.ChatlistRepo.GetChatList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to retrieve chatlist information"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ChatList": chatlists})
}

func HandleChatlistSearchByContact(c *gin.Context) {
	var chatlists []models.ChatList
	queryParam := c.Query("keyword")
	if queryParam == "" {
        // If query parameter is not present
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Query parameter 'keyword' is missing",
        })
		return
    }
	chatlists, err := repositories.ChatlistRepo.SearchChatListByContact(queryParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to retrieve chatlist information"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ChatList": chatlists})
}

func HandleChatlistSearchByMessage(c *gin.Context) {
	var chatlists []models.MessageSearchResult
	queryParam := c.Query("keyword")
	if queryParam == "" {
        // If query parameter is not present
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Query parameter 'keyword' is missing",
        })
		return
    }
	chatlists, err := repositories.ChatlistRepo.SearchChatListByMessage(queryParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to retrieve chatlist information"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ChatList": chatlists})
}

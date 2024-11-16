package handlers

import (
    "net/http"
    "task-golang-db/models"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

// HomepageHandler is the handler for the /account/list route
func Homepage(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var accounts []models.Account
        if err := db.Find(&accounts).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Failed to fetch accounts data",
            })
            return
        }
        c.JSON(http.StatusOK, gin.H{
            "accounts": accounts,
        })
    }
}

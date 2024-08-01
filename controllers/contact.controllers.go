package controllers

import (
    "context"
    "database/sql"
    "net/http"
    "strconv"
    "time"

    db "github.com/ozz-life/geoff89__golang-crudsqlc-rest/db/sqlc"
    "github.com/ozz-life/geoff89__golang-crudsqlc-rest/schemas"

    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
)

type ContactController struct {
    db  *db.Queries
    ctx context.Context
}

func NewContactController(db *db.Queries, ctx context.Context) *ContactController {
    return &ContactController{db, ctx}
}

// Create contact  handler
func (cc *ContactController) CreateContact(ctx *gin.Context) {
    var payload *schemas.CreateContact

    if err := ctx.ShouldBindJSON(&payload); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"status": "Failed payload", "error": err.Error()})
        return
    }

    now := time.Now()
    args := &db.CreateContactParams{
        FirstName:   payload.FirstName,
        LastName:    payload.LastName,
        PhoneNumber: payload.PhoneNumber,
        Street:      payload.Street,
        CreatedAt:   now,
        UpdatedAt:   now,
    }

    contact, err := cc.db.CreateContact(ctx, *args)

    if err != nil {
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "Failed retrieving contact", "error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"status": "successfully created contact", "contact": contact})
}
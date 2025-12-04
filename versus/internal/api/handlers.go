package api

import (
	"context"
	"net/http"
	"time"

	"github.com/chessclub2205-dev/versus-service/internal/match"
	"github.com/chessclub2205-dev/versus-service/internal/payments"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

type serviceDeps struct {
	pm *payments.Manager
	mm *match.Matchmaker
}

func RegisterRoutes(r *gin.Engine, pm *payments.Manager, mm *match.Matchmaker) {
	s := &serviceDeps{pm: pm, mm: mm}
	v1 := r.Group("/v1")
	{
		v1.POST("/join", s.joinHandler)
		v1.POST("/cancel", s.cancelHandler)
		v1.POST("/result", s.resultHandler)
	}
	r.GET("/health", func(c *gin.Context) { c.JSON(200, gin.H{"ok": true}) })
}

type joinReq struct {
	UserID      string `json:"user_id" binding:"required"`
	StakeSlices int    `json:"stake_slices" binding:"required"`
	Rating      int    `json:"rating"`
}

func (s *serviceDeps) joinHandler(c *gin.Context) {
	var req joinReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	uid, err := uuid.FromString(req.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user_id"})
		return
	}
	r := match.Request{
		RequestID: uuid.Must(uuid.NewV4()).String(),
		UserID:    uid,
		Rating:    req.Rating,
		Stake:     req.StakeSlices,
		Timestamp: time.Now().Unix(),
	}
	ctx := context.Background()
	if err := s.mm.Enqueue(ctx, r); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"request_id": r.RequestID, "status": "queued"})
}

type resultReq struct {
	MatchID  string `json:"match_id" binding:"required"`
	WinnerID string `json:"winner_id" binding:"required"`
	Token    string `json:"token" binding:"required"` // simple auth
}

func (s *serviceDeps) resultHandler(c *gin.Context) {
	var req resultReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// simple server auth
	if req.Token != "server-secret-token" { // use env var in prod
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	mid, err := uuid.FromString(req.MatchID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid match_id"})
		return
	}
	wid, err := uuid.FromString(req.WinnerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid winner_id"})
		return
	}
	ctx := context.Background()
	if err := s.pm.SettleMatch(ctx, mid, wid); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": "settled"})
}

func (s *serviceDeps) cancelHandler(c *gin.Context) {
	// cancellation logic left for integration: remove from redis queue and refund if locked
	c.JSON(200, gin.H{"status": "ok"})
}

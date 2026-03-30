package controllers

import (
	"errors"
	"net/http"
	"voting-system/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type VoteRequest struct {
	VoterID     uint `json:"voter_id" binding:"required"`
	CandidateID uint `json:"candidate_id" binding:"required"`
}

type CreateVoterRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

func (h *Handler) GetVoters(c *gin.Context) {
	var voters []models.Voter
	if err := h.DB.Preload("Candidate").Find(&voters).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data voter"})
		return
	}

	c.JSON(http.StatusOK, voters)
}

func (h *Handler) CreateVoter(c *gin.Context) {
	var req CreateVoterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Payload tidak valid"})
		return
	}

	voter := models.Voter{
		Name:  req.Name,
		Email: req.Email,
	}

	if err := h.DB.Create(&voter).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Gagal membuat voter (email mungkin sudah dipakai)"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Voter berhasil ditambahkan", "voter": voter})
}

// Fungsi untuk pemilih melakukan voting
func (h *Handler) Vote(c *gin.Context) {
	var voteData VoteRequest
	if err := c.ShouldBindJSON(&voteData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Payload tidak valid"})
		return
	}

	err := h.DB.Transaction(func(tx *gorm.DB) error {
		var voter models.Voter
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&voter, voteData.VoterID).Error; err != nil {
			return err
		}

		if voter.Voted {
			return errors.New("Telah dipilih")
		}

		var candidate models.Candidate
		if err := tx.First(&candidate, voteData.CandidateID).Error; err != nil {
			return err
		}

		voter.Voted = true
		voter.CandidateID = voteData.CandidateID
		if err := tx.Save(&voter).Error; err != nil {
			return err
		}

		if err := tx.Model(&models.Candidate{}).
			Where("id = ?", voteData.CandidateID).
			UpdateColumn("votes", gorm.Expr("votes + ?", 1)).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		switch {
		case err.Error() == "already_voted":
			c.JSON(http.StatusConflict, gin.H{"error": "Voter sudah melakukan voting"})
		case errors.Is(err, gorm.ErrRecordNotFound):
			c.JSON(http.StatusNotFound, gin.H{"error": "Voter atau kandidat tidak ditemukan"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memproses voting"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Voting berhasil"})
}

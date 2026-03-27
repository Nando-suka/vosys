package controllers

import (
	"net/http"
	"voting-system/models"

	"github.com/gin-gonic/gin"
)

// Menampilkan semua kandidat
func (h *Handler) GetCandidates(c *gin.Context) {
	var candidates []models.Candidate
	h.DB.Find(&candidates)
	c.JSON(http.StatusOK, candidates)
}

// Menampilkan Kandidat Berdasarkan ID
func (h *Handler) GetCandidateByID(c *gin.Context) {
	id := c.Param("id")
	var candidate models.Candidate

	if err := h.DB.First(&candidate, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kandidat tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, candidate)
}

// Menampilkan hasil voting
func (h *Handler) GetCandidateRanking(c *gin.Context) {
	var candidates []models.Candidate
	h.DB.Order("votes desc").Find(&candidates)
	c.JSON(http.StatusOK, candidates)
}

// Menambahkan kandidat baru
func (h *Handler) CreateCandidate(c *gin.Context) {
	var candidate models.Candidate

	// Bind JSON dari request ke struct
	if err := c.ShouldBindJSON(&candidate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Simpan ke database
	if err := h.DB.Create(&candidate).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menambahkan kandidat"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Kandidat berhasil ditambahkan", "candidate": candidate})

}

func (h *Handler) DeleteCandidate(c *gin.Context) {
	// Ambil ID kandidat dari URL
	id := c.Param("id")

	// Cari kandidat berdasarkan ID
	var candidate models.Candidate
	if err := h.DB.First(&candidate, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kandidat tidak ditemukan"})
		return
	}

	// Hapus kandidat dari database
	if err := h.DB.Delete(&candidate).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus kandidat"})
		return
	}

	// Kirim respons sukses
	c.JSON(http.StatusOK, gin.H{"message": "Kandidat berhasil dihapus"})
}

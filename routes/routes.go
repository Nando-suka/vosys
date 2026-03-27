package routes

import (
	"voting-system/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	h := controllers.NewHandler(db)
	r.StaticFile("/", "./web/index.html")

	// Endpoint untuk kandidat
	r.GET("/candidates", h.GetCandidates)               // Lihat semua kandidat
	r.GET("/candidates/:id", h.GetCandidateByID)        // Lihat kandidat yang ada berdasarkan ID
	r.GET("/candidates/ranking", h.GetCandidateRanking) // Update Ranking
	r.POST("/candidates", h.CreateCandidate)            // Tambah kandidat baru
	r.DELETE("/candidates/:id", h.DeleteCandidate)      // Hapus kandidat
	r.POST("/vote", h.Vote)                             // Voting utama via voter
	r.GET("/voters", h.GetVoters)                       // Lihat semua voter
	r.POST("/voters", h.CreateVoter)                    // Tambah voter baru

	return r
}

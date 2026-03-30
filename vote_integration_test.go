package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"voting-system/models"
	"voting-system/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setupTestRouter(t *testing.T) (*gorm.DB, *gin.Engine) {
	t.Helper()

	dsn := os.Getenv("root:@tcp(127.0.0.1:3306)/dbvoting")
	if dsn == "" {
		t.Skip("TEST_DB_DSN belum di-set, test integrasi MySQL dilewati")
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("gagal koneksi test database: %v", err)
	}

	if err := db.AutoMigrate(&models.Candidate{}, &models.Voter{}); err != nil {
		t.Fatalf("gagal migrasi test database: %v", err)
	}

	db.Exec("DELETE FROM voters")
	db.Exec("DELETE FROM candidates")

	gin.SetMode(gin.TestMode)
	return db, routes.SetupRouter(db)
}

func TestVoteSuccess(t *testing.T) {
	db, router := setupTestRouter(t)

	candidate := models.Candidate{Name: "Alice", Country: "ID"}
	voter := models.Voter{Name: "Budi", Email: "budi.vote@test.local"}

	if err := db.Create(&candidate).Error; err != nil {
		t.Fatalf("gagal membuat candidate: %v", err)
	}
	if err := db.Create(&voter).Error; err != nil {
		t.Fatalf("gagal membuat voter: %v", err)
	}

	payload := map[string]uint{
		"voter_id":     voter.ID,
		"candidate_id": candidate.ID,
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPost, "/vote", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d, body=%s", w.Code, w.Body.String())
	}
}

func TestVoteTwiceReturnsConflict(t *testing.T) {
	db, router := setupTestRouter(t)

	candidate := models.Candidate{Name: "Bob", Country: "ID"}
	voter := models.Voter{Name: "Cici", Email: "cici.vote@test.local", Voted: true}

	if err := db.Create(&candidate).Error; err != nil {
		t.Fatalf("gagal membuat candidate: %v", err)
	}
	if err := db.Create(&voter).Error; err != nil {
		t.Fatalf("gagal membuat voter: %v", err)
	}

	payload := map[string]uint{
		"voter_id":     voter.ID,
		"candidate_id": candidate.ID,
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPost, "/vote", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusConflict {
		t.Fatalf("expected 409, got %d, body=%s", w.Code, w.Body.String())
	}
}

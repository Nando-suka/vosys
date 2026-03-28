# Voting System (Golang)

Simple learning project to build a voting system using:
- Go
- Gin (HTTP framework)
- GORM
- MySQL

This project is designed for first-time backend learning: CRUD data, voting flow, transaction safety, and integration testing.

## Features

- Create and list candidates
- Delete candidate
- Create and list voters
- Vote endpoint with one-voter-one-vote rule
- Ranking endpoint sorted by total votes
- Simple web demo page for manual testing (`/`)
- MySQL integration tests

## Project Structure

```text
voting-system/
  controllers/
  database/
  models/
  routes/
  web/
  main.go
  vote_integration_test.go
```

## API Endpoints

- `GET /` -> demo page
- `GET /candidates`
- `GET /candidates/:id`
- `GET /candidates/ranking`
- `POST /candidates`
- `DELETE /candidates/:id`
- `GET /voters`
- `POST /voters`
- `POST /vote`

### Example Payloads

Create candidate:

```json
{
  "name": "Alice",
  "country": "Indonesia"
}
```

Create voter:

```json
{
  "name": "Budi",
  "email": "budi@example.com"
}
```

Vote:

```json
{
  "voter_id": 1,
  "candidate_id": 1
}
```

## Prerequisites

- Go 1.23+
- MySQL running locally
- Database created (example: `dbvoting`)

## Run Locally

1. Set environment variable:

Windows PowerShell:

```powershell
$env:DB_DSN="root:@tcp(127.0.0.1:3306)/dbvoting?charset=utf8mb4&parseTime=True&loc=Local"
```

2. Run app:

```powershell
go run .
```

3. Open:

`http://localhost:8080`

## Run Tests (MySQL Integration)

Use a separate database for tests, for example `dbvoting_test`.

Windows PowerShell:

```powershell
$env:TEST_DB_DSN="root:@tcp(127.0.0.1:3306)/dbvoting_test?charset=utf8mb4&parseTime=True&loc=Local"
go test ./... -v
```

If `TEST_DB_DSN` is not set, integration tests will be skipped.

## Push This Project to GitHub

From inside `voting-system` directory:

```powershell
git init
git add .
git commit -m "Initial voting system project"
git branch -M main
git remote add origin https://github.com/<your-username>/voting-system.git
git push -u origin main
```

If this folder is already inside another git repository, create a new GitHub repository from that current root and push normally from your current branch.

## Learning Notes

- `POST /vote` uses a database transaction and row locking to reduce race conditions.
- A voter can vote only once (`voted = true` after success).
- Candidate votes are incremented atomically in SQL.

---

Made for learning backend fundamentals in Golang.
TBH I'm still vibecoding all of these code. However, I will plan to learn little by little to improve my golang skills and I'm so happy if anyone could give me any feedback, harsh or kind accepted :)

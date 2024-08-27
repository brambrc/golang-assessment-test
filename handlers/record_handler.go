package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"math/rand"
	_ "mezink-goland-assessment/models"
	"net/http"
	"time"
)

type RecordHandler struct {
	DB *sqlx.DB
}

type RequestPayload struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	MinCount  int    `json:"minCount"`
	MaxCount  int    `json:"maxCount"`
}

type RecordItem struct {
	ID         int64     `json:"id"`
	CreatedAt  time.Time `json:"createdAt"`
	TotalMarks int       `json:"totalMarks"`
}

type newRecord struct {
	ID        int64         `db:"id" json:"id"`
	Name      string        `db:"name" json:"name"`
	Marks     pq.Int64Array `db:"marks" json:"marks"`
	CreatedAt time.Time     `db:"createdat" json:"createdAt"`
}

type ResponsePayload struct {
	Code    int          `json:"code"`
	Msg     string       `json:"msg"`
	Records []RecordItem `json:"records"`
}

type ResponseFetchTable struct {
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Records []newRecord `json:"records"`
}

// generating random array elements from min and max marks
func generateRandomMarks(minSum, maxSum int, numMarks int) []int {
	if minSum > maxSum || numMarks <= 0 {
		return nil
	}

	marks := make([]int, numMarks)
	sum := 0

	for i := 0; i < numMarks-1; i++ {
		remainingSum := maxSum - sum - (numMarks - i - 1)
		if remainingSum <= 1 {
			remainingSum = 1
		}

		mark := rand.Intn(remainingSum) + 1
		marks[i] = mark
		sum += mark
	}

	// calculate the last mark to ensure the total sum is within the range
	lastMark := minSum - sum
	if lastMark < 1 {
		lastMark = 1
	}
	if lastMark > maxSum-sum {
		lastMark = maxSum - sum
	}

	marks[numMarks-1] = lastMark

	return marks
}

func (h *RecordHandler) InsertAndFetch(w http.ResponseWriter, r *http.Request) {
	var req RequestPayload
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	rand.Seed(time.Now().UnixNano())

	// calling function random marks that sum between minCount and maxCount
	sampleMarks := generateRandomMarks(req.MinCount, req.MaxCount, 3)

	// Insert a sample record into the database
	sampleName := "John Doe"
	createdAt := time.Now()

	insertQuery := `
        INSERT INTO records (name, marks, createdAt)
        VALUES ($1, $2, $3)
        RETURNING id
    `
	var id int64
	err := h.DB.QueryRowx(insertQuery, sampleName, pq.Array(sampleMarks), createdAt).Scan(&id)
	if err != nil {
		http.Error(w, "Failed to insert record", http.StatusInternalServerError)
		fmt.Printf("Insert Error: %v\n", err)
		return
	}

	// convert startDate and endDate to time.Time
	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		http.Error(w, "Invalid startDate format", http.StatusBadRequest)
		return
	}

	endDate, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		http.Error(w, "Invalid endDate format", http.StatusBadRequest)
		return
	}

	// Adjusted query using LATERAL join to handle unsetting array elements
	fetchQuery := `
        SELECT id, createdAt, SUM(m) as totalMarks
        FROM records, LATERAL unnest(marks) as m
        WHERE createdAt BETWEEN $1 AND $2
        GROUP BY id, createdAt
        HAVING SUM(m) BETWEEN $3 AND $4
    `

	rows, err := h.DB.Queryx(fetchQuery, startDate, endDate, req.MinCount, req.MaxCount)
	if err != nil {
		http.Error(w, "Database query failed", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var records []RecordItem
	for rows.Next() {
		var record RecordItem
		if err := rows.StructScan(&record); err != nil {
			http.Error(w, "Failed to scan record", http.StatusInternalServerError)
			return
		}
		records = append(records, record)
	}
	var response ResponsePayload
	if len(records) < 1 {
		response = ResponsePayload{
			Code:    1,
			Msg:     "Success, No records found for the given date range",
			Records: records,
		}
	} else {
		response = ResponsePayload{
			Code:    0,
			Msg:     "Success",
			Records: records,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

func (h *RecordHandler) FetchTable(w http.ResponseWriter, r *http.Request) {
	query := `
        SELECT id, name, marks, createdAt
        FROM records
    `

	rows, err := h.DB.Queryx(query)
	if err != nil {
		http.Error(w, "Database query failed", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var records []newRecord
	for rows.Next() {
		var record newRecord
		if err := rows.StructScan(&record); err != nil {
			http.Error(w, "Failed to scan record", http.StatusInternalServerError)
			return
		}
		records = append(records, record)
	}

	response := ResponseFetchTable{
		Code:    0,
		Msg:     "Success",
		Records: records,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

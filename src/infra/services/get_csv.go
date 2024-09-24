package services

import (
	"time"
)

type GetCsvResponse struct {
	OK    bool   `json:"ok"`
	ID    int64  `json:"id"`
	ReqID string `json:"req_id"`
}

func Get_csv_service() (GetCsvResponse, error) {
	return GetCsvResponse{OK: true, ID: time.Now().UnixNano(), ReqID: "get csv"}, nil
}

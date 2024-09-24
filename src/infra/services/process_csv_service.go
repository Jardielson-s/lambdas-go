package services

import (
	"time"
)

type CreateResponse struct {
	OK    bool   `json:"ok"`
	ID    int64  `json:"id"`
	ReqID string `json:"req_id"`
}

func Process_csv_service() (CreateResponse, error) {
	return CreateResponse{OK: true, ID: time.Now().UnixNano(), ReqID: "process csv"}, nil
}

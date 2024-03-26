package server

type BackFillOneInchRequest struct {
	QueryID int64  `json:"query_id" binding:"required"`
	Version string `json:"version"`
}

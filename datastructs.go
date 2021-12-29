package main

type StatusType string

const (
	StatusOk  StatusType = "ok"
	StatusErr StatusType = "err"
)

type Status struct {
	StatusType `json:"status"`
	Error      string `json:"error"`
}

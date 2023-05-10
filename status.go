package main

type status struct {
	ID            int    `json:"id"`
	ApplicationId int    `json:"applicationId"`
	Type          string `json:"type"`
	Note          string `json:"note"`
	Timestamp     string `json:"timestamp"`
}

package model

import "time"

type Alert struct {
	Labels      map[string]string `json:"labels"`
	Annotations map[string]string `json:"annotations"`
	StartsAt    time.Time         `json:"starts_at"`
	EndsAt      time.Time         `json:"ends_at"`
}
type Notification struct {
	Alertname   string `json:"alertname"`
	Instance    string `json:"instance"`
	Job         string `json:"job"`
	ServcieName string `json:"servcie_name"`
	Severity    string `json:"severity"`
	Description string `json:"description"`
	Summary     string `json:"summary"`
}

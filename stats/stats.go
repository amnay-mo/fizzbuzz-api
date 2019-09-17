package stats

// Stats represents an endpoint's hit count
type Stats struct {
	Parameters string `json:"parameters"`
	Count      int    `json:"count"`
}

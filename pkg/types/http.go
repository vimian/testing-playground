package types

// HTTPResponse represents a standard structure for HTTP responses.
type HTTPResponse struct {
	Error  string      `json:"error,omitempty"`
	Status string      `json:"status,omitempty"`
	Result interface{} `json:"result,omitempty"`
}

// Float64s is a slice of float64 numbers.
type Float64s []float64

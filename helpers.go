package entrasync

import abstractions "github.com/microsoft/kiota-abstractions-go"

func advancedQueryHeaders() *abstractions.RequestHeaders {
	h := abstractions.NewRequestHeaders()
	h.Add("ConsistencyLevel", "eventual")
	return h
}

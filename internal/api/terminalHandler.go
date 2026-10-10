package api

import (
	"Terminal_Transaction/internal/report"
	"encoding/json"
	"net/http"
)

func (h *Handler) GetTerminals(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	terminalReport, err := report.TerminalReport(ctx, h.db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(terminalReport)
	if err != nil {
		return
	}

}

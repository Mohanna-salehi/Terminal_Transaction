package api

import (
	"Terminal_Transaction/internal/repository"
	"encoding/json"
	"net/http"
)

func (h *Handler) GetDailyTransaction(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	dailyReport, err :=
		repository.FindDailyTerminalReport(ctx, h.db)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(dailyReport); err != nil {
		return
	}
}

func (h *Handler) GetLastTransaction(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	lastTransaction, err := repository.FindLastTransactionReport(ctx, h.db)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(lastTransaction)
	if err != nil {
		return
	}
}

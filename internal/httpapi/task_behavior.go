package httpapi

import (
	"chargeguard/internal/charging"
	"net/http"
)

func TaskHTTPHandler(w http.ResponseWriter, r *http.Request) {
	err := charging.SendReminder()
	if charging.IsRetryableReminder(err) {
		http.Error(w, "retry", http.StatusServiceUnavailable)
		return
	}
	if err != nil {
		http.Error(w, "permanent failure", http.StatusUnprocessableEntity)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

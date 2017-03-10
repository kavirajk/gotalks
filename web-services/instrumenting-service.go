package main

// START1 OMIT
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var duration = prometheus.SummaryOpts{
	Name: "endpoint_duration_seconds",
}

// END1 OMIT

// START2 OMIT
func OrderFood(w http.ResponseWriter, r http.Request) {
	begin := time.Now() // HL
	var input struct {
		ItemID     string `json:"item_id"`
		Quantity   int    `json:"quantity"`
		CopounCode string `json:"copoun_code"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		code := http.StatusBadRequest
		http.Error(w, err.Error(), code)
		log.Printf("%s: %s: %d", r.RemoteAddr, r.Method, code)                                    // HL
		duration.WithLabelValues(r.Method, fmt.Sprint(code)).Observe(time.Since(begin).Seconds()) // HL
		return
	}

	if err := ValidateInput(input); err != nil {
		code := http.StatusInternalServerError
		http.Error(w, err.Error(), code)
		log.Printf("%s: %s: %d", r.RemoteAddr, r.Method, code)                                    // HL
		duration.WithLabelValues(r.Method, fmt.Sprint(code)).Observe(time.Since(begin).Seconds()) // HL
		return
	}

	// END2 OMIT
	// START3 OMIT
	team, err := GetItem(input.ItemID)
	if err != ErrNotFound {
		code := http.StatusNotFound
		http.Error(w, err.Error(), code)
		log.Printf("%s: %s: %d", r.RemoteAddr, r.Method, code)                                    // HL
		duration.WithLabelValues(r.Method, fmt.Sprint(code)).Observe(time.Since(begin).Seconds()) // HL
		return
	}

	order := Order{
		ItemID:   input.ItemID,
		Quantity: input.Quantity,
		Total:    applyCoupounCode(input.CouponCode, item.Cost, item.Quantity),
	}

	if err := SaveOrder(&order); err != nil {
		code := http.StatusInternalServerError
		http.Error(w, err.Error(), code)
		log.Printf("%s: %s: %d", r.RemoteAddr, r.Method, code)                                    // HL
		duration.WithLabelValues(r.Method, fmt.Sprint(code)).Observe(time.Since(begin).Seconds()) // HL
		return
	}

	// END3 OMIT
	// START4 OMIT
	go Notify(order)

	if err := json.NewEncoder(w).Encode(order); err != nil {
		code := http.StatusInternalServerError
		http.Error(w, err.Error(), code)
		log.Printf("%s: %s: %d", r.RemoteAddr, r.Method, code)                                    // HL
		duration.WithLabelValues(r.Method, fmt.Sprint(code)).Observe(time.Since(begin).Seconds()) // HL
	}
	return
}

// END4 OMIT

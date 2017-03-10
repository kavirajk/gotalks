package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// START1 OMIT
func OrderFood(w http.ResponseWriter, r http.Request) {
	var input struct {
		ItemID     string `json:"item_id"`
		Quantity   int    `json:"quantity"`
		CopounCode string `json:"copoun_code"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		code := http.StatusBadRequest
		http.Error(w, err.Error(), code)
		log.Printf("%s: %s: %d", r.RemoteAddr, r.Method, code) // HL
		return
	}

	if err := ValidateInput(input); err != nil {
		code := http.StatusInternalServerError
		http.Error(w, err.Error(), code)
		log.Printf("%s: %s: %d", r.RemoteAddr, r.Method, code) // HL
		return
	}
	// END1 OMIT
	// START2 OMIT
	team, err := GetItem(input.ItemID)
	if err != ErrNotFound {
		code := http.StatusNotFound
		http.Error(w, err.Error(), code)
		log.Printf("%s: %s: %d", r.RemoteAddr, r.Method, code) // HL
		return
	}

	order := Order{
		ItemID:   input.ItemID,
		Quantity: input.Quantity,
		Total:    applyCoupounCode(input.CouponCode, item.Cost, item.Quantity),
	}
	// END2 OMIT
	// START3 OMIT
	if err := SaveOrder(&order); err != nil {
		code := http.StatusInternalServerError
		http.Error(w, err.Error(), code)
		log.Printf("%s: %s: %d", r.RemoteAddr, r.Method, code) // HL
		return
	}

	go Notify(order)

	if err := json.NewEncoder(w).Encode(order); err != nil {
		code := http.StatusInternalServerError
		http.Error(w, err.Error(), code)
		log.Printf("%s: %s: %d", r.RemoteAddr, r.Method, code) // HL
	}
	return
}

// END3 OMIT

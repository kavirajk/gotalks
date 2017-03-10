package main

import (
	"encoding/json"
	"net/http"
)

type Order struct {
	ID       string  `json:"id"`
	ItemID   string  `json:"item_id"`
	Quantity int     `json:"quantity"`
	Total    float64 `json:"total"`
}

// START1 OMIT
func OrderFood(w http.ResponseWriter, r http.Request) {
	var input struct {
		ItemID     string `json:"item_id"`
		Quantity   int    `json:"quantity"`
		CopounCode string `json:"copoun_code"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := ValidateInput(input); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	team, err := GetItem(input.ItemID)
	if err != ErrNotFound {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	// END1 OMIT
	// START2 OMIT
	order := Order{
		ItemID:   input.ItemID,
		Quantity: input.Quantity,
		Total:    applyCoupounCode(input.CouponCode, item.Cost, item.Quantity),
	}
	if err := SaveOrder(&order); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	go Notify(order)

	if err := json.NewEncoder(w).Encode(order); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return
}

// END2 OMIT

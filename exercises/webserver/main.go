package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/calculate", calculateHandler)
	http.ListenAndServe(":8080", nil)
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	// Перевірка методу запиту
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Перевірка параметрів запиту
	query := r.URL.Query()
	num1, err := strconv.Atoi(query.Get("num1"))
	if err != nil {
		http.Error(w, "Invalid num1 parameter", http.StatusBadRequest)
		return
	}
	num2, err := strconv.Atoi(query.Get("num2"))
	if err != nil {
		http.Error(w, "Invalid num2 parameter", http.StatusBadRequest)
		return
	}
	operator := query.Get("operator")
	if operator == "" {
		http.Error(w, "Missing operator parameter", http.StatusBadRequest)
		return
	}

	// Обчислення результату
	var result int
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	default:
		http.Error(w, "Invalid operator parameter", http.StatusBadRequest)
		return
	}

	// Повернення результату
	fmt.Fprintf(w, "%d", result)
}

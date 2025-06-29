package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Calculator struct {
	Result float64
	Error  string
}

func main() {
	// Serve static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	
	// Handle routes
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/calculate", handleCalculate)
	
	fmt.Println("Server berjalan di http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, Calculator{})
}

func handleCalculate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	expression := r.FormValue("expression")
	result, err := evaluateExpression(expression)
	
	calc := Calculator{}
	if err != nil {
		calc.Error = err.Error()
	} else {
		calc.Result = result
	}

	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, calc)
}

func evaluateExpression(expr string) (float64, error) {
	// Remove spaces and convert to lowercase
	expr = strings.ReplaceAll(expr, " ", "")
	expr = strings.ToLower(expr)
	
	// Simple expression evaluator
	// This is a basic implementation - in production you'd want a more robust parser
	
	// Handle basic operations
	if strings.Contains(expr, "+") {
		parts := strings.Split(expr, "+")
		if len(parts) != 2 {
			return 0, fmt.Errorf("invalid expression")
		}
		a, err := strconv.ParseFloat(parts[0], 64)
		if err != nil {
			return 0, fmt.Errorf("invalid number: %s", parts[0])
		}
		b, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			return 0, fmt.Errorf("invalid number: %s", parts[1])
		}
		return a + b, nil
	}
	
	if strings.Contains(expr, "-") {
		parts := strings.Split(expr, "-")
		if len(parts) != 2 {
			return 0, fmt.Errorf("invalid expression")
		}
		a, err := strconv.ParseFloat(parts[0], 64)
		if err != nil {
			return 0, fmt.Errorf("invalid number: %s", parts[0])
		}
		b, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			return 0, fmt.Errorf("invalid number: %s", parts[1])
		}
		return a - b, nil
	}
	
	if strings.Contains(expr, "*") {
		parts := strings.Split(expr, "*")
		if len(parts) != 2 {
			return 0, fmt.Errorf("invalid expression")
		}
		a, err := strconv.ParseFloat(parts[0], 64)
		if err != nil {
			return 0, fmt.Errorf("invalid number: %s", parts[0])
		}
		b, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			return 0, fmt.Errorf("invalid number: %s", parts[1])
		}
		return a * b, nil
	}
	
	if strings.Contains(expr, "/") {
		parts := strings.Split(expr, "/")
		if len(parts) != 2 {
			return 0, fmt.Errorf("invalid expression")
		}
		a, err := strconv.ParseFloat(parts[0], 64)
		if err != nil {
			return 0, fmt.Errorf("invalid number: %s", parts[0])
		}
		b, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			return 0, fmt.Errorf("invalid number: %s", parts[1])
		}
		if b == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		return a / b, nil
	}
	
	// If no operator found, try to parse as single number
	result, err := strconv.ParseFloat(expr, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid expression")
	}
	
	return result, nil
} 
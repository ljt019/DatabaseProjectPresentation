package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func withCORS(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.WriteHeader(http.StatusOK)
			return
		}
		handler(w, r)
	}
}

func executeSQLFile(filePath string) error {
	// 1. Parse the SQL query from a file
	query, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("Error reading query file: %w", err)
	}

	// Split the query by semicolon
	queries := strings.Split(string(query), ";")

	for _, q := range queries {
		q = strings.TrimSpace(q) // remove leading and trailing spaces
		if q == "" {
			continue // skip empty queries
		}

		// 2. Prepare the SQL statement
		stmt, err := db.Prepare(q)
		if err != nil {
			return fmt.Errorf("Error preparing statement: %w", err)
		}
		defer stmt.Close()

		// 3. Execute the statement
		_, err = stmt.Exec()
		if err != nil {
			return fmt.Errorf("Error executing statement: %w", err)
		}
	}

	return nil
}

func exportTableToJSON(w http.ResponseWriter, tableName string) error {
	// Check if the 'Email' column exists
	var emailExists bool
	err := db.QueryRow(fmt.Sprintf("SELECT COUNT(*) > 0 FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_NAME = '%s' AND COLUMN_NAME = 'Email'", tableName)).Scan(&emailExists)
	if err != nil {
		return fmt.Errorf("Error checking for Email column: %w", err)
	}

	// Construct the query based on the presence of the Email column
	var query string
	if emailExists {
		query = fmt.Sprintf("SELECT * FROM %s", tableName)
	} else {
		query = fmt.Sprintf("SELECT * FROM %s", tableName)
	}

	// Query all rows from the table
	rows, err := db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		return err
	}

	// Prepare a slice of maps to hold the row data
	var results []map[string]interface{}

	// Iterate over the rows
	for rows.Next() {
		// Prepare a slice to hold the column values
		columnValues := make([]interface{}, len(columns))
		columnPointers := make([]interface{}, len(columns))
		for i := range columnValues {
			columnPointers[i] = &columnValues[i]
		}

		// Scan the row values into the column pointers
		err := rows.Scan(columnPointers...)
		if err != nil {
			return err
		}

		// Create a map to hold the column name and value pairs
		rowMap := make(map[string]interface{})
		for i, colName := range columns {
			val := columnValues[i]

			// Convert []byte to string for readability in JSON
			if b, ok := val.([]byte); ok {
				rowMap[colName] = string(b)
			} else {
				rowMap[colName] = val
			}
		}

		results = append(results, rowMap)
	}

	// Check for errors from iterating over rows
	if err = rows.Err(); err != nil {
		return err
	}

	// Set the content type to application/json
	w.Header().Set("Content-Type", "application/json")

	// Encode the results to JSON and write to the response writer
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(results)
	if err != nil {
		return err
	}

	return nil
}

func executeSQLFileToJSON(w http.ResponseWriter, filePath string) error {
	// Read the SQL query from the file
	query, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("Error reading query file: %w", err)
	}

	// Execute the query and fetch results
	rows, err := db.Query(string(query))
	if err != nil {
		return fmt.Errorf("Error executing query: %w", err)
	}
	defer rows.Close()

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		return fmt.Errorf("Error getting columns: %w", err)
	}

	// Prepare a slice of maps to hold the row data
	var results []map[string]interface{}

	// Iterate over the rows
	for rows.Next() {
		// Prepare a slice to hold the column values
		columnValues := make([]interface{}, len(columns))
		columnPointers := make([]interface{}, len(columns))
		for i := range columnValues {
			columnPointers[i] = &columnValues[i]
		}

		// Scan the row values into the column pointers
		err := rows.Scan(columnPointers...)
		if err != nil {
			return fmt.Errorf("Error scanning row: %w", err)
		}

		// Create a map to hold the column name and value pairs
		rowMap := make(map[string]interface{})
		for i, colName := range columns {
			val := columnValues[i]

			// Convert []byte to string for readability in JSON
			if b, ok := val.([]byte); ok {
				rowMap[colName] = string(b)
			} else {
				rowMap[colName] = val
			}
		}

		results = append(results, rowMap)
	}

	// Check for errors from iterating over rows
	if err = rows.Err(); err != nil {
		return fmt.Errorf("Error iterating over rows: %w", err)
	}

	// Set the content type to application/json
	w.Header().Set("Content-Type", "application/json")

	// Encode the results to JSON and write to the response writer
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(results)
	if err != nil {
		return fmt.Errorf("Error encoding results to JSON: %w", err)
	}

	return nil
}

func exportQueryToJSON(w http.ResponseWriter, query string) {
	// Execute the query
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error executing query: %v", err), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting columns: %v", err), http.StatusInternalServerError)
		return
	}

	// Prepare a slice of maps to hold the row data
	var results []map[string]interface{}

	// Iterate over the rows
	for rows.Next() {
		// Prepare a slice to hold the column values
		columnValues := make([]interface{}, len(columns))
		columnPointers := make([]interface{}, len(columns))
		for i := range columnValues {
			columnPointers[i] = &columnValues[i]
		}

		// Scan the row values into the column pointers
		err := rows.Scan(columnPointers...)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error scanning row: %v", err), http.StatusInternalServerError)
			return
		}

		// Create a map to hold the column name and value pairs
		rowMap := make(map[string]interface{})
		for i, colName := range columns {
			val := columnValues[i]

			// Convert []byte to string for readability in JSON
			if b, ok := val.([]byte); ok {
				rowMap[colName] = string(b)
			} else {
				rowMap[colName] = val
			}
		}

		results = append(results, rowMap)
	}

	// Check for errors from iterating over rows
	if err = rows.Err(); err != nil {
		http.Error(w, fmt.Sprintf("Error iterating over rows: %v", err), http.StatusInternalServerError)
		return
	}

	// Set the content type to application/json
	w.Header().Set("Content-Type", "application/json")

	// Encode the results to JSON and write to the response writer
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(results)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error encoding JSON: %v", err), http.StatusInternalServerError)
	}
}

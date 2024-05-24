package main

import (
	"fmt"
	"net/http"
)

func registerReturnEndpoints() {
	// Dangerous
	http.HandleFunc("/clearAllData", withCORS(clearAllData))
	http.HandleFunc("/deleteAllTables", withCORS(deleteAllTables))

	// Register HTTP handlers for each table with CORS handling
	http.HandleFunc("/exhibits", withCORS(handleExhibit))
	http.HandleFunc("/employees", withCORS(handleEmployee))
	http.HandleFunc("/sponsorships", withCORS(handleSponsorship))
	http.HandleFunc("/notes", withCORS(handleNote))
	http.HandleFunc("/parts", withCORS(handlePart))
	http.HandleFunc("/jotforms", withCORS(handleJotForm))
	http.HandleFunc("/issues", withCORS(handleIssue))
}

func clearAllData(w http.ResponseWriter, r *http.Request) {
	tables := []string{
		"Employee",
		"Issue",
		"Exhibit",
		"Sponsorship",
		"Note",
		"Part",
		"JotForm",
		// Add other tables as needed
	}

	tx, err := db.Begin()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error starting transaction: %v", err), http.StatusInternalServerError)
		return
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			http.Error(w, fmt.Sprintf("Error clearing data: %v", p), http.StatusInternalServerError)
		}
	}()

	// Disable foreign key checks
	if _, err := tx.Exec("SET FOREIGN_KEY_CHECKS = 0"); err != nil {
		tx.Rollback()
		http.Error(w, fmt.Sprintf("Error disabling foreign key checks: %v", err), http.StatusInternalServerError)
		return
	}

	for _, table := range tables {
		_, err := tx.Exec(fmt.Sprintf("DELETE FROM %s", table))
		if err != nil {
			tx.Rollback()
			http.Error(w, fmt.Sprintf("Error deleting data from table %s: %v", table, err), http.StatusInternalServerError)
			return
		}

		// Reset AUTO_INCREMENT value
		_, err = tx.Exec(fmt.Sprintf("ALTER TABLE %s AUTO_INCREMENT = 1", table))
		if err != nil {
			tx.Rollback()
			http.Error(w, fmt.Sprintf("Error resetting AUTO_INCREMENT for table %s: %v", table, err), http.StatusInternalServerError)
			return
		}
	}

	// Enable foreign key checks
	if _, err := tx.Exec("SET FOREIGN_KEY_CHECKS = 1"); err != nil {
		tx.Rollback()
		http.Error(w, fmt.Sprintf("Error enabling foreign key checks: %v", err), http.StatusInternalServerError)
		return
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		http.Error(w, fmt.Sprintf("Error committing transaction: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("All data cleared successfully"))
}

func deleteAllTables(w http.ResponseWriter, r *http.Request) {
	tables := []string{
		"Employee",
		"Issue",
		"Exhibit",
		"Sponsorship",
		"Note",
		"Part",
		"JotForm",
		// Add other tables as needed
	}

	tx, err := db.Begin()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error starting transaction: %v", err), http.StatusInternalServerError)
		return
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			http.Error(w, fmt.Sprintf("Error deleting tables: %v", p), http.StatusInternalServerError)
		}
	}()

	// Disable foreign key checks
	if _, err := tx.Exec("SET FOREIGN_KEY_CHECKS = 0"); err != nil {
		tx.Rollback()
		http.Error(w, fmt.Sprintf("Error disabling foreign key checks: %v", err), http.StatusInternalServerError)
		return
	}

	for _, table := range tables {
		_, err := tx.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", table))
		if err != nil {
			tx.Rollback()
			http.Error(w, fmt.Sprintf("Error deleting table %s: %v", table, err), http.StatusInternalServerError)
			return
		}
	}

	// Enable foreign key checks
	if _, err := tx.Exec("SET FOREIGN_KEY_CHECKS = 1"); err != nil {
		tx.Rollback()
		http.Error(w, fmt.Sprintf("Error enabling foreign key checks: %v", err), http.StatusInternalServerError)
		return
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		http.Error(w, fmt.Sprintf("Error committing transaction: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("All tables deleted successfully"))
}

func handleExhibit(w http.ResponseWriter, r *http.Request) {
	if err := exportTableToJSON(w, "Exhibit"); err != nil {
		http.Error(w, fmt.Sprintf("Error exporting table Exhibit: %v", err), http.StatusInternalServerError)
	}
}

func handleEmployee(w http.ResponseWriter, r *http.Request) {
	if err := exportTableToJSON(w, "Employee"); err != nil {
		http.Error(w, fmt.Sprintf("Error exporting table Employee: %v", err), http.StatusInternalServerError)
	}
}

func handleSponsorship(w http.ResponseWriter, r *http.Request) {
	if err := exportTableToJSON(w, "Sponsorship"); err != nil {
		http.Error(w, fmt.Sprintf("Error exporting table Sponsorship: %v", err), http.StatusInternalServerError)
	}
}

func handleNote(w http.ResponseWriter, r *http.Request) {
	if err := exportTableToJSON(w, "Note"); err != nil {
		http.Error(w, fmt.Sprintf("Error exporting table Note: %v", err), http.StatusInternalServerError)
	}
}

func handlePart(w http.ResponseWriter, r *http.Request) {
	if err := exportTableToJSON(w, "Part"); err != nil {
		http.Error(w, fmt.Sprintf("Error exporting table Part: %v", err), http.StatusInternalServerError)
	}
}

func handleJotForm(w http.ResponseWriter, r *http.Request) {
	if err := exportTableToJSON(w, "JotForm"); err != nil {
		http.Error(w, fmt.Sprintf("Error exporting table JotForm: %v", err), http.StatusInternalServerError)
	}
}

func handleIssue(w http.ResponseWriter, r *http.Request) {
	if err := exportTableToJSON(w, "Issue"); err != nil {
		http.Error(w, fmt.Sprintf("Error exporting table Issue: %v", err), http.StatusInternalServerError)
	}
}

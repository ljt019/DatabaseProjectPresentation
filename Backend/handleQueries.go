package main

import (
	"fmt"
	"net/http"
)

func registerQueryEndpoints() {
	// Create Table Endpoints
	http.HandleFunc("/runCreateEmployeeTable", withCORS(CreateEmployeeTable))
	http.HandleFunc("/runCreateExhibitTable", withCORS(CreateExhibitTable))
	http.HandleFunc("/runCreateIssueTable", withCORS(CreateIssueTable))
	http.HandleFunc("/runCreatePartTable", withCORS(CreatePartTable))
	http.HandleFunc("/runCreateNoteTable", withCORS(CreateNoteTable))
	http.HandleFunc("/runCreateSponsorshipTable", withCORS(CreateSponsorshipTable))
	http.HandleFunc("/runCreateJotformTable", withCORS(CreateJotformTable))

	// Insert Endpoints
	http.HandleFunc("/runAddEmployeeQuery", withCORS(InsertEmployeeQueries))
	http.HandleFunc("/runAddExhibitQuery", withCORS(InsertExhibitsQueries))
	http.HandleFunc("/runAddIssueQuery", withCORS(InsertIssuesQueries))
	http.HandleFunc("/runAddPartQuery", withCORS(InsertPartsQueries))
	http.HandleFunc("/runAddNoteQuery", withCORS(InsertNotesQueries))
	http.HandleFunc("/runAddSponsorshipQuery", withCORS(InsertSponsorshipsQueries))
	http.HandleFunc("/runAddJotformQuery", withCORS(InsertJotformsQueries))

	// Update Endpoints
	http.HandleFunc("/runUpdateEmployeeQuery", withCORS(UpdateEmployeeQuery))
	http.HandleFunc("/runUpdateExhibitQuery", withCORS(UpdateExhibitQuery))

	//Delete Endpoints
	http.HandleFunc("/runDeleteEmployeeQuery", withCORS(DeleteEmployeeQuery))
	http.HandleFunc("/runDeleteExhibitQuery", withCORS(DeleteExhibitQuery))

	//Schema Modification Endpoints
	http.HandleFunc("/runModifyEmployeeQuery", withCORS(ModifyEmployeeQuery))

	// Trigger Endpoints
	http.HandleFunc("/runRegisterTriggerBeforeInsertIssueEmployee", withCORS(RegisterTriggerBeforeInsertIssueEmployee))
	http.HandleFunc("/runTestTriggerBeforeInsertIssueEmployee", withCORS(TestTriggerBeforeInsertIssueEmployee))

	http.HandleFunc("/runRegisterTriggerBeforeInsertIssueStatus", withCORS(RegisterTriggerBeforeInsertIssueStatus))
	http.HandleFunc("/runTestTriggerBeforeInsertIssueStatus", withCORS(TestTriggerBeforeInsertIssueStatus))

	http.HandleFunc("/runRegisterBeforeInsertJotForm", withCORS(RegisterBeforeInsertJotForm))
	http.HandleFunc("/runTestBeforeInsertJotForm", withCORS(TestBeforeInsertJotForm))

	http.HandleFunc("/runRegisterBeforeInsertNote", withCORS(RegisterBeforeInsertNote))
	http.HandleFunc("/runTestBeforeInsertNote", withCORS(TestBeforeInsertNote))

	http.HandleFunc("/runRegisterAfterInsertIssueEmployee", withCORS(RegisterAfterInsertIssueEmployee))
	http.HandleFunc("/runTestAfterInsertIssueEmployee", withCORS(TestAfterInsertIssueEmployee))

	// View Endpoints
	http.HandleFunc("/runViewEmployeeIssues", withCORS(ViewEmployeeIssues))
	http.HandleFunc("/runViewExhibitDetails", withCORS(ViewExhibitDetails))
	http.HandleFunc("/runViewJotFormIssues", withCORS(ViewJotFormIssues))
	http.HandleFunc("/runViewOperationalNeedsRepair", withCORS(ViewOperationalNeedsRepair))
	http.HandleFunc("/runViewPartsPendingOrders", withCORS(ViewPartsPendingOrders))

	// Retrieval Endpoints
	http.HandleFunc("/runRetrievalQuery1", withCORS(RetrievalQuery1))
	http.HandleFunc("/runRetrievalQuery2", withCORS(RetrievalQuery2))
	http.HandleFunc("/runRetrievalQuery3", withCORS(RetrievalQuery3))
	http.HandleFunc("/runRetrievalQuery4", withCORS(RetrievalQuery4))
	http.HandleFunc("/runRetrievalQuery5", withCORS(RetrievalQuery5))
	http.HandleFunc("/runRetrievalQuery6", withCORS(RetrievalQuery6))

}

func CreateEmployeeTable(w http.ResponseWriter, r *http.Request) {
	err := executeSQLFile("./queries/creates/createEmployeeTable.sql")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send a response
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Employee table created successfully")
}

func CreateExhibitTable(w http.ResponseWriter, r *http.Request) {
	err := executeSQLFile("./queries/creates/createExhibitTable.sql")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send a response
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Exhibit table created successfully")
}

func CreateIssueTable(w http.ResponseWriter, r *http.Request) {
	err := executeSQLFile("./queries/creates/createIssueTable.sql")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send a response
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Issue table created successfully")
}

func CreatePartTable(w http.ResponseWriter, r *http.Request) {
	err := executeSQLFile("./queries/creates/createPartTable.sql")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send a response
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Part table created successfully")
}

func CreateNoteTable(w http.ResponseWriter, r *http.Request) {
	err := executeSQLFile("./queries/creates/createNoteTable.sql")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send a response
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Note table created successfully")
}

func CreateSponsorshipTable(w http.ResponseWriter, r *http.Request) {
	err := executeSQLFile("./queries/creates/createSponsorshipTable.sql")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send a response
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Sponsorship table created successfully")
}

func CreateJotformTable(w http.ResponseWriter, r *http.Request) {
	err := executeSQLFile("./queries/creates/createJotformTable.sql")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send a response
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Jotform table created successfully")
}

func InsertEmployeeQueries(w http.ResponseWriter, r *http.Request) {
	err := executeSQLFile("./queries/inserts/insertEmployees.sql")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send a response
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Employee added successfully")
}

func InsertExhibitsQueries(w http.ResponseWriter, r *http.Request) {
	err := executeSQLFile("./queries/inserts/insertExhibits.sql")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send a response
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Employee added successfully")
}

func InsertIssuesQueries(w http.ResponseWriter, r *http.Request) {
	err := executeSQLFile("./queries/inserts/insertIssues.sql")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send a response
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Employee added successfully")
}

func InsertPartsQueries(w http.ResponseWriter, r *http.Request) {
	err := executeSQLFile("./queries/inserts/insertParts.sql")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send a response
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Employee added successfully")
}

func InsertNotesQueries(w http.ResponseWriter, r *http.Request) {
	err := executeSQLFile("./queries/inserts/insertNotes.sql")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send a response
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Employee added successfully")
}

func InsertSponsorshipsQueries(w http.ResponseWriter, r *http.Request) {
	err := executeSQLFile("./queries/inserts/insertSponsorships.sql")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send a response
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Employee added successfully")
}

func InsertJotformsQueries(w http.ResponseWriter, r *http.Request) {
	err := executeSQLFile("./queries/inserts/insertJotforms.sql")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send a response
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Employee added successfully")
}

func UpdateEmployeeQuery(w http.ResponseWriter, r *http.Request) {
	err := executeSQLFile("./queries/updates/updateEmployee.sql")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send a response
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Employee updated successfully")
}

func UpdateExhibitQuery(w http.ResponseWriter, r *http.Request) {
	err := executeSQLFile("./queries/updates/updateExhibit.sql")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send a response
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Exhibit updated successfully")
}

func DeleteEmployeeQuery(w http.ResponseWriter, r *http.Request) {
	err := executeSQLFile("./queries/deletes/deleteEmployee.sql")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send a response
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Employee deleted successfully")
}

func DeleteExhibitQuery(w http.ResponseWriter, r *http.Request) {
	err := executeSQLFile("./queries/deletes/deleteExhibit.sql")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send a response
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Exhibit deleted successfully")
}

func ModifyEmployeeQuery(w http.ResponseWriter, r *http.Request) {
	err := executeSQLFile("./queries/schema-mod/modifyEmployees.sql")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send a response
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Employee schema modified successfully")
}

func RegisterTriggerBeforeInsertIssueEmployee(w http.ResponseWriter, r *http.Request) {
	err := executeSQLFile("./queries/triggers/registerTriggerBeforeInsertIssueEmployee.sql")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send a response
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Trigger registered successfully")
}

func TestTriggerBeforeInsertIssueEmployee(w http.ResponseWriter, r *http.Request) {
	err := executeSQLFile("./queries/triggers/testBeforeInsertIssueEmployee.sql")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send a response
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Trigger tested successfully")
}

func RegisterTriggerBeforeInsertIssueStatus(w http.ResponseWriter, r *http.Request) {
	err := executeSQLFile("./queries/triggers/registerTriggerBeforeInsertIssueStatus.sql")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send a response
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Trigger registered successfully")
}

func TestTriggerBeforeInsertIssueStatus(w http.ResponseWriter, r *http.Request) {
	err := executeSQLFile("./queries/triggers/testBeforeInsertIssueStatus.sql")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send a response
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Trigger tested successfully")
}

func RegisterBeforeInsertJotForm(w http.ResponseWriter, r *http.Request) {
	err := executeSQLFile("./queries/triggers/registerBeforeInsertJotForm.sql")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send a response
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Trigger registered successfully")
}

func TestBeforeInsertJotForm(w http.ResponseWriter, r *http.Request) {
	err := executeSQLFile("./queries/triggers/testBeforeInsertJotForm.sql")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send a response
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Trigger tested successfully")
}

func RegisterBeforeInsertNote(w http.ResponseWriter, r *http.Request) {
	err := executeSQLFile("./queries/triggers/registerBeforeInsertNote.sql")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send a response
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Trigger registered successfully")
}

func TestBeforeInsertNote(w http.ResponseWriter, r *http.Request) {
	err := executeSQLFile("./queries/triggers/testBeforeInsertNote.sql")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send a response
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Trigger tested successfully")
}

func RegisterAfterInsertIssueEmployee(w http.ResponseWriter, r *http.Request) {
	err := executeSQLFile("./queries/triggers/registerAfterInsertIssueEmployee.sql")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send a response
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Trigger registered successfully")
}

func TestAfterInsertIssueEmployee(w http.ResponseWriter, r *http.Request) {
	err := executeSQLFile("./queries/triggers/testAfterInsertIssueEmployee.sql")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send a response
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Trigger tested successfully")
}

func ViewEmployeeIssues(w http.ResponseWriter, r *http.Request) {
	// Execute the SQL to create the view
	err := executeSQLFile("./queries/views/viewEmployeeIssues.sql")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating view EmployeeIssues: %v", err), http.StatusInternalServerError)
		return
	}

	// Use the exportTableToJSON function to fetch and return data from the view
	if err := exportTableToJSON(w, "EmployeeIssues"); err != nil {
		http.Error(w, fmt.Sprintf("Error exporting table EmployeeIssues: %v", err), http.StatusInternalServerError)
	}
}

func ViewExhibitDetails(w http.ResponseWriter, r *http.Request) {
	// Execute the SQL to create the view
	err := executeSQLFile("./queries/views/viewExhibitDetails.sql")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating view ExhibitDetails: %v", err), http.StatusInternalServerError)
		return
	}

	// Use the exportTableToJSON function to fetch and return data from the view
	if err := exportTableToJSON(w, "ExhibitDetails"); err != nil {
		http.Error(w, fmt.Sprintf("Error exporting table ExhibitDetails: %v", err), http.StatusInternalServerError)
	}
}

func ViewJotFormIssues(w http.ResponseWriter, r *http.Request) {
	// Execute the SQL to create the view
	err := executeSQLFile("./queries/views/viewJotFormIssues.sql")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating view JotFormIssues: %v", err), http.StatusInternalServerError)
		return
	}

	// Use the exportTableToJSON function to fetch and return data from the view
	if err := exportTableToJSON(w, "JotFormIssues"); err != nil {
		http.Error(w, fmt.Sprintf("Error exporting table JotFormIssues: %v", err), http.StatusInternalServerError)
	}
}

func ViewOperationalNeedsRepair(w http.ResponseWriter, r *http.Request) {
	// Execute the SQL to create the view
	err := executeSQLFile("./queries/views/viewOperationalNeedsRepair.sql")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating view OperationalNeedsRepair: %v", err), http.StatusInternalServerError)
		return
	}

	// Use the exportTableToJSON function to fetch and return data from the view
	if err := exportTableToJSON(w, "OperationalNeedsRepair"); err != nil {
		http.Error(w, fmt.Sprintf("Error exporting table OperationalNeedsRepair: %v", err), http.StatusInternalServerError)
	}
}

func ViewPartsPendingOrders(w http.ResponseWriter, r *http.Request) {
	// Execute the SQL to create the view
	err := executeSQLFile("./queries/views/viewPartsPendingOrders.sql")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating view PartsPendingOrders: %v", err), http.StatusInternalServerError)
		return
	}

	// Use the exportTableToJSON function to fetch and return data from the view
	if err := exportTableToJSON(w, "PartsPendingOrders"); err != nil {
		http.Error(w, fmt.Sprintf("Error exporting table PartsPendingOrders: %v", err), http.StatusInternalServerError)
	}
}

func RetrievalQuery1(w http.ResponseWriter, r *http.Request) {
	query := `
	SELECT i.IssueID, i.RepairDescription, i.PriorityLevel, e.FirstName, e.LastName
	FROM Issue i
	JOIN Employee e ON i.EMPLOYEE_AccountID = e.AccountID
	WHERE i.CompletedStatus = 'Pending';`
	exportQueryToJSON(w, query)
}

func RetrievalQuery2(w http.ResponseWriter, r *http.Request) {
	query := `
	SELECT COUNT(*) AS OperationalExhibits
	FROM Exhibit
	WHERE CurrentStatus = 'Operational';`
	exportQueryToJSON(w, query)
}

func RetrievalQuery3(w http.ResponseWriter, r *http.Request) {
	query := `
	SELECT AccountID, PreferredName
	FROM Employee
	WHERE Role IN ('Employee', 'Manager');`
	exportQueryToJSON(w, query)
}

func RetrievalQuery4(w http.ResponseWriter, r *http.Request) {
	query := `
	SELECT ex.Name AS ExhibitName, n.Note, n.Date
	FROM Exhibit ex
	JOIN Note n ON ex.Name = n.EXHIBIT_Name
	WHERE ex.Name = 'Moon Chair';`
	exportQueryToJSON(w, query)
}

func RetrievalQuery5(w http.ResponseWriter, r *http.Request) {
	query := `
	SELECT ex.Name AS ExhibitName, s.SponsorName, s.StartDate, s.EndDate
	FROM Exhibit ex
	JOIN Sponsorship s ON ex.Name = s.EXHIBIT_Name;`
	exportQueryToJSON(w, query)
}

func RetrievalQuery6(w http.ResponseWriter, r *http.Request) {
	query := `
	SELECT e.FirstName, e.LastName
	FROM Employee e
	WHERE e.AccountID IN (
		SELECT i.EMPLOYEE_AccountID
		FROM Issue i
		JOIN Exhibit ex ON i.EXHIBIT_Name = ex.Name
		WHERE ex.CurrentStatus = 'Needs Repair'
	);`
	exportQueryToJSON(w, query)
}

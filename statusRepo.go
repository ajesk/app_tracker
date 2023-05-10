package main

import "fmt"

func queryStatusesByApplication(applicationId int) ([]status, error) {
	db := getDb()
	statuses := []status{}

	rows, err := db.Query(
		fmt.Sprintf(
			`SELECT id, type, note, timestamp 
			FROM %s
			WHERE application_id = %d`,
			ApplicationStatusTable, applicationId),
	)
	if err != nil {
		return statuses, err
	}

	defer rows.Close()

	for rows.Next() {
		var row status

		err := rows.Scan(&row.ID, &row.Type, &row.Note, &row.Timestamp)
		if err != nil {
			return statuses, err
		}

		statuses = append(statuses, row)
	}

	return statuses, nil
}

func insertApplicationStatus(stat status, applicationId int) (int, error) {
	db := getDb()
	defer db.Close()
	var id int

	err := db.QueryRow(fmt.Sprintf(
		`INSERT INTO %s (application_id, type, note) 
		VALUES($1, $2, $3) 
		RETURNING id`, ApplicationStatusTable),
		applicationId, stat.Type, stat.Note,
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

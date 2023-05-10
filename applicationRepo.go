package main

import "fmt"

func queryAllApplications() ([]application, error) {
	db := getDb()
	applications := []application{}

	rows, err := db.Query(
		fmt.Sprintf(
			`SELECT a.id, a.title, a.max_salary, a.min_salary, a.notes, c.name
			FROM %s a
			JOIN %s c ON a.company = c.id`, ApplicationTable, CompanyTable),
	)

	if err != nil {
		return applications, err
	}

	defer rows.Close()

	for rows.Next() {
		var row application

		err := rows.Scan(&row.ID, &row.Title, &row.MaxSalary, &row.MinSalary, &row.Notes, &row.Company)
		if err != nil {
			return applications, err
		}

		applications = append(applications, row)
	}

	return applications, nil
}

func queryApplicationById(id int) (application, error) {
	db := getDb()
	defer db.Close()
	var app application

	row := db.QueryRow(
		fmt.Sprintf(
			`SELECT a.id, a.title, a.max_salary, a.min_salary, a.notes, c.name
			FROM %s a
			JOIN %s c ON a.company = c.id
			WHERE a.id = %d`, ApplicationTable, CompanyTable, id),
	)
	err := row.Scan(&app.ID, &app.Title, &app.MaxSalary, &app.MinSalary, &app.Notes, &app.Company)
	if err != nil {
		return app, err
	}

	return app, nil
}

func insertApplication(app application, companyId int, userId int) (int, error) {
	db := getDb()
	defer db.Close()
	var id int

	err := db.QueryRow(fmt.Sprintf(
		`INSERT INTO %s (title, max_salary, min_salary, notes, company, applicant) 
		VALUES($1, $2, $3, $4, $5, $6) 
		RETURNING id`, ApplicationTable),
		app.Title, app.MaxSalary, app.MinSalary, app.Notes, companyId, userId,
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

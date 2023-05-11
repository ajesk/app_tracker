package main

func queryAllCompanies() ([]company, error) {
	db := getDb()
	companies := []company{}

	rows, err := db.Query("SELECT id, name, industry FROM " + CompanyTable)
	if err != nil {
		return companies, err
	}

	defer rows.Close()

	for rows.Next() {
		var row company

		err := rows.Scan(&row.ID, &row.Name, &row.Industry)
		if err != nil {
			return companies, err
		}

		companies = append(companies, row)
	}

	return companies, nil
}

func queryCompanyById(id int) (company, error) {
	db := getDb()
	defer db.Close()
	var co company

	row := db.QueryRow("SELECT id, name, industry FROM "+CompanyTable+" WHERE id = $1", id)

	err := row.Scan(&co.ID, &co.Name, &co.Industry)
	if err != nil {
		return co, err
	}

	return co, nil
}

func insertCompany(co company) (int, error) {
	db := getDb()
	defer db.Close()
	var id int

	err := db.QueryRow("INSERT INTO "+CompanyTable+"(name, industry) VALUES($1, $2) RETURNING id", co.Name, co.Industry).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

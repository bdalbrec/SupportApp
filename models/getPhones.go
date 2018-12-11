package models

type Phone struct{
	Name string
	Number string
}

func GetPhones() ([]*Phone, error){
	rows, err := db.Query("SELECT Name, Number FROM Phone ORDER BY Name")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ps := make([]*Phone, 0)
	for rows.Next() {
		p := new(Phone)
		err := rows.Scan(&p.Name, &p.Number)
		if err != nil {
			return nil, err
		}
		ps = append(ps, p)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return ps, nil
}
package models

type Link struct{
	Name string
	Address string
}

func GetLinks(c string) ([]*Link, error){
	rows, err := db.Query("SELECT Name, Address FROM Links WHERE Category = ? ORDER BY Name", c)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ls := make([]*Link, 0)
	for rows.Next() {
		l := new(Link)
		err := rows.Scan(&l.Name, &l.Address)
		if err != nil {
			return nil, err
		}
		ls = append(ls, l)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return ls, nil
}
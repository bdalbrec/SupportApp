package models

type Category struct{
	Name string
	Number string
}

func GetCategories() ([]*Category, error){
	rows, err := db.Query("SELECT * FROM Category ORDER BY Number")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cs := make([]*Category, 0)
	for rows.Next() {
		c := new(Category)
		err := rows.Scan(&c.Name, &c.Number)
		if err != nil {
			return nil, err
		}
		cs = append(cs, c)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return cs, nil
}
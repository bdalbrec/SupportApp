package models

// Link struct already declared in getLinks.go
type FullLink struct {
	Id string
	Name string
	Address string
	Tags string
	Category string
}

func GetAllLinks() ([]*FullLink, error){
	rows, err := db.Query("SELECT Id, Name, Address, Tags, Category FROM Links ORDER BY Name")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ls := make([]*FullLink, 0)
	for rows.Next() {
		l := new(FullLink)
		err := rows.Scan(&l.Id, &l.Name, &l.Address, &l.Tags, &l.Category)
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
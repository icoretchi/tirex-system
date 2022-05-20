package locality_manager

import (
	"database/sql"
	"fmt"
)

//mySQLRepo mysql repo
type mySQLRepo struct {
	db *sql.DB
}

//NewMySQLRepo create new repository
func NewMySQLRepo(db *sql.DB) (Repository, error) {
	return &mySQLRepo{
		db: db,
	}, nil
}

//Create locality
func (r *mySQLRepo) Create(l *Locality) (LocalityCode, error) {
	stmt, err := r.db.Prepare(`
		insert into localities (code, statistical_code, name, status) 
		values(?,?,?,?)`)
	if err != nil {
		return l.Code, err
	}
	_, err = stmt.Exec(l.Code, l.StatisticalCode, l.Name, l.Status)
	if err != nil {
		return l.Code, err
	}
	err = stmt.Close()
	if err != nil {
		return l.Code, err
	}
	return l.Code, nil
}

//Get locality
func (r *mySQLRepo) Get(code LocalityCode) (*Locality, error) {
	return getLocality(code, r.db)
}

func (r *mySQLRepo) GetLocalityByCode(code LocalityCode) (interface{}, error) {
	l := Locality{}
	err := r.db.QueryRow("SELECT * FROM localities where code = ?", code).Scan(&l.Code, &l.StatisticalCode, &l.Name, &l.Status)

	if err != nil {
		return "", err
	}
	return l, nil
}

func getLocality(code LocalityCode, db *sql.DB) (*Locality, error) {
	stmt, err := db.Prepare(`select code, statistical_code, name, status from localities where code = ?`)
	if err != nil {
		return nil, err
	}
	var l Locality
	rows, err := stmt.Query(code)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&l.Code, &l.StatisticalCode, &l.Name, &l.Status)
	}
	return &l, nil
}

//Update locality
func (r *mySQLRepo) Update(l *Locality) error {
	_, err := r.db.Exec("update localities set code = ?, statistical_code = ?, name = ?, status = ? where code = ?",
		l.Code, l.StatisticalCode, l.Name, l.Status, l.Code)
	if err != nil {
		return err
	}
	return nil
}

//Search localities
func (r *mySQLRepo) Search(query string) ([]*Locality, error) {
	stmt, err := r.db.Prepare(`select code from localities where name like ?`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var ids []LocalityCode
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var i LocalityCode
		err = rows.Scan(&i)
		if err != nil {
			return nil, err
		}
		ids = append(ids, i)
	}
	if len(ids) == 0 {
		return nil, fmt.Errorf("not found")
	}
	var localities []*Locality
	for _, id := range ids {
		u, err := getLocality(id, r.db)
		if err != nil {
			return nil, err
		}
		localities = append(localities, u)
	}
	return localities, nil
}

//List users
func (r *mySQLRepo) List() ([]*Locality, error) {
	stmt, err := r.db.Prepare(`select code from localities`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var ids []LocalityCode
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var i LocalityCode
		err = rows.Scan(&i)
		if err != nil {
			return nil, err
		}
		ids = append(ids, i)
	}
	if len(ids) == 0 {
		return nil, fmt.Errorf("not found")
	}
	var localities []*Locality
	for _, id := range ids {
		l, err := getLocality(id, r.db)
		if err != nil {
			return nil, err
		}
		localities = append(localities, l)
	}
	return localities, nil
}

//Delete locality
func (r *mySQLRepo) Delete(code LocalityCode) error {
	_, err := r.db.Exec("delete from localities where code = ?", code)
	if err != nil {
		return err
	}
	return nil
}

package database

import (
	"fmt"
	"gym/ankur/models"
)

func (m *Repository) GetGyms(owner_id float64) []models.Gym {
	sqlQuery := `SELECT id,name,address,owner_id FROM gyms WHERE owner_id=$1`
	rows, err := m.Db.Query(sqlQuery, owner_id)
	if err != nil {
		panic(err)
	}
	var ListOfGyms []models.Gym
	for rows.Next() {
		var GymDetails models.Gym
		err := rows.Scan(&GymDetails.Id, &GymDetails.Name, &GymDetails.Address, &GymDetails.Owner_id)
		if err != nil {
			panic(err)
		}
		ListOfGyms = append(ListOfGyms, GymDetails)
	}
	return ListOfGyms
}

func (m *Repository) CreateGym(name string, address string, owner_id int) string {
	sqlQuery := `INSERT INTO gyms(name,address,owner_id) values($1,$2,$3)`
	_, err := m.Db.Exec(sqlQuery, name, address, owner_id)

	if err != nil {
		panic(err)
	}
	return "New Gym Created"
}
func (m *Repository) GymId(name string, cid int) {
	sql := `SELECT id FROM gyms where name=$1 `
	var gym_id int
	row := m.Db.QueryRow(sql, name)
	row.Scan(&gym_id)
	sql2 := `UPDATE customers SET gym_id=$1 WHERE customer_id=$2`
	_, err := m.Db.Exec(sql2, gym_id, cid)
	if err != nil {

		panic(err)
	}
}

func (m *Repository) DeleteGymsOfOwners(owenr_id int) {

	sql := `DELETE FROM gyms WHERE owner_id=$1`
	_, err := m.Db.Exec(sql, owenr_id)
	if err != nil {
		panic(err)
	}
}

func (m *Repository) DeleteGyms(gym_id int, owner_id int) {
	fmt.Println(gym_id)
	sql0 := `DELETE FROM customers WHERE gym_id =$1`
	_, err := m.Db.Exec(sql0, gym_id)
	if err != nil {
		panic(err)
	}
	sql := `DELETE FROM gyms WHERE id=$1 and owner_id=$2`
	_, err = m.Db.Exec(sql, gym_id, owner_id)
	if err != nil {
		panic(err)
	}
}

package database

import (
	"fmt"
	"gym/ankur/models"
)

func (m *Repository) GetCustomerDetails(owner_id int) []models.CustomerDetails {
	var ListOfCustomers []models.CustomerDetails
	sqlQuery := `SELECT u.id,u.name,gyms.name,cu.current_weight,cu.target_weight,u2.name
	FROM users as u
	inner join customers as cu
	on cu.customer_id=u.id
	inner join gyms as gyms
	on gyms.id=cu.gym_id
	left join users as u2
	on cu.trainer_id=u2.id`
	rows, err := m.Db.Query(sqlQuery)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var SingleCustomer models.CustomerDetails
		err := rows.Scan(&SingleCustomer.Id, &SingleCustomer.Name, &SingleCustomer.Gym_name, &SingleCustomer.Current_weight, &SingleCustomer.Targeted_Weight, &SingleCustomer.Personal_trainer)
		if err != nil {
			panic(err)
		}
		ListOfCustomers = append(ListOfCustomers, SingleCustomer)
	}
	return ListOfCustomers
}

func (m *Repository) WeightSection(curr int, tar int, cid int) {
	sql := `UPDATE customers SET current_weight=$1,target_weight=$2 WHERE customer_id=$3`
	_, err := m.Db.Exec(sql, curr, tar, cid)
	if err != nil {
		fmt.Println("line 54")
		panic(err)
	}
}
func (m *Repository) DeleteCustomer(id int, owner_id int) {
	sql := `DELETE from customers WHERE customer_id=$1 and gym_id in (SELECT id FROM gyms WHERE owner_id=$2)`
	_, err := m.Db.Exec(sql, id, owner_id)
	if err != nil {
		panic(err)
	}
	fmt.Println(id)
	sql1 := `DELETE FROM users WHERE id=$1`
	_, err = m.Db.Exec(sql1, id)
	if err != nil {
		panic(err)
	}
}

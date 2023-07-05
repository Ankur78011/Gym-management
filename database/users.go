package database

import (
	"fmt"
	"gym/ankur/models"

	"golang.org/x/crypto/bcrypt"
)

func (m *Repository) GetOwner() []models.Owner {
	sqlQuer := `SELECT id,name,email FROM users WHERE usertype='owner'`
	rows, err := m.Db.Query(sqlQuer)
	if err != nil {
		fmt.Println("Asdfghj")
		panic(err)

	}
	var ListOfOwners []models.Owner
	for rows.Next() {
		var Owner models.Owner
		err = rows.Scan(&Owner.Id, &Owner.Name, &Owner.Email)
		ListOfOwners = append(ListOfOwners, Owner)
	}
	return ListOfOwners
}

func (m *Repository) CreateOwner(name string, email string, password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		panic(err)
	}
	sqlQuery := `INSERT INTO users(usertype,name,email,password) values('owner',$1,$2,$3)`
	_, err = m.Db.Exec(sqlQuery, name, email, hashedPassword)
	if err != nil {
		return "something went wrong"
	}
	return "Owner Registered Successfully"
}

func (m *Repository) InserCustomer(name string) int {
	sql := `INSERT INTO users(usertype,name) values('customer',$1) Returning Id`
	LastId := m.Db.QueryRow(sql, name)
	var id int
	LastId.Scan(&id)
	sql_customer_table := `INSERT INTO customers (customer_id) values($1)`
	_, err := m.Db.Exec(sql_customer_table, id)
	if err != nil {
		panic(err)
	}
	return id
}
func (m *Repository) Personal_trainer(name string, customer_id int) {
	sql := `Select id FROM users WHERE usertype='trainer' AND name=$1`
	var id int
	row := m.Db.QueryRow(sql, name)
	err := row.Scan(&id)
	if id == 0 {
		m.CreateTrainers(name)
		sql2 := `Select id FROM users WHERE usertype='trainer' AND name=$1`
		row := m.Db.QueryRow(sql2, name)
		row.Scan(&id)
	}
	sql2 := `UPDATE customers SET trainer_id =$1 WHERE customer_id=$2`
	_, err = m.Db.Exec(sql2, id, customer_id)
	if err != nil {
		panic(err)
	}
}

func (m *Repository) GetTrainers() []models.Trainer {
	sql := `SELECT id,name FROM users WHERE usertype='trainer'`
	var ListOfTrainers []models.Trainer
	rows, err := m.Db.Query(sql)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var singleTrainer models.Trainer
		err := rows.Scan(&singleTrainer.Id, &singleTrainer.Name)
		if err != nil {
			panic(err)
		}
		ListOfTrainers = append(ListOfTrainers, singleTrainer)
	}
	return ListOfTrainers
}

func (m *Repository) CreateTrainers(name string) {
	sql := `INSERT INTO users (usertype,name) values('trainer',$1) `
	_, err := m.Db.Exec(sql, name)
	if err != nil {
		panic(err)
	}

}

func (m *Repository) DeleteOwnerFromCustomer(owenr_id int) {
	sql := `DELETE FROM customers WHERE  gym_id IN (SELECT id FROM gyms WHERE owner_id=$1)`
	_, err := m.Db.Exec(sql, owenr_id)
	if err != nil {
		panic(err)
	}
}

func (m *Repository) DelteOwnerFromUsers(owenr_id int) {
	sql := `DELETE FROM users WHERE id=$1`
	_, err := m.Db.Exec(sql, owenr_id)
	if err != nil {
		panic(err)
	}
}

func (m *Repository) DeleteTrainer(id int) {
	sql := `DELETE from  customers WHERE trainer_id=$1`
	_, err := m.Db.Exec(sql, id)
	if err != nil {
		panic(err)
	}
	sql1 := `DELETE from users WHERE id=$1`
	_, err = m.Db.Exec(sql1, id)
	if err != nil {
		panic(err)
	}
}
func (m *Repository) Login(email string, password string) (string, bool) {
	sql := `SELECT password FROM users WHERE email=$1`
	row := m.Db.QueryRow(sql, email)
	var hashPassword string
	err := row.Scan(&hashPassword)
	if err != nil {
		return "Email or password Incorrect", false
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	if err != nil {
		return "Email or password Incorrect", false
	}
	return "Welcome", true

}
func (m *Repository) GetId(email string) int {
	sql := `SELECT id from users WHERE email=$1`
	row := m.Db.QueryRow(sql, email)
	var id int
	row.Scan(&id)
	return id

}

func (m *Repository) ValidateUser(id float64) bool {
	sql := `SELECT name from users WHERE id=$1`
	var username string
	row := m.Db.QueryRow(sql, id)
	err := row.Scan(&username)
	if err != nil {
		return false
	}
	return true
}

func (m *Repository) GetReports(cid int) []models.Report {
	var ListOfReports []models.Report
	sql := `SELECT customer_id,TO_CHAR(created_at,'Month'),EXTRACT(YEAR from created_at),weight from stats WHERE customer_id=$1`
	rows, err := m.Db.Query(sql, cid)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var Report models.Report
		err := rows.Scan(&Report.Customer_id, &Report.Month, &Report.Year, &Report.Weight)
		if err != nil {
			panic(err)
		}
		ListOfReports = append(ListOfReports, Report)
	}
	return ListOfReports
}
func (m *Repository) InsertReports(cid int, weight int) {
	sql := `insert into stats(customer_id,weight) values($1,$2)`
	_, err := m.Db.Exec(sql, cid, weight)
	if err != nil {
		panic(err)
	}
}
func (m *Repository) CheckCustomer(id int) bool {
	sql := `SELECT id FROM users WHERE id =$1`
	var cid int
	res := m.Db.QueryRow(sql, id)
	err := res.Scan(&cid)
	if err != nil {
		return false
	}
	return true
}

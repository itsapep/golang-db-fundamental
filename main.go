package main

import (
	"fmt"
	"golang_db_fundamental/config"
	"golang_db_fundamental/repository"
	"golang_db_fundamental/usecase"
	"golang_db_fundamental/utils"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // kalau lupa import -> unknown driver "postgres"
)

// func main() {
// 	dbHost := "localhost"
// 	dbPort := "5432"
// 	dbUser := "postgres"
// 	dbPassword := "password"
// 	dbName := "db_enigma_shop"

// 	//urutan url koneksi ke db postgres
// 	//postgres://dbUser:dbPassword@dbHost:dbPort/dbName
// 	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)
// 	db, err := sqlx.Connect("postgres", dsn)
// 	if err != nil {
// 		panic(err.Error())
// 	} else {
// 		log.Println("Connected")
// 	}

// 	defer func(db *sqlx.DB) {
// 		err := db.Close()
// 		if err != nil {
// 			panic(err.Error())
// 		}
// 	}(db)

// 	/*
// 		memanipulasi: insert, update, delete
// 	*/

// 	//insert
// 	// customers := []map[string]interface{}{
// 	// 	{
// 	// 		"id":      "C002",
// 	// 		"name":    "Budi",
// 	// 		"address": "Jakarta",
// 	// 		"phone":   "089089",
// 	// 		"email":   "budi@mail.com",
// 	// 		"saldo":   100000,
// 	// 	},
// 	// 	{
// 	// 		"id":      "C003",
// 	// 		"name":    "Ani",
// 	// 		"address": "Medan",
// 	// 		"phone":   "89786321",
// 	// 		"email":   "ani@mail.com",
// 	// 		"saldo":   100000,
// 	// 	},
// 	// 	{
// 	// 		"id":      "C004",
// 	// 		"name":    "Joko",
// 	// 		"address": "Semarang",
// 	// 		"phone":   "089089",
// 	// 		"email":   "joko@mail.com",
// 	// 		"saldo":   100000,
// 	// 	},
// 	// }

// 	customer := model.Customer{
// 		Id:      "C009",
// 		Name:    "Jontor",
// 		Address: "Solo",
// 		Phone:   "09890709",
// 		Email:   "jontor@gmail.com",
// 		Balance: 900000,
// 	}
// 	// _, err = db.NamedExec(util.INSERT_CUSTOMER, &customer)
// 	// if err != nil {
// 	// 	log.Println(err)
// 	// } else {
// 	// 	log.Println("Insert success")
// 	// }

// 	// with prepared statement
// 	// -> menghindari sql injection
// 	// ->cobde lebih efisien karena bisa digunakan berulang kali
// 	// ->database,preparex(query)
// 	// ->MustExcel(value)
// 	stmt, err := db.Preparex(utils.INSERT_CUSTOMER_PS)
// 	if err != nil {
// 		log.Println(err.Error())
// 	}

// 	stmt.MustExec(customer.Id, customer.Name, customer.Address, customer.Phone, customer.Email, customer.Balance)
// 	log.Println("Insert success")

// }

// func main() {
// 	dsn := GetDataSourceName(newConfigWithEnv())
// 	db, err := sqlx.Connect("postgres", dsn)
// 	if err != nil {
// 		panic(err.Error())
// 	} else {
// 		fmt.Println("connected")
// 	}

// 	defer func(db *sqlx.DB) {
// 		err := db.Close()
// 		if err != nil {
// 			panic(err.Error())
// 		} else {
// 			fmt.Println("disconnected")
// 		}
// 	}(db)

// WITH PREPARED STATEMENT
// -> Menghindari adanya sql injection
// -> Code lebih efisien karena bisa digunakan berulang kali
// -> database.Preparex(query)
// -> stmt.MustExec(value)

// 	customerStruct := model.Customer{
// 		Id:      "C011",
// 		Name:    "Indo aPEP",
// 		Address: "tUBAN",
// 		Phone:   "08998765432",
// 		Email:   "indo.apep@gmail.com",
// 		Balance: 50000000,
// 	}
// 	stmt, err := db.Preparex(utils.INSERT_CUSTOMER_PS)
// 	if err != nil {
// 		log.Println(err.Error())
// 	} else {
// 		log.Println("success")
// 	}

// 	stmt.MustExec(customerStruct.Id, customerStruct.Name, customerStruct.Address, customerStruct.Phone, customerStruct.Email, customerStruct.Balance)
// }

// func newConfigWithEnv() Config {
// 	// dbHost := os.Getenv("GOROOT")
// 	dbHost := os.Getenv("TEMPDBHOST")
// 	dbPort := os.Getenv("TEMPDBPORT")
// 	dbName := os.Getenv("TEMPDBNAME")
// 	dbUser := os.Getenv("TEMPDBUSER")
// 	dbPassword := os.Getenv("TEMPDBPASSWORD")
// 	return Config{
// 		dbHost,
// 		dbPort,
// 		dbName,
// 		dbUser,
// 		dbPassword,
// 	}
// }

func main() {
	config := config.NewConfigDB()
	db := config.DbConn()
	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			utils.IsError(err)
		}
	}(db)
	cstRepo := repository.NewCustomerRepository(db)
	cstUse := usecase.NewCustomerUseCase(cstRepo)

	// INSERT
	// cstId := utils.GenerateId()
	// customer := model.Customer{
	// 	Id:      cstId,
	// 	Name:    "Jution Kirana",
	// 	Address: "Ragunan",
	// 	Phone:   "08292929",
	// 	Email:   "jutionck@gmail.com",
	// 	Balance: 150000,
	// }
	// cstUse.RegisterNewCustomer(&customer)

	// DELETE
	// customerId := "C004"
	// err := cstUse.DeleteCustomer(customerId)
	// if err != nil {
	// 	fmt.Println("error test")
	// 	fmt.Println(err.Error())
	// }

	// UPDATE
	//customerUpdate := model.Customer{
	//	Name:    "Jution Aja",
	//	Address: "Ragunan",
	//	Phone:   "08292929",
	//	Email:   "jutionck@gmail.com",
	//	Id:      "1c63f19d-e896-4116-a847-2dbb75d7eae8",
	//}
	//cstRepo.Update(&customerUpdate)

	// FIND
	// customer, err := cstUse.FindCustomerByID("1")
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// fmt.Println(customer)

	// Aggregation
	var customerCount int
	customerCount, _ = cstUse.GetTotalActiveCustomer()
	fmt.Println("Total customer aktif: ", customerCount)

	customerBalance, _ := cstUse.GetTotalBalanceActiveCustomer()
	fmt.Println("Total balance customer aktif: ", customerBalance)
}

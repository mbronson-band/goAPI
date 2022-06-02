package main

import (
	//"encoding/json"
	"fmt"
	"log"

	//"net/http"
	"database/sql"

	_ "github.com/snowflakedb/gosnowflake"
)

/*type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Test Title", Desc: "Test Description", Content: "bruh"},
	}
	fmt.Fprintf(w, "Hell ya")
	json.NewEncoder(w).Encode(articles)

}

func testPostArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test Post endpoint")
}

func homePage(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequests() {
	myRouter := mux.NewRouter.StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", testPostArticles).Methods("POST")

	log.Fatal(http.ListenAndServe(":8082", myRouter))

}
*/

const (
	// PersonTableName is the name of the table for the person model
	PersonTableName = "person"
	// PersonFirstNameCol is the column name of the model's first name
	PersonFirstNameCol = "first_name"
	// PersonLastNameCol is the column name of the model's last name
	PersonLastNameCol = "last_name"
	// PersonAgeCol is the column name of the model's age
	PersonAgeCol = "age"
)

type Person struct {
	FirstName string
	LastName  string
	Age       uint
}

func CreatePersonTable(db *sql.DB) (sql.Result, error) {
	return db.Exec(
		fmt.Sprintf("CREATE TABLE %s (%s varchar(255), %s varchar(255), %s int)",
			PersonTableName,
			PersonFirstNameCol,
			PersonLastNameCol,
			PersonAgeCol,
		),
	)
}

func InsertPerson(db *sql.DB, person Person) (sql.Result, error) {
	return db.Exec(
		fmt.Sprintf("INSERT INTO %s VALUES(?, ?, ?)", PersonTableName),
		person.FirstName,
		person.LastName,
		person.Age,
	)
}

func SelectPerson(db *sql.DB, firstName, lastName string, age uint, result *Person) error {
	row := db.QueryRow(
		fmt.Sprintf(
			"SELECT * FROM %s WHERE %s=? AND %s=? AND %s=?",
			PersonTableName,
			PersonFirstNameCol,
			PersonLastNameCol,
			PersonAgeCol,
		),
		firstName,
		lastName,
		age,
	)
	var retFirstName, retLastName string
	var retAge uint
	if err := row.Scan(&retFirstName, &retLastName, &retAge); err != nil {
		return err
	}
	result.FirstName = retFirstName
	result.LastName = retLastName
	result.Age = retAge
	return nil
}

func UpdatePerson(db *sql.DB, firstName, lastName string, age uint, newPerson Person) error {
	_, err := db.Exec(
		fmt.Sprintf(
			"UPDATE %s SET %s=?,%s=?,%s=? WHERE %s=? AND %s=? AND %s=?",
			PersonTableName,
			PersonFirstNameCol,
			PersonLastNameCol,
			PersonAgeCol,
			PersonFirstNameCol,
			PersonLastNameCol,
			PersonAgeCol,
		),
		newPerson.FirstName,
		newPerson.LastName,
		newPerson.Age,
		firstName,
		lastName,
		age,
	)
	return err
}

func DeletePerson(db *sql.DB, firstName, lastName string, age uint) error {
	_, err := db.Exec(
		fmt.Sprintf(
			"DELETE FROM %s WHERE %s=? AND %s=? AND %s=?",
			PersonTableName,
			PersonFirstNameCol,
			PersonLastNameCol,
			PersonAgeCol,
		),
		firstName,
		lastName,
		age,
	)
	return err
}

func main() {

	db, err := sql.Open("snowflake", "milesbronson:Gearmonkey1!@ipb74218.us-east-1/api_sample/public")
	if err != nil {
		log.Fatalf("error opening DB (%s)", err)
	}

	log.Printf("Creating new table")
	if _, crErr := CreatePersonTable(db); crErr != nil {
		log.Fatalf("Error creating table (%s)", crErr)
	}
	log.Printf("Created")

	me := Person{FirstName: "Miles", LastName: "Bronson", Age: 20}
	log.Printf("Inserting %+v into the DB", me)
	if _, insErr := InsertPerson(db, me); insErr != nil {
		log.Fatalf("Error inserting new person into the DB (%s)", insErr)
	}
	log.Printf("Inserted")

	log.Printf("Selecting person from the DB")
	selectedMe := Person{}
	if err := SelectPerson(db, me.FirstName, me.LastName, me.Age, &selectedMe); err != nil {
		log.Fatalf("Error selecting person from the DB (%s)", err)
	}
	log.Printf("Selected %+v from the DB", selectedMe)

	log.Printf("Updating person in the DB")
	updatedMe := Person{
		FirstName: "Miles",
		LastName:  "Bronson",
		Age:       25,
	}
	if err := UpdatePerson(db, selectedMe.FirstName, selectedMe.LastName, selectedMe.Age, updatedMe); err != nil {
		log.Fatalf("Error updating person in the DB (%s)", err)
	}

	log.Printf("Deleting person from the DB")
	if delErr := DeletePerson(db, selectedMe.FirstName, selectedMe.LastName, selectedMe.Age); delErr != nil {
		log.Fatalf("Error deleting person from the DB (%s)", delErr)
	}
	log.Printf("Deleted")

	//handleRequests()
}

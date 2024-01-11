package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type student struct {
    id    string
    name  string
    age   int
    grade int
}

func connect() (*sql.DB, error) {
    db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3307)/belajar_golang")
    if err != nil {
        return nil, err
    }

    return db, nil
}

func sqlPrepare() {
    db, err := connect()
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    defer db.Close()

    stmt, err := db.Prepare("select name, grade from tb_student where id = ?")
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    var result1 = student{}
    stmt.QueryRow("E001").Scan(&result1.name, &result1.grade)
    fmt.Printf("name: %s\ngrade: %d\n", result1.name, result1.grade)

    var result2 = student{}
    stmt.QueryRow("W001").Scan(&result2.name, &result2.grade)
    fmt.Printf("name: %s\ngrade: %d\n", result2.name, result2.grade)

    var result3 = student{}
    stmt.QueryRow("B001").Scan(&result3.name, &result3.grade)
    fmt.Printf("name: %s\ngrade: %d\n", result3.name, result3.grade)
}

func main() {
    sqlPrepare()
}

// Query Create Table
// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"

// 	_ "github.com/go-sql-driver/mysql"
// )

// func main() {

// 	dbUser := "root"
// 	dbPass := ""
// 	dbName := "belajar_golang"

// 	// Membuat string koneksi ke MySQL
// dataSourceName := fmt.Sprintf("%s:%s@tcp(localhost:3307)/%s", dbUser, dbPass, dbName)


// 	// Membuka koneksi ke MySQL
// 	db, err := sql.Open("mysql", dataSourceName)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	// Mengecek koneksi ke MySQL
// 	err = db.Ping()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Membuat tabel
// 	createTableQuery := `
// 	CREATE TABLE IF NOT EXISTS tb_student (
// 		id VARCHAR(5) NOT NULL,
// 		name VARCHAR(255) NOT NULL,
// 		age INT NOT NULL,
// 		grade INT NOT NULL
// 	  ) ENGINE=InnoDB DEFAULT CHARSET=latin1;
// 	`

// 	_, err = db.Exec(createTableQuery)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Tabel berhasil dibuat!")
// }


// Query add data
// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"

// 	_ "github.com/go-sql-driver/mysql"
// )

// func main() {
// 
// 	dbUser := "root"
// 	dbPass := ""
// 	dbName := "belajar_golang"

// 	// Membuat string koneksi ke MySQL
// 	dataSourceName := fmt.Sprintf("%s:%s@tcp(localhost:3307)/%s", dbUser, dbPass, dbName)

// 	// Membuka koneksi ke MySQL
// 	db, err := sql.Open("mysql", dataSourceName)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	// Mengecek koneksi ke MySQL
// 	err = db.Ping()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Menambah data tabel
// 	insertTableQuery := `
// 		INSERT INTO tb_student (id, name, age, grade) VALUES
// 		('B001', 'Jason Bourne', 29, 1),
// 		('B002', 'James Bond', 27, 1),
// 		('E001', 'Ethan Hunt', 27, 2),
// 		('W001', 'John Wick', 28, 2);
// 	`

// 	_, err = db.Exec(insertTableQuery)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Menambahkan primary key setelah data dimasukkan
// 	alterTableQuery := `
// 		ALTER TABLE tb_student ADD PRIMARY KEY (id);
// 	`

// 	_, err = db.Exec(alterTableQuery)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Tabel berhasil dibuat!")
// }

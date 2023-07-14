package main

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
	Age  int
}
type Join struct {
	Name  string
	Age   int
	Email string
	// gorm.Model
}
type Total struct {
	Date  time.Time
	Total int
}

func main() {
	// Connection -->
	dsn := "host=localhost user=postgres password=danotbebek dbname=test_db_camp port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// // Migrate the schema -->
	// db.AutoMigrate(&User{})

	// // Create -->
	// db.Create(&User{Name: "Aditira", Age: 23})

	// // Read -->
	// // SELECT * FROM "users" WHERE "users"."id" = 1 AND "users"."deleted_at" IS NULL ORDER BY "users"."id" LIMIT 1
	// var user User
	// db.First(&user, 1)                     // temukan user dengan menggunakan primary key dan simpan di variabel user
	// db.First(&user, "name = ?", "Aditira") // temukan user dengan nama Aditira

	// fmt.Println(user)

	// // Insert 1 -->
	// user := User{Name: "Aditira", Age: 24}
	// result := db.Create(&user)
	// // INSERT INTO users (name, age, created_at, updated_at, deleted_at) VALUES ('Aditira', 24, NOW(), NOW(), NULL) RETURNING id

	// fmt.Println("Error: ", result.Error)       // mengembalikan error
	// fmt.Println("ID: ", user.ID)               // mengembalikan primary key dari data yang dimasukkan
	// fmt.Println("Rows: ", result.RowsAffected) // mengembalikan jumlah records yang dimasukkan

	// Insert 2 -->
	// user := User{Name: "Aditira", Age: 24}
	// db.Select("Name", "Age").Create(&user)
	// INSERT INTO users (name, age, created_at, updated_at, deleted_at) VALUES ('Aditira', 24, NOW(), NOW(), NULL)

	// Insert 3 -->
	// user := User{Name: "Aditira", Age: 24}
	// db.Omit("Age").Create(&user)
	// INSERT INTO users (name, created_at, updated_at, deleted_at) VALUES ('Aditira', NOW(), NOW(), NULL)

	// // Insert 4 -->
	// users := [3]User{{Name: "Dito", Age: 22}, {Name: "Eddy", Age: 30}, {Name: "Imam", Age: 19}}
	// result := db.Create(&users)

	// fmt.Println("Rows: ", result.RowsAffected)

	// for _, user := range users {
	// 	fmt.Println(user.ID, user.Name, user.Age)
	// }

	// // Query -->
	// result := User{}
	// response := db.Model(&User{}).First(&result) // Get record pertama yang diurutkan berdasarkan primary key
	// // SELECT * FROM users ORDER BY id LIMIT 1;
	// fmt.Println("Jumlah record :", response.RowsAffected) // mengembalikan jumlah record yang ditemukan
	// fmt.Println("Error :", response.Error)                // mengembalikan error atau nil
	// errors.Is(response.Error, gorm.ErrRecordNotFound)     // check error ErrRecordNotFound
	// fmt.Println(result)

	// db.Model(&User{}).Take(&result) // Get satu record, tanpa pengurutan yang spesifik
	// // SELECT * FROM users LIMIT 1
	// fmt.Println("-------------------------")
	// fmt.Println(result)

	// db.Model(&User{}).Last(&result) // Get record terakhir, diurutkan berdasarkan primary key
	// // SELECT * FROM users ORDER BY id DESC LIMIT 1;
	// fmt.Println("-------------------------")
	// fmt.Println(result)

	// // Joining -->
	// results := []Join{}
	// db.Table("users").Select("users.name, users.age, joins.email").Joins("left join joins on joins.id = users.id").Scan(&results)
	// // SELECT users.name, users.age, emails.email FROM `users` left join emails on emails.id = users.id

	// fmt.Println(results)

	// // QueryRow -->
	// results := []User{}
	// rows, err := db.Table("users").Select("name, age").Rows()
	// // SELECT name, age FROM `users`
	// defer rows.Close()
	// for rows.Next() { // Next akan menyiapkan hasil baris berikutnya untuk dibaca dengan metode Scan.
	// 	db.ScanRows(rows, &results)
	// }

	// fmt.Println(results)

	// // SUM Age -->
	// result := Total{}
	// rows, err := db.Table("users").Select("date(created_at) as date, sum(age) as total").Group("date(created_at)").Rows()
	// // SELECT date(created_at) as date, sum(age) as total FROM public.users GROUP BY date(created_at);
	// defer rows.Close()
	// for rows.Next() { // Next akan menyiapkan hasil baris berikutnya untuk dibaca dengan metode Scan.
	// 	err := rows.Scan(&result.Date, &result.Total)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println(result)
	// }

	// // Custom Query -->
	// var result User
	// db.Raw("SELECT id, name, age FROM users WHERE id = ?", 3).Scan(&result)
	// fmt.Println(result)

	// db.Raw("SELECT id, name, age FROM users WHERE name = ?", "Dito").Scan(&result)
	// fmt.Println(result)

	// var age int
	// db.Raw("SELECT SUM(age) FROM users WHERE id IN ?", []int64{1, 2, 3}).Scan(&age)
	// fmt.Println(age)

	// // Update -->
	// user := User{}
	// db.First(&user)

	// user.Name = "Aditira Jamhuri"
	// user.Age = 18
	// db.Save(&user)
	// // UPDATE users SET name='Aditira Jamhuri', age=18, updated_at = '2022-10-19 21:34:10' WHERE id=1;

	// // Update dengan kondisi
	// db.Model(&User{}).Where("id = ?", 1).Update("name", "Jamhuri")
	// // UPDATE users SET name='Jamhuri', updated_at='2022-10-19 21:34:10' WHERE id=1;

	// db.Model(&user).Updates(User{Name: "Steven", Age: 17})
	// // UPDATE users SET name='Steven', age=17, updated_at = '2022-10-19 21:34:10' WHERE id = 3;

	// db.Model(&user).Updates(map[string]interface{}{"name": "Steven", "age": 17})
	// // UPDATE users SET name='Steven', age=17 updated_at='2022-10-19 21:34:10' WHERE id=3;

	// Delate -->
	db.Where("name = ?", "Imam").Delete(&User{})
	// DELETE from users where id = 4 AND name = "Imam";

	db.Delete(&User{}, 10)
	// DELETE FROM users WHERE id = 10;

	db.Delete(&User{}, "10")
	// DELETE FROM users WHERE id = 10;

	// db.Delete(&users, []int{1, 2, 3})
	// // DELETE FROM users WHERE id IN (1,2,3);
}

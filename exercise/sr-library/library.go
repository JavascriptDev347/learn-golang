//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
	"fmt"
	"time"
)

type LendingDetail struct {
	CheckOutTime time.Time
	CheckInTime  *time.Time
}
type Member struct {
	Name  string
	Books map[string]LendingDetail
}
type Book struct {
	Title  string
	Total  int
	Lended int
}

type Library struct {
	Books   map[string]Book
	Members map[string]Member
}

func checkout(library *Library, title string, memberId string) {
	book, exists := library.Books[title]
	if !exists {
		fmt.Println("Xatolik: Bunday kitob yo'q!")
		return
	}

	if book.Lended >= book.Total {
		fmt.Println("Uzr, bu kitob tugagan.")
		return
	}

	member, exists := library.Members[memberId]
	if !exists {
		fmt.Println("Xatolik: Bunday Member yo'q!")
		return
	}

	book.Lended++
	library.Books[title] = book

	member.Books[title] = LendingDetail{
		CheckOutTime: time.Now(),
		CheckInTime:  nil,
	}

	library.Members[memberId] = member

	fmt.Println("Muvaffaqiyatli:", memberId, "uchun", title, "berildi.")
}

func checkIn(lib *Library, title string, memberId string) {
	member, exists := lib.Members[memberId]
	if !exists {
		fmt.Println("Xatolik: Bunday Member yo'q!")
		return
	}

	detail, hasBook := member.Books[title]
	if !hasBook {
		fmt.Println("Xatolik: Bu Memberda bunaqa kitob yo'q!")
	}

	now := time.Now()
	detail.CheckInTime = &now
	member.Books[title] = detail
	lib.Members[memberId] = member

	book := lib.Books[title]
	book.Lended--
	lib.Books[title] = book

	fmt.Println("Rahmat!", memberId, title, "kitobini qaytardi.")
}

func main() {
	// 1. Kutubxonani ishga tushiramiz (Maplarni 'make' qilish esdan chiqmasin!)
	library := Library{
		Books:   make(map[string]Book),
		Members: make(map[string]Member),
	}

	// 2. Kitoblarni qo'shamiz
	library.Books["Go Programming"] = Book{Title: "Go Programming", Total: 3, Lended: 0}
	library.Books["O'tkan kunlar"] = Book{Title: "O'tkan kunlar", Total: 2, Lended: 0}
	library.Books["Algorithm"] = Book{Title: "Algorithm", Total: 5, Lended: 0}
	library.Books["Clean Code"] = Book{Title: "Clean Code", Total: 1, Lended: 0}

	// 3. A'zolarni qo'shamiz
	// Har bir a'zoning o'z 'Books' mapi bor, uni ham 'make' qilish kerak
	library.Members["Ali"] = Member{Name: "Ali", Books: make(map[string]LendingDetail)}
	library.Members["Vali"] = Member{Name: "Vali", Books: make(map[string]LendingDetail)}
	library.Members["Sardor"] = Member{Name: "Sardor", Books: make(map[string]LendingDetail)}

	fmt.Println("--- KUTUBXONA TIZIMI ---")

	// 4. Jarayonlar
	checkout(&library, "Go Programming", "Ali")
	checkout(&library, "O'tkan kunlar", "Sardor")

	// Biroz vaqt o'tkazamiz (simulyatsiya)
	time.Sleep(1 * time.Second)

	checkIn(&library, "Go Programming", "Ali")

	// Yakuniy natijani ko'rish uchun kichik chop etish
	printStatus(&library)
}

func printStatus(lib *Library) {
	fmt.Println("\n--- Status ---")
	for title, book := range lib.Books {
		fmt.Printf("Kitob: %s | Jami: %d | Berilgan: %d\n", title, book.Total, book.Lended)
	}
}

package controllers

import (
	"bufio"
	"fmt"
	"library_management/models"
	"library_management/services"
	"os"
	"strconv"
)

func StartConsole(library *services.Library) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nLibrary Management System")
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed Books by Member")
		fmt.Println("7. Add Member")
		fmt.Println("0. Exit")
		fmt.Print("Choose an option: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("Enter Book ID: ")
			scanner.Scan()
			bookID, _ := strconv.Atoi(scanner.Text())
			fmt.Print("Enter Book Title: ")
			scanner.Scan()
			title := scanner.Text()
			fmt.Print("Enter Author: ")
			scanner.Scan()
			author := scanner.Text()

			book := models.Book{ID: bookID, Title: title, Author: author, Status: "Available"}
			library.AddBook(book)
			fmt.Println("Book added successfully.")

		case "2":
			fmt.Print("Enter Book ID to remove: ")
			scanner.Scan()
			bookID, _ := strconv.Atoi(scanner.Text())
			err := library.RemoveBook(bookID)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Book removed successfully.")
			}

		case "3":
			fmt.Print("Enter Book ID to borrow: ")
			scanner.Scan()
			bookID, _ := strconv.Atoi(scanner.Text())
			fmt.Print("Enter Member ID: ")
			scanner.Scan()
			memberID, _ := strconv.Atoi(scanner.Text())

			err := library.BorrowBook(bookID, memberID)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Book borrowed successfully.")
			}

		case "4":
			fmt.Print("Enter Book ID to return: ")
			scanner.Scan()
			bookID, _ := strconv.Atoi(scanner.Text())
			fmt.Print("Enter Member ID: ")
			scanner.Scan()
			memberID, _ := strconv.Atoi(scanner.Text())

			err := library.ReturnBook(bookID, memberID)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Book returned successfully.")
			}

		case "5":
			fmt.Println("Available Books:")
			for _, b := range library.ListAvailableBooks() {
				fmt.Println(b.ID, "-", b.Title, "by", b.Author)
			}

		case "6":
			fmt.Print("Enter Member ID: ")
			scanner.Scan()
			memberID, _ := strconv.Atoi(scanner.Text())
			books, err := library.ListBorrowedBooks(memberID)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Borrowed Books:")
				for _, b := range books {
					fmt.Println(b.ID, "-", b.Title, "by", b.Author)
				}
			}

		case "7":
			fmt.Print("Enter Member ID: ")
			scanner.Scan()
			memberID, _ := strconv.Atoi(scanner.Text())
			fmt.Print("Enter Member Name: ")
			scanner.Scan()
			name := scanner.Text()
			library.Members[memberID] = models.Member{ID: memberID, Name: name}
			fmt.Println("Member added successfully.")

		case "0":
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

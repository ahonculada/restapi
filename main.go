package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

// Book Struct (Model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Case Struct (Model)
type Case struct {
	Results    *Results `json:"results"`
	Error      bool     `json:"error"`
	Error_Test string   `json:"error_text"`
	Run_Time   string   `json:"run_time"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type Results struct {
	ID                    *ID                    `json:"id"`
	Full_Name             *Full_Name             `json:"full_name"`
	County                *County                `json:"county"`
	Address               *Address               `json:"addr1"`
	Phone_Home            *Phone_Home            `json:"phone_home"`
	Phone_Business        *Phone_Business        `json:"phone_business"`
	Phone_Fax             *Phone_Fax             `json:"phone_fax"`
	Phone_Mobile          *Phone_Mobile          `json:"phone_mobile"`
	Phone_Other           *Phone_Other           `json:"phone_other"`
	Legal_Problem_Code    *Legal_Problem_Code    `json:"legal_problem_code"`
	Unique_ID             *Unique_ID             `json:"unique_id"`
	DOB                   *DOB                   `json:"dob"`
	Identification_Number *Identification_Number `json:"identification_number"`
	Email                 *Email                 `json:"email"`
	Current_Disposition   *Current_Disposition   `json:"current_disposition"`
	Date_Open             *Date_Open             `json:"date_open"`
	Close_Date            *Close_Date            `json:"close_date"`
	Close_Reason          *Close_Reason          `json:"close_reason"`
	Disposition_Status    *Disposition_Status    `json:"disposition_status"`
	Disposition_Date      *Disposition_Date      `json:"disposition_date"`
	Primary_Advocate      *Primary_Advocate      `json:"primary_advocate"`
	URL                   string                 `json:"url"`
}

type ID struct {
	Raw_Value  string `json:"raw_value"`
	Text_Value string `json:"text_value"`
}

type Full_Name struct {
	Raw_Value  string `json:"raw_value"`
	Text_Value string `json:"text_value"`
}

type County struct {
	Raw_Value  string `json:"raw_value"`
	Text_Value string `json:"text_value"`
}

type Address struct {
	Raw_Value  string `json:"raw_value"`
	Text_Value string `json:"text_value"`
}

type Phone_Home struct {
	Raw_Value  string `json:"raw_value"`
	Text_Value string `json:"text_value"`
}

type Phone_Business struct {
	Raw_Value  string `json:"raw_value"`
	Text_Value string `json:"text_value"`
}

type Phone_Fax struct {
	Raw_Value  string `json:"raw_value"`
	Text_Value string `json:"text_value"`
}

type Phone_Mobile struct {
	Raw_Value  string `json:"raw_value"`
	Text_Value string `json:"text_value"`
}

type Phone_Other struct {
	Raw_Value  string `json:"raw_value"`
	Text_Value string `json:"text_value"`
}

type Legal_Problem_Code struct {
	Raw_Value  string `json:"raw_value"`
	Text_Value string `json:"text_value"`
}

type Unique_ID struct {
	Raw_Value  string `json:"raw_value"`
	Text_Value string `json:"text_value"`
}

type DOB struct {
	Raw_Value  string `json:"raw_value"`
	Text_Value string `json:"text_value"`
}

type Identification_Number struct {
	Raw_Value  string `json:"raw_value"`
	Text_Value string `json:"text_value"`
}

type Email struct {
	Raw_Value  string `json:"raw_value"`
	Text_Value string `json:"text_value"`
}

type Current_Disposition struct {
	Raw_Value  string `json:"raw_value"`
	Text_Value string `json:"text_value"`
}

type Date_Open struct {
	Raw_Value  string `json:"raw_value"`
	Text_Value string `json:"text_value"`
}

type Close_Date struct {
	Raw_Value  string `json:"raw_value"`
	Text_Value string `json:"text_value"`
}

type Close_Reason struct {
	Raw_Value  string `json:"raw_value"`
	Text_Value string `json:"text_value"`
}

type Disposition_Status struct {
	Raw_Value  string `json:"raw_value"`
	Text_Value string `json:"text_value"`
}

type Disposition_Date struct {
	Raw_Value  string `json:"raw_value"`
	Text_Value string `json:"text_value"`
}

type Primary_Advocate struct {
	Raw_Value  string `json:"raw_value"`
	Text_Value string `json:"text_value"`
}

// Init books var as a slice Book struct
var books []Book

// Init cases var as a slice Case struct
var cases []Case

// Get all Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get all Cases
func getCases(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control_Allow-Origin", "*")
	json.NewEncoder(w).Encode(cases)
}

// Get single Book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // get params
	// loop through books and find with id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// Get a single Case
func getCase(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // get params
	// loop through cases and match dob, identification_number, full_name
	// for _, item := range cases {}
}

// Create a New Book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book

	_ = json.NewDecoder(r.Body).Decode(&book)

	book.ID = strconv.Itoa(rand.Intn(1000000)) // Mock ID - not safe (could generate same id)
	books = append(books, book)
	json.NewEncoder(w).Encode(book)

}

// Update Book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)

			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)

}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)

}

func main() {
	// Init Router
	r := mux.NewRouter()

	// Mock Data
	cases = append(cases,
		Case{
			Results: &Results{ID: &ID{Raw_Value: "5", Text_Value: "5"},
				Full_Name:             &Full_Name{Raw_Value: "TEST TEST", Text_Value: "TEST TEST"},
				County:                &County{Raw_Value: "", Text_Value: "N/A"},
				Address:               &Address{Raw_Value: "667 Broadway", Text_Value: "667 Broadway"},
				Phone_Home:            &Phone_Home{Raw_Value: "123-456-8999", Text_Value: "(123) 456-8999"},
				Phone_Business:        &Phone_Business{Raw_Value: "", Text_Value: "N/A"},
				Phone_Fax:             &Phone_Fax{Raw_Value: "", Text_Value: "N/A"},
				Phone_Mobile:          &Phone_Mobile{Raw_Value: "", Text_Value: "N/A"},
				Phone_Other:           &Phone_Other{Raw_Value: "", Text_Value: "N/A"},
				Legal_Problem_Code:    &Legal_Problem_Code{Raw_Value: "8964", Text_Value: "01 Bankruptcy/Debtor Relief"},
				Unique_ID:             &Unique_ID{Raw_Value: "52ed3287-6984-476c-8092-67a6c0a80a22", Text_Value: "52ed3287-6984-476c-8092-67a6c0a80a22"},
				DOB:                   &DOB{Raw_Value: "1970-05-10", Text_Value: "05/10/1970"},
				Identification_Number: &Identification_Number{Raw_Value: "14-0000005", Text_Value: "14-0000005"},
				Email:                 &Email{Raw_Value: "jhogue@legalserver.org", Text_Value: "jhogue@legalserver.org"},
				Current_Disposition:   &Current_Disposition{Raw_Value: "Case: Open", Text_Value: "Case: Open"},
				Date_Open:             &Date_Open{Raw_Value: "2014-02-01", Text_Value: "02/01/2014"},
				Close_Date:            &Close_Date{Raw_Value: "", Text_Value: "N/A"},
				Close_Reason:          &Close_Reason{Raw_Value: "", Text_Value: "N/A"},
				Disposition_Status:    &Disposition_Status{Raw_Value: "N/A", Text_Value: "N/A"},
				Primary_Advocate:      &Primary_Advocate{Raw_Value: "124", Text_Value: "Meagen Peden Agnew"},
				URL:                   "/matter/dymaic-profile/view/5"},
			Error: false, Error_Test: "", Run_Time: "0.16840791702271s"})

	// Route Handlers / Endpoints
	// r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/matter/api/basic_case_info", getCases).Methods("GET")
	//r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	//r.HandleFunc("/matter/api/basic_case_info/{id}", getCase).Methods("GET")
	//r.HandleFunc("/api/books", createBook).Methods("POST")
	//r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	//r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))

}

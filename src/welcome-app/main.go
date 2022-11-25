package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
	"encoding/json"
)

// Create a struct that holds information to be displayed in our HTML file
type Welcome struct {
	Name string
	Time string
}

// type JsonResponse struct{
// 	Value1 string `json:"key1"`
// 	Value2 string `json:"key2"`
// 	JsonNested JsonNested `json:"jsonNested"`
// }
// type JsonNested struct{
// 	NestedValue1 string `json:"nestedkey1"`
// 	NestedValue2 string `json:"nestedkey2"`
// }

type JsonContactInfo struct{
	Value1 string `json:"key1"`
	Value2 string `json:"key2"`
	Value3 string `json:"key3"`
	JsonNested JsonNested `json:"JsonNested"`
	Value4 string `json:"key4"`
	JsonNestedContactInfo JsonNestedContactInfo `json: "JsonNestedContactInfo"`
}
type JsonNested struct{
	NestedValue1 string `json:"nestedkey1"`
	NestedValue2 string `json:"nestedkey2"`
}

type JsonNestedContactInfo struct{
	NestedValue3 string `json:"nestedkey3"`
	NestedValue4 string `json:"nestedkey4"`
}


// Go application entrypoint
func main() {

	welcome := Welcome{"Anonymous", time.Now().Format(time.Stamp)}

	//We tell Go exactly where we can find our html file. We ask Go to parse the html file (Notice
	// the relative path). We wrap it in a call to template.Must() which handles any errors and halts if there are fatal errors

	templates := template.Must(template.ParseFiles("templates/welcome-template.html"))
	
	// nested := JsonNested{
	// 	NestedValue1: "first nested val",
	// 	NestedValue2 : "second nested val",
	// 	JsonNested: nested,
	// }

	// jsonResp := JsonResponse{
	// 	Value1: "some Data",
	// 	Value2: "other data",
	// }

	nested := JsonNested{
		NestedValue1: "1234 Windbrook Lane",
		NestedValue2 : "Sunnydale, Georgia",
	}

	NestedContactInfo := JsonNestedContactInfo{
		NestedValue3: "Email: windbrook123@gmail.com",
		NestedValue4: "Phone: 706-123-4567",
	}

	jsonContactInfo := JsonContactInfo{
		Value1: "FirstName: John",
		Value2: "LastName: Smith",
		Value3: "Address:",
		JsonNested: nested,
		Value4: "Contact Information:",
		JsonNestedContactInfo: NestedContactInfo,
	}
	//Our HTML comes with CSS that go needs to provide when we run the app. Here we tell go to create
	// a handle that looks in the static directory, go then uses the "/static/" as a url that our
	//html can refer to when looking for our css and other files.
	//fmt.Fprint("%+v\n\n\n\n", jsonResp)
	http.Handle("/static/", //final url can be anything
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static")))) //Go looks in the relative "static" directory first using http.FileServer(), then matches it to a

	//This method takes in the URL path "/" and a function that takes in a response writer, and a http request.
	// **** THIS IS THE MAIN PATH /
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		//Takes the name from the URL query e.g ?name=Martin, will set welcome.Name = Martin.
		if name := r.FormValue("name"); name != "" {
			welcome.Name = name
		}
		//If errors show an internal server error message
		//I also pass the welcome struct to the welcome-template.html file.
		if err := templates.ExecuteTemplate(w, "welcome-template.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})


	// http.HandleFunc("/jsonResponse", func(w http.ResponseWriter, r *http.Request){
	// 	json.NewEncoder(w).Encode(jsonResp)
	// })

	http.HandleFunc("/jsonContactInfo", func(w http.ResponseWriter, r *http.Request){
		json.NewEncoder(w).Encode(jsonContactInfo)
	})

	fmt.Println("Listening")
	fmt.Println(http.ListenAndServe(":8080", nil))
}

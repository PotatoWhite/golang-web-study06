package main

import (
	"html/template"
	"os"
)

type User struct {
	Name  string
	Email string
	Age   int
}

func (u User) IsOld() bool {
	return u.Age > 30
}

func main() {
	tmpl, err := template.New("Tmpl1").ParseFiles("templates/temp01.tmpl", "templates/temp02.tmpl")
	if err != nil {
		panic(err)
	}

	user01 := User{Name: "potato", Email: "potato@example.com", Age: 20}
	user02 := User{Name: "carrot", Email: "carrot@example.com", Age: 40}

	users := []User{user01, user02}

	tmpl.ExecuteTemplate(os.Stdout, "temp02.tmpl", user01)
	tmpl.ExecuteTemplate(os.Stdout, "temp02.tmpl", user02)
	tmpl.ExecuteTemplate(os.Stdout, "temp02.tmpl", users)

}

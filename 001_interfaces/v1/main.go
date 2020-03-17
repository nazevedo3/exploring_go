package main

import "fmt"

// Create my two "databases."
type mongo map[int]person
type postg map[int]person

//Define my person struct
type person struct {
	first string
}

//Get method pulls a person from the database.
func (m mongo) get(i int) person {
	return m[i]
}

//Post method stores a person in the database.
func (m mongo) post(i int, p person) {
	m[i] = p
}

//Get method pulls a person from the database.
func (pg postg) get(i int) person {
	return pg[i]
}

//Post method stores a person in the database.
func (pg postg) post(i int, p person) {
	pg[i] = p
}

//Accessor allows you to get or post to either database type (Mongo or Postress)
type accessor interface {
	get(i int) person
	post(i int, p person)
}

//Put stores info into a database through the accessor interface.
func put(a accessor, i int, p person) {
	a.post(i, p)
}

//Retrieves pulls info from a database through the accessor interface.
func retrieve(a accessor, i int) person {
	return a.get(i)
}

func main() {

	mg := mongo{}
	postg := postg{}

	p1 := person{"Nathan"}
	p2 := person{"Quica"}

	put(mg, 1, p1)
	put(postg, 1, p2)

	fmt.Println("Stored in Mongo", retrieve(mg, 1))
	fmt.Println("Stored in Postgress", retrieve(postg, 1))

}

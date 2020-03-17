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

type personService struct {
	a accessor
}

func (ps personService) get(i int) (person, error) {
	p := ps.a.get(i)

	if p.first == "" {
		return person{}, fmt.Errorf("No person with n of %d", i)
	}

	return p, nil
}

func (ps personService) put(i int, p person) {
	ps.a.post(i, p)
}

func main() {

	// mg := mongo{}
	postg := postg{}

	p1 := person{"Nathan"}
	p2 := person{"Quica"}

	ps := personService{
		a: postg,
	}

	ps.a.post(1, p1)
	ps.a.post(2, p2)

	fmt.Println(ps.get(1))
	fmt.Println(ps.get(2))
	fmt.Println(ps.get(3)) //error

}

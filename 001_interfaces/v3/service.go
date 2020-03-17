package architecture

import "fmt"

// Person is how the architecture stores a person
type Person struct {
	First string
}

// Accessor allows you to get or post to either database type (Mongo or Postress)
// When retriving a person, if they do not exist, return the zero value
type Accessor interface {
	Get(i int) Person
	Post(i int, p Person)
}

type PersonService struct {
	a Accessor
}

func NewPersonService(a Accessor) PersonService {
	return PersonService{
		a: a,
	}
}

func (ps PersonService) Get(i int) (Person, error) {
	p := ps.a.Get(i)

	if p.First == "" {
		return Person{}, fmt.Errorf("No person with n of %d", i)
	}

	return p, nil
}

func (ps PersonService) Put(i int, p Person) {
	ps.a.Post(i, p)
}

//Put stores info into a database through the accessor interface.
func Put(a Accessor, i int, p Person) {
	a.Post(i, p)
}

//Retrieves pulls info from a database through the accessor interface.
func Retrieve(a Accessor, i int) Person {
	return a.Get(i)
}

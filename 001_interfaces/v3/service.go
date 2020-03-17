package architecture

import "fmt"

//Person struct
type Person struct {
	First string
}

//Accessor allows you to get or post to either database type (Mongo or Postress)
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

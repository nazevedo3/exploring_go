package postgres

import architecture "github.com/nazevedo3/exploring_go"

type Db map[int]architecture.Person

//Get method pulls a person from the database.
func (pg Db) Get(i int) architecture.Person {
	return pg[i]
}

//Post method stores a person in the database.
func (pg Db) Post(i int, p architecture.Person) {
	pg[i] = p
}

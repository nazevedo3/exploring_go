package mongo

import architecture "github.com/nazevedo3/exploring_go"

type Db map[int]architecture.Person

//Get method pulls a person from the database.
func (m Db) Get(i int) architecture.Person {
	return m[i]
}

//Post method stores a person in the database.
func (m Db) Post(i int, p architecture.Person) {
	m[i] = p
}

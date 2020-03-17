package architecture

import (
	"fmt"
	"testing"
)

type Db map[int]Person

//Get method pulls a person from the database.
func (m Db) Get(i int) Person {
	return m[i]
}

//Post method stores a person in the database.
func (m Db) Post(i int, p Person) {
	m[i] = p
}
func TestPut(t *testing.T) {
	mdb := Db{}
	p := Person{"James"}

	Put(mdb, 1, p)

	got := mdb.Get(1)

	if got != p {
		t.Fatalf("Want %v, got %v", p, got)
	}

}

func ExamplePut() {
	mdb := Db{}
	p := Person{"James"}

	Put(mdb, 1, p)
	got := mdb.Get(1)
	fmt.Println(got)
	// Output: {James}

}

package main

import (
	"fmt"

	architecture "github.com/nazevedo3/exploring_go"
	"github.com/nazevedo3/exploring_go/storage/postgres"
)

func main() {

	//dbm := mongo.Db{}
	dgp := postgres.Db{}

	p1 := architecture.Person{"Nathan"}
	p2 := architecture.Person{"Quica"}

	ps := architecture.NewPersonService(dgp)

	ps.Put(1, p1)
	ps.Put(2, p2)

	fmt.Println(ps.Get(1))
	fmt.Println(ps.Get(2))
	fmt.Println(ps.Get(3)) //error

}

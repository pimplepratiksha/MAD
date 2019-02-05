package main
import (
	"fmt"
	"os"
	dbrepo "../assignment1/dbrepository"
	mongoutils "../assignment1/utils"
)

func main() {
	//pass mongohost through the environment
	mongoSession, _ := mongoutils.RegisterMongoSession(os.Getenv("MONGO_HOST"))
	dbname := "restaurant"
	repoaccess := dbrepo.NewMongoRepository(mongoSession, dbname)
	fmt.Println(repoaccess)
	
	repoaccess.Insert("newset")
	
}

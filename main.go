package main
import (
	"fmt"
	"os"
	"strings"
	dbrepo "../assignment1/dbrepository"
	mongoutils "../assignment1/utils"
	domain "../assignment1/domain"
)

func main() {
	//pass mongohost through the environment
	mongoSession, _ := mongoutils.RegisterMongoSession(os.Getenv("MONGO_HOST"))
	dbname := "restaurant"
	repoaccess := dbrepo.NewMongoRepository(mongoSession, dbname)
	fmt.Println(repoaccess)
	
	var ip string
	var final []*domain.Restaurant 
	var err error
	if len(os.Args)>1{
		ip=os.Args[1]
		arr:=strings.Split(ip,"=")
		switch(arr[0]){
			case "--type_of_food":
				final,err=repoaccess.FindByTypeOfFood(arr[1])
			case "--postcode":
				final,err=repoaccess.FindByTypeOfPostCode(arr[1])
			default:
				fmt.Println("invalid argument")
				return 
		}
		if err!=nil{
			fmt.Println(err)
			return 
		}
	
		for _,z:=range final {
			fmt.Println(z)
		}
	}
}

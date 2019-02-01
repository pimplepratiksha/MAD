package dbrepository

import (
	//"github.com/priteshgudge/mongorestaurantsample/domain"
	domain "../domain"
	mgo "gopkg.in/mgo.v2"
	bson "gopkg.in/mgo.v2/bson"
	"encoding/json"
	"bufio"
	"os"
)

//MongoRepository mongodb repo
type MongoRepository struct {
	mongoSession *mgo.Session
	db           string
}

var collectionName = "newrestaurant"

//*****************************************************************************************************

//NewMongoRepository create new repository
func NewMongoRepository(mongoSession *mgo.Session, db string) *MongoRepository {
	return &MongoRepository{
		mongoSession: mongoSession,
		db:           db,
	}
}

func (r *MongoRepository) Insert(filename string) (int,error){
	fname,err:=os.Open(filename)
	if err!=nil{
		return 0,err
	}
	defer fname.Close()
	fp:=bufio.NewScanner(fname)
	var final=&domain.Restaurant{}
	rcnt:=0
	for fp.Scan(){
			rcnt+=1
		json.Unmarshal([]byte(fp.Text()),final)
		final.DBID=domain.NewID()
		_,err:=r.Store(final)
		if	err!=nil{
			return rcnt,err
		}
	}
	
	return rcnt,nil
}

//Find a Restaurant(reader)
func (r *MongoRepository) Get(id domain.ID) (*domain.Restaurant, error) {
	result := domain.Restaurant{}
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	err := coll.Find(bson.M{"_id": id}).One(&result)
	switch err {
		case nil:
			return &result, nil
		case mgo.ErrNotFound:
			return nil, domain.ErrNotFound
		default:
			return nil, err
	}
}

//get all restaurants(reader)
func (r *MongoRepository) GetAll() ([]*domain.Restaurant, error) {
	result := []*domain.Restaurant{}
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	err := coll.Find(bson.M{}).All(&result)
	switch err {
		case nil:
			return result, nil
		case mgo.ErrNotFound:
			return nil, domain.ErrNotFound
		default:
			return nil, err
	}
}

//Find a Restaurant By Name(reader)
func (r *MongoRepository) FindByName(name string) ([]*domain.Restaurant, error) {
	result := []*domain.Restaurant{}
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	err := coll.Find(bson.M{"name":bson.RegEx{name,"i"}}).All(&result) 	
	switch err {
		case nil:
			return result, nil
		case mgo.ErrNotFound:
			return nil, domain.ErrNotFound
		default:
			return nil, err
	}
}

//Store a Restaurant record(writer)
func (r *MongoRepository) Store(b *domain.Restaurant) (domain.ID, error) {
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	if domain.ID(0) == b.DBID {
		b.DBID = domain.NewID()
	}
	_, err := coll.UpsertId(b.DBID, b)

	if err != nil {
		return domain.ID(0), err
	}
	return b.DBID, nil
}

//Delete a Restaurant record(writer)
func (r *MongoRepository) Delete(id domain.ID)(error) {
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	err := coll.Remove(bson.M{"_id": id})
	switch err {
		case nil:
			return nil
		case mgo.ErrNotFound:
			return domain.ErrNotFound
		default:
			return err
	}
}

//Find a Restaurant By Type Of Food(filter)
func (r *MongoRepository) FindByTypeOfFood(foodtype string) ([]*domain.Restaurant, error) {
	result := []*domain.Restaurant{}
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	err := coll.Find(bson.M{"typeOfFood": foodtype}).All(&result) 	
	switch err {
		case nil:
			return result, nil
		case mgo.ErrNotFound:
			return nil, domain.ErrNotFound
		default:
			return nil, err
	}
}

//Find a Restaurant By Type Of Post Code(filter)
func (r *MongoRepository) FindByTypeOfPostCode(postcode string) ([]*domain.Restaurant, error) {
	result := []*domain.Restaurant{}
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	err := coll.Find(bson.M{"postcode": postcode}).All(&result) 	
	
	switch err {
		case nil:
			return result, nil
		case mgo.ErrNotFound:
			return nil, domain.ErrNotFound
		default:
			return nil, err
	}
}

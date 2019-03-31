package repoimpl

import (
	"context"
	models "go-jwt/model"
	repo "go-jwt/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepoImpl struct {
	Db *mongo.Database
}

func NewUserRepo(db *mongo.Database) repo.UserRepo {
	return &UserRepoImpl{
		Db: db,
	}
}

func (mongo *UserRepoImpl) FindUserByEmail(email string) (models.User, error) {
	user := models.User{}

	result := mongo.Db.Collection("users").
						FindOne(context.Background(), 
								bson.M{"email" :email})	

	err := result.Decode(&user)	

	if user == (models.User{}) {
		return user, models.ERR_USER_NOT_FOUND
	}

	if err != nil {
		return user, err 
	}

	return user, nil
}

func (mongo *UserRepoImpl) Insert(user models.User) error {
	bbytes, _ := bson.Marshal(user)
	
	_, err := mongo.Db.Collection("users").InsertOne(context.Background(), bbytes)
	if err != nil {
		return err
	}

	return nil
}

func (mongo *UserRepoImpl) CheckLoginInfo(email, password string) (models.User, error) {
	user := models.User{}

	result := mongo.Db.Collection("users").
						FindOne(context.Background(), 
								bson.M{
									"email" :email,
									"password": password,
								})	

	err := result.Decode(&user)	
	if err != nil {
		return user, err 
	}

	return user, nil
}












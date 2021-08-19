package controller

import (
	"Project_Eular/models"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

const connectionString = "mongodb://localhost:27017"
const dbName = "Eular"
const userCollName = "user"
const questionCollName = "question"

var SECRET_KEY = []byte("MY_SECRET_KEY")

type User struct {
	Token     string
	UserEmail string
	UserName  string
}

var userCollection *mongo.Collection
var questionCollection *mongo.Collection

func init() {
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	userCollection = client.Database(dbName).Collection(userCollName)
	questionCollection = client.Database(dbName).Collection(questionCollName)
}

func GetLandingPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	var existingUser models.User
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	json.Unmarshal(reqBody, &user)
	user.Password = getHash([]byte(user.Password))
	var filter = bson.M{"email": user.Email}
	existingUserError := userCollection.FindOne(context.Background(), filter).Decode(&existingUser)
	if existingUserError == nil {
		fmt.Println("User Already Exist")
		http.Error(w, existingUserError.Error(), 500)
	} else if len(existingUser.Email) > 0 {
		_, _ = w.Write([]byte(`{"response":"User Already Exist!"}`))
	} else {
		result, err := userCollection.InsertOne(context.Background(), user)
		if err != nil {
			panic(err)
		}
		json.NewEncoder(w).Encode(result)
	}
}

func getHash(password []byte) string {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	return string(hash)
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	var dbUser models.User
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	_ = json.Unmarshal(reqBody, &user)
	filter := bson.M{"email": user.Email}
	dbError := userCollection.FindOne(context.Background(), filter).Decode(&dbUser)
	if dbError != nil {
		_, _ = w.Write([]byte(`{"message":"` + err.Error() + `"}`))
		return
	}
	userPassword := []byte(user.Password)
	dbUserPassword := []byte(dbUser.Password)

	passwordError := bcrypt.CompareHashAndPassword(dbUserPassword, userPassword)

	if passwordError != nil {
		log.Println(passwordError)
		_, _ = w.Write([]byte(`{"response":"Wrong Password!"}`))
		return
	}

	jwtToken, err := GenerateJWT()
	if err != nil {
		_, _ = w.Write([]byte(`{"response":"Wrong Password!"}`))
		return
	}
	response := User{
		Token:     jwtToken,
		UserEmail: user.Email,
		UserName:  user.Username,
	}
	//_, _ = w.Write([]byte(response))
	var responseBytes []byte
	var err2 error

	responseBytes, err2 = json.Marshal(response)
	if err2 != nil {
		print(err2)
		return
	}
	_, _ = w.Write([]byte(responseBytes))
	//json.NewEncoder(w).Encode(string(responseBytes))
}

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString(SECRET_KEY)
	if err != nil {
		log.Println("Error in JWT token generation")
		return "", err
	}
	return tokenString, nil
}

func InsertQuestions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-COntrol-Allow-Origin", "Content-Type")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	var question models.Question
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	json.Unmarshal(reqBody, &question)
	insertQuestionResponse, questionInsertError := questionCollection.InsertOne(context.Background(), question)
	if questionInsertError != nil {
		http.Error(w, questionInsertError.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(insertQuestionResponse)
}

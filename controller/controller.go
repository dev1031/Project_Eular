package controller

import (
	"Project_Eular/models"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

const connectionString = "mongodb://localhost:27017"
const dbName = "Eular"
const userCollName = "user"
const questionCollName = "question"
const commentCollName = "comment"

var SECRET_KEY = []byte("MY_SECRET_KEY")

type User struct {
	Token     string
	UserEmail string
	UserName  string
}

var userCollection *mongo.Collection
var questionCollection *mongo.Collection
var commentCollection *mongo.Collection

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
	commentCollection = client.Database(dbName).Collection(commentCollName)
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
	w.Header().Set("Access-Control-Allow-Origin", "Content-Type")
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

func Comments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Origin", "Content-Type")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	var comment models.Comments
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(errors.New("something went wrong"))
		http.Error(w, err.Error(), 500)
		return
	}
	json.Unmarshal(reqBody, &comment)
	result, resultError := commentCollection.InsertOne(context.Background(), comment)
	if resultError != nil {
		http.Error(w, resultError.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(result)
}

func Profile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Origin", "Content-Type")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	var requestedUser models.User
	id, dberr := primitive.ObjectIDFromHex(params["id"])
	if dberr != nil {
		fmt.Println(dberr)
	}
	var filter = bson.M{"_id": id}
	error := userCollection.FindOne(context.Background(), filter).Decode(&requestedUser)
	if error != nil {
		fmt.Println(error)
	}
	json.NewEncoder(w).Encode(requestedUser)
}

func ImageUpload(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Origin", "Content-Type")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	fmt.Println("File Upload Endpoint Hit")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	tempFile, err := ioutil.TempFile("uploads", "img-*.png")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Uploaded File: %+v\n", tempFile.Name())

	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
	//finding the user and uploading the profile image field for databse
	var params = mux.Vars(r)
	//var requestedUser models.User
	id, dberr := primitive.ObjectIDFromHex(params["id"])
	if dberr != nil {
		fmt.Println(dberr)
	}
	var filter = bson.M{"_id": id}
	_, error := userCollection.UpdateOne(context.Background(), filter, bson.D{{"$set", bson.D{{"image", tempFile.Name()}}}})
	// return that we have successfully uploaded our file!
	if error != nil {
		fmt.Println(error)
	}
	fmt.Fprintf(w, "Successfully Uploaded File\n")
}

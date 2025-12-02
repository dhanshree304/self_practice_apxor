package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"backend/db"
)

type User struct {
	ID    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name  string             `json:"name"`
	Email string             `json:"email"`
}


//.Header() → Gets the header map of the response.
//Headers describe what type of data we are sending.
//ctx → A context object.
//Used to control timeouts and cancel operations.
//cancel → A function used to stop/cleanup the context.
//defer → Tells Go to run this line at the end of the function.
//cancel() → Calls the cancel function created above.
//It frees resources and stops the context.
//result → A variable that will store the output from MongoDB
//(usually the inserted document ID).


// CREATE user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	result, err := db.Collection.InsertOne(ctx, user)
//ctx → The context with timeout (max 10 seconds).
//If the DB insert takes longer than 10 seconds → it cancels.
	
	if err != nil {

		//“Did InsertOne fail?”
		//If yes → handle the error.

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(result)
	//json → Go’s JSON package.
	//.NewEncoder(...) → Creates a JSON encoder.
	//(w) → Sends the output to the HTTP response.
	//.Encode(result) → Converts result into JSON and writes it to the client.
}




	

// READ ALL user or get all user
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := db.Collection.Find(ctx, bson.M{})
//bson.M{} →
//A MongoDB filter.
//bson.M → BSON map (key-value format)
//{} → empty filter → means fetch all documents.

//Fetches all documents from MongoDB.
//Checks for errors.
//Loops through each document.
//Converts each document into a User struct.
//Stores all users in a users slice.
//Closes the cursor after reading.
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var users []User //users is slice hold multiple object [{},{},{}]
	//users → A slice to hold multiple User objects.
	//[]User → Slice of User type. ..holds objects with User keys
//cursor.Next(ctx) →
//Moves the cursor to the next document.
//Returns true if there is another document.

	for cursor.Next(ctx) {
		var user User //user contain one doc {} from DB
		cursor.Decode(&user)//user is one doc in that slice from mongoDB
		users = append(users, user)//users slice me user objects append kr rhe hai
	}
//cursor.Decode →
//Converts the MongoDB document into a Go struct.
//&user →
//Pass the address of user → so the data is filled in that struct.
json.NewEncoder(w).Encode(users)
//Encode(users):-Converts users slice into JSON and writes it to the client

}






// UPDATE  user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.URL.Query().Get("id")
	objID, _ := primitive.ObjectIDFromHex(id)
	//Converts a hex string into a MongoDB ObjectID.

	var user User
	json.NewDecoder(r.Body).Decode(&user)

// ctx:=
// Context to control the lifetime of the DB operation.
// cancelFunction to stop the context.
// Short declaration.
// context.WithTimeout
// Create a context that expires after X time.

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
//bson.M
//MongoDB key-value map (like JSON object).


	update := bson.M{//single object in DB
		"$set": bson.M{
			"name":  user.Name,
			"email": user.Email,
		},
	}

	result, err := db.Collection.UpdateOne(ctx, bson.M{"_id": objID}, update)//The update object we created.
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}





//  DELETE user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.URL.Query().Get("id")
	objID, _ := primitive.ObjectIDFromHex(id)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
 
 	result, err := db.Collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}


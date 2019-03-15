package main

import(
  "fmt"
  "net/http"
  "time"
  "log"
  "context"
  "encoding/json"

  "github.com/gorilla/mux"
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
  "go.mongodb.org/mongo-driver/bson/primitive"
)
//func getClientDB(ctx context.Context) (*mongo.Database, error){
func getClientDB() (*mongo.Client){

  client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://test:test@imaginex-developer-shard-00-00-ue0if.mongodb.net:27017/dbname?ssl=true&authmechanism=SCRAM-SHA-1&authSource=admin"))
  if err != nil {
    log.Fatal(err)
    return nil
  }
  ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
  defer cancel()
  err = client.Connect(ctx)
  // db := client.Database("db_notify_track")
  // if err != nil {
  //   log.Fatal(err)
  //   return nil
  // }

  fmt.Println("Connected to Mongo")
  return client
}

var collection = getClientDB().Database("db_notify_track").Collection("col_notify_dte")
//Index funcion de inicio de API
func Index(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "Hola, este es el inicio")
}

func Response(w http.ResponseWriter, status int, results ErrorMsg){
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(200)
  json.NewEncoder(w).Encode(results)
}

func ErrorShow(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  var result ErrorMsg

  errorID := params["id"]

  fmt.Println(errorID)

  id, err := primitive.ObjectIDFromHex(errorID)
  if err != nil{
    log.Fatal(err)
  }

  filter := bson.D{{"_id", id}}

  err = collection.FindOne(context.TODO(), filter).Decode(&result)
  if err != nil{
    log.Fatal(err)
  }
  Response(w, 200, result)
  fmt.Printf("Found a single error: %v\n", result)

}

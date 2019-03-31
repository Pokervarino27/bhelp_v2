package main

import(
  "fmt"
  "net/http"
  "time"
  "log"
  "context"
  "encoding/json"
  "strconv"

  "github.com/gorilla/mux"
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
  //"go.mongodb.org/mongo-driver/bson/primitive"
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
  err = client.Ping(context.TODO(), nil)
  fmt.Println("Connected to Mongo")
  return client
}

var collection = getClientDB().Database("db_notify_track").Collection("col_notify_dte")
//Index funcion de inicio de API
func Index(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "Hola, este es el inicio")
}

func Response(w http.ResponseWriter, status int, result ErrorMsg){
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(200)
  json.NewEncoder(w).Encode(result)
}

func ErrorShow(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  var result ErrorMsg

  cpn := params["cpnid"]

  i, err := strconv.Atoi(cpn)
  if err != nil {
    fmt.Println(err)
  }
  // id, err := primitive.ObjectIDFromHex(errorID)
  // if err != nil{
  //   log.Fatal(err)
  // }

  filter := bson.D{{"cpnid", i}}
  err = collection.FindOne(context.TODO(), filter).Decode(&result)
  if err != nil{
    log.Fatal(err)
  }
  Response(w, 200, result)
  fmt.Printf("Found a single error: %v\n", result)
}

func ErrorAdd(w http.ResponseWriter, r *http.Request){
  decoder := json.NewDecoder(r.Body)

  var errorData ErrorMsg
  err := decoder.Decode(&errorData)
  if err != nil {
    panic(err)
  }
  defer r.Body.Close()

  insertResult, err := collection.InsertOne(context.TODO(), errorData)
  if err != nil{
    log.Fatal(err)
    w.WriteHeader(500)
    return
  }

  fmt.Println("Error Insert: ", insertResult.InsertedID)
  Response(w, 200, errorData)
}

func ErrorUpdate(w http.ResponseWriter, r *http.Request){

  params := mux.Vars(r)

  cpnID := params["cpnid"]
  cpn, err := strconv.Atoi(cpnID)
  if err != nil{
    fmt.Println(err)
    return
  }
  decoder := json.NewDecoder(r.Body)
  var errorData ErrorMsg

  err = decoder.Decode(&errorData)
  if err != nil{
    panic(err)
    w.WriteHeader(500)
    return
  }

  defer r.Body.Close()

  filter := bson.D{{"cpnid", cpn}}
  update := bson.D{{"$set", errorData}}

  updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
  if err != nil{
    log.Fatal(err)
  }
  Response(w, 200, errorData)
  fmt.Printf("Matched %v documents and updated %v documents. \n", updateResult.MatchedCount, updateResult.ModifiedCount)
}

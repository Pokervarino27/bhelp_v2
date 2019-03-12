package main

import(
  "fmt"
  "net/http"
  "time"
  "context"

  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)
func getClientDB(ctx context.Context) (*mongo.Database, error){

  client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://test:test@imaginex-developer-shard-00-00-ue0if.mongodb.net:27017"))
  if err != nil { return nil, err}
  ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
  defer cancel()
  err = client.Connect(ctx)

  db := client.Database("db_notify_track")
  if err != nil { return nil, err}

  fmt.Println("Connected to Mongo")
  return db, nil

}
//Index funcion de inicio de API
func Index(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "Hola, este es el inicio")
}

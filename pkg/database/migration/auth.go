package migration

import (
	"context"
	"log"

	"github.com/Kami0rn/MicroService/config"
	"github.com/Kami0rn/MicroService/module/auth"
	"github.com/Kami0rn/MicroService/pkg/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func authDbConn(pctx context.Context , cfg *config.Config) *mongo.Database{
	return database.DbConn(pctx, cfg).Database("auth_db")
}

func AuthMigrate (pctx context.Context , cfg *config.Config) {
	db := authDbConn(pctx,cfg)
	defer db.Client().Disconnect(pctx)

	col := db.Collection("auth")
	//Auth
	indexs, _ := col.Indexes().CreateMany(pctx, []mongo.IndexModel{
		{ Keys : bson.D{{"_id",1}}},
		{ Keys : bson.D{{"player_id" , 1}}},
		{ Keys : bson.D{{"refresh_token" , 1}}},
	})
	for _ , index := range indexs {
		log.Printf("Index: %s", index)
	}

	//Roles

	col = db.Collection("roles")

	indexs, _ = col.Indexes().CreateMany(pctx, []mongo.IndexModel{
		{ Keys : bson.D{{"_id",1}}},
		{ Keys : bson.D{{"code" , 1}}},
	})
	for _ , index := range indexs {
		log.Printf("Index: %s", index)
	}

	//roles data
	documents := func () []any {
		roles := []*auth.Role{
			{
				Title: "user",
				Code: 0,
			},
			{
				Title: "admin",
				Code: 1 ,
			},
		}

		docs := make([]any,0)
		for _,r := range roles {
			docs = append(docs,r)
		}
		return docs
	}()

	results, err := col.InsertMany(pctx, documents,nil)
	if err != nil {
		panic(err)
	}
	log.Printf("Migrate auth completed: %s" , results.InsertedIDs)

}
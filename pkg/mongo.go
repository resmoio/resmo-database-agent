package pkg

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"resmo-db-mapper/pkg/config"
	"resmo-db-mapper/pkg/resmo"
	"time"
)

func RunMongoQueries(ctx context.Context, config config.Config, dbType string) error {
	opts := options.Client().ApplyURI(config.DSN)
	log.Printf("connecting to database: %s\n", config.DSN)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return fmt.Errorf("error connecting to MongoDB: %w", err)
	}
	log.Printf("connected to database successfully\n")
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Printf("error disconnecting from MongoDB: %s", err)
		}
	}()

	cursor, err := client.ListDatabases(ctx, bson.M{})
	if err != nil {
		return fmt.Errorf("error listing databases: %w", err)
	}

	var databases []MongoDatabase
	var collections []MongoCollection
	var allUsers []MongoUser
	var allRoles []MongoRole

	for _, elem := range cursor.Databases {
		database := MongoDatabase{
			Name:  elem.Name,
			Size:  elem.SizeOnDisk,
			Empty: elem.Empty,
		}

		databases = append(databases, database)

		collectionInfo, err := listCollections(client, elem.Name)
		if err != nil {
			return fmt.Errorf("error listing collections: %w", err)
		}
		collections = append(collections, collectionInfo...)

		users, err := listMongoUsers(client, elem.Name)
		if err != nil {
			return fmt.Errorf("error listing users: %w", err)
		}
		allUsers = append(allUsers, users...)

		roles, err := listMongoRoles(ctx, client, elem.Name)
		if err != nil {
			return fmt.Errorf("error listing roles: %w", err)
		}
		allRoles = append(allRoles, roles...)
	}

	if err := resmo.Ingest(ctx, config, dbType, "mongo_database", databases); err != nil {
		return fmt.Errorf("error ingesting MongoDB database data: %w", err)
	}

	if err := resmo.Ingest(ctx, config, dbType, "mongo_collection", collections); err != nil {
		return fmt.Errorf("error ingesting MongoDB collection data: %w", err)
	}

	if err := resmo.Ingest(ctx, config, dbType, "mongo_user", allUsers); err != nil {
		return fmt.Errorf("error ingesting MongoDB user data: %w", err)
	}

	if err := resmo.Ingest(ctx, config, dbType, "mongo_role", allRoles); err != nil {
		return fmt.Errorf("error ingesting MongoDB role data: %w", err)
	}

	return nil
}

func listMongoRoles(ctx context.Context, client *mongo.Client, dbName string) ([]MongoRole, error) {
	rolesResult := struct {
		Roles []MongoRole `bson:"roles"`
	}{}
	err := client.Database(dbName).RunCommand(ctx, bson.D{{Key: "rolesInfo", Value: 1}}).Decode(&rolesResult)
	if err != nil {
		return nil, fmt.Errorf("error while listing mongo roles: %w", err)
	}

	return rolesResult.Roles, nil
}

func listMongoUsers(client *mongo.Client, dbName string) ([]MongoUser, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	usersResult := struct {
		Users []MongoUser `bson:"users"`
	}{}
	err := client.Database(dbName).RunCommand(ctx, bson.D{{Key: "usersInfo", Value: 1}}).Decode(&usersResult)
	if err != nil {
		return nil, fmt.Errorf("error while listing mongo users: %w", err)
	}
	return usersResult.Users, nil
}

func listCollections(client *mongo.Client, dbName string) ([]MongoCollection, error) {
	database := client.Database(dbName)

	collections, err := database.ListCollectionNames(context.Background(), bson.D{})
	if err != nil {
		return nil, fmt.Errorf("error while listing mongo collections: %w", err)
	}

	var collectionInfo []MongoCollection
	for _, collection := range collections {
		coll := database.Collection(collection)
		collectionData := MongoCollection{
			Name:     coll.Name(),
			Database: dbName,
		}
		collectionInfo = append(collectionInfo, collectionData)
	}

	return collectionInfo, nil
}

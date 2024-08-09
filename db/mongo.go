package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type mongoClient struct {
	client *mongo.Client
}

type mongoDatabase struct {
	db *mongo.Database
}
type mongoCollection struct {
	collection *mongo.Collection
}

type mongoSingleResult struct {
	singleResult *mongo.SingleResult
}

type mongoCursor struct {
	cursor *mongo.Cursor
}

type mongoSession struct {
	mongo.Session
}

///////////////////// mongoClient

func NewClient(connection string) (Client, error) {

	time.Local = time.UTC
	c, err := mongo.NewClient(options.Client().ApplyURI(connection))

	return &mongoClient{client: c}, err
}

func (mc *mongoClient) UseSession(ctx context.Context, fn func(mongo.SessionContext) error) error {
	return mc.client.UseSession(ctx, fn)
}

func (mc *mongoClient) StartSession() (mongo.Session, error) {
	session, err := mc.client.StartSession()
	return &mongoSession{session}, err
}

func (mc *mongoClient) Connect(ctx context.Context) error {
	return mc.client.Connect(ctx)
}

func (mc *mongoClient) Disconnect(ctx context.Context) error {
	return mc.client.Disconnect(ctx)
}

func (mc *mongoClient) Ping(ctx context.Context) error {
	return mc.client.Ping(ctx, readpref.Primary())
}

func (mc *mongoClient) Database(dbName string) Database {
	db := mc.client.Database(dbName)
	return &mongoDatabase{db: db}
}

///////////////////// mongoDatabase

func (md *mongoDatabase) Collection(colName string) Collection {
	collection := md.db.Collection(colName)
	return &mongoCollection{collection: collection}
}

func (md *mongoDatabase) Client() Client {
	client := md.db.Client()
	return &mongoClient{client: client}
}

///////////////////// mongoCollection

func (mc *mongoCollection) FindOne(ctx context.Context, filter interface{}) SingleResult {
	singleResult := mc.collection.FindOne(ctx, filter)
	return &mongoSingleResult{singleResult: singleResult}
}

func (mc *mongoCollection) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return mc.collection.UpdateOne(ctx, filter, update, opts[:]...)
}

func (mc *mongoCollection) InsertOne(ctx context.Context, document interface{}) (interface{}, error) {
	id, err := mc.collection.InsertOne(ctx, document)
	return id.InsertedID, err
}

func (mc *mongoCollection) InsertMany(ctx context.Context, document []interface{}) ([]interface{}, error) {
	res, err := mc.collection.InsertMany(ctx, document)
	return res.InsertedIDs, err
}

func (mc *mongoCollection) DeleteOne(ctx context.Context, filter interface{}) (int64, error) {
	count, err := mc.collection.DeleteOne(ctx, filter)
	return count.DeletedCount, err
}

func (mc *mongoCollection) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (Cursor, error) {
	findResult, err := mc.collection.Find(ctx, filter, opts...)
	return &mongoCursor{cursor: findResult}, err
}

func (mc *mongoCollection) Aggregate(ctx context.Context, pipeline interface{}) (Cursor, error) {
	aggregateResult, err := mc.collection.Aggregate(ctx, pipeline)
	return &mongoCursor{cursor: aggregateResult}, err
}

func (mc *mongoCollection) UpdateMany(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return mc.collection.UpdateMany(ctx, filter, update, opts[:]...)
}

func (mc *mongoCollection) CountDocuments(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	return mc.collection.CountDocuments(ctx, filter, opts...)
}

///////////////////// mongoSingleResult

func (sr *mongoSingleResult) Decode(v interface{}) error {
	return sr.singleResult.Decode(v)
}

///////////////////// mongoCursor

func (mr *mongoCursor) Close(ctx context.Context) error {
	return mr.cursor.Close(ctx)
}

func (mr *mongoCursor) Next(ctx context.Context) bool {
	return mr.cursor.Next(ctx)
}

func (mr *mongoCursor) Decode(v interface{}) error {
	return mr.cursor.Decode(v)
}

func (mr *mongoCursor) All(ctx context.Context, result interface{}) error {
	return mr.cursor.All(ctx, result)
}

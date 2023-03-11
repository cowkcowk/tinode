//go:build mongodb
// +build mongodb

// Package mongodb is a database adapter for MongoDB.
package mongodb

import (
	mdb "go.mongodb.org/mongo-driver/mongo"
)

// adapter holds MongoDB connection data.
type adapter struct {
	conn   *mdb.Client
	db     *mdb.Database
	dbName string
	// Maximum number of records to return
	maxResults int
	// Maximum number of message records to return
	maxMessageResults int
	version           int
	ctx               context.Context
	useTransactions   bool
}

// Open initializes mongodb session
func (a *adapter) Open(jsonconfig json.RawMessage) error {
	if a.conn != nil {
		return errors.New("adapter mongodb is already connected")
	}

	if len(jsonconfig) < 2 {
		return errors.New("adapter mongodb missing config")
	}

	var err error

}

// Close the adapter
func (a *adapter) Close() error {
	var err error
	if a.conn != nil {
		err = a.conn.Disconnect(a.ctx)
		a.conn = nil
		a.version = -1
	}
	return err
}

// IsOpen checks if the adapter is ready for use
func (a *adapter) IsOpen() bool {
	return a.conn != nil
}

// CreateDb creates the database optionally dropping an existing database first.
func (a *adapter) CreateDb(reset bool) error {
	if reset {
		if err := a.db.Drop(a.ctx); err != nil {
			return err
		}
	} else if a.isDbInitialized() {
		return errors.New("Database already initialized")
	}
	// Collections (tables) do not need to be explicitly created since MongoDB creates them with first write operation

	indexes := []struct {
		Collection string
		Field      string
		IndexOpts  mdb.IndexModel
	}{
		// Users
		// Index on 'user.state' for finding suspended and soft-deleted users.
		{
			Collection: "users",
			Field:      "state",
		},
		// Index on 'user.tags' array so user can be found by tags.
		{
			Collection: "users",
			Field:      "tags",
		},
		// Index for 'user.devices.deviceid' to ensure Device ID uniqueness across users.
		// Partial filter set to avoid unique constraint for null values (when user object have no devices).
		{
			Collection: "users",
			IndexOpts: mdb.IndexModel{
				Keys: b.M{"devices.deviceid": 1},
				Options: mdbopts.Index().
					SetUnique(true).
					SetPartialFilterExpression(b.M{"devices.deviceid": b.M{"$exists": true}}),
			},
		},

		// User authentication records {_id, userid, secret}
		// Should be able to access user's auth records by user id
		{
			Collection: "auth",
			Field:      "userid",
		},

		// Subscription to a topic. The primary key is a topic:user string
		{
			Collection: "subscriptions",
			Field:      "user",
		},
		{
			Collection: "subscriptions",
			Field:      "topic",
		},

		// Topics stored in database
		// Index on 'owner' field for deleting users.
		{
			Collection: "topics",
			Field:      "owner",
		},
		// Index on 'state' for finding suspended and soft-deleted topics.
		{
			Collection: "topics",
			Field:      "state",
		},
		// Index on 'topic.tags' array so topics can be found by tags.
		// These tags are not unique as opposite to 'user.tags'.
		{
			Collection: "topics",
			Field:      "tags",
		},

		// Stored message
		// Compound index of 'topic - seqid' for selecting messages in a topic.
		{
			Collection: "messages",
			IndexOpts:  mdb.IndexModel{Keys: b.D{{"topic", 1}, {"seqid", 1}}},
		},
		// Compound index of hard-deleted messages
		{
			Collection: "messages",
			IndexOpts:  mdb.IndexModel{Keys: b.D{{"topic", 1}, {"delid", 1}}},
		},
		// Compound multi-index of soft-deleted messages: each message gets multiple compound index entries like
		// 		 [topic, user1, delid1], [topic, user2, delid2],...
		{
			Collection: "messages",
			IndexOpts:  mdb.IndexModel{Keys: b.D{{"topic", 1}, {"deletedfor.user", 1}, {"deletedfor.delid", 1}}},
		},

		// Log of deleted messages
		// Compound index of 'topic - delid'
		{
			Collection: "dellog",
			IndexOpts:  mdb.IndexModel{Keys: b.D{{"topic", 1}, {"delid", 1}}},
		},

		// User credentials - contact information such as "email:jdoe@example.com" or "tel:+18003287448":
		// Id: "method:credential" like "email:jdoe@example.com". See types.Credential.
		// Index on 'credentials.user' to be able to query credentials by user id.
		{
			Collection: "credentials",
			Field:      "user",
		},

		// Records of file uploads. See types.FileDef.
		// Index on 'fileuploads.usecount' to be able to delete unused records at once.
		{
			Collection: "fileuploads",
			Field:      "usecount",
		},
	}

	var err error
	for _, idx := range indexes {
		if idx.Field != "" {
			_, err = a.db.Collection(idx.Collection).Indexes().CreateOne(a.ctx, mdb.IndexModel{Keys: b.M{idx.Field: 1}})
		} else {
			_, err = a.db.Collection(idx.Collection).Indexes().CreateOne(a.ctx, idx.IndexOpts)
		}
		if err != nil {
			return err
		}
	}

	// Collection "kvmeta" with metadata key-value pairs.
	// Key in "_id" field.
	// Record current DB version.
	if _, err := a.db.Collection("kvmeta").InsertOne(a.ctx, map[string]interface{}{"_id": "version", "value": adpVersion}); err != nil {
		return err
	}

	// Create system topic 'sys'.
	return createSystemTopic(a)
}


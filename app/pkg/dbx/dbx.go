// Package dbx is database connections and queries
package dbx

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"log/slog"

	"github.com/dhinesh20022806/api_s3/app/pkg/env"
	// import pq packages
	_ "github.com/lib/pq"
)

var db *sql.DB

// Bucket is used to store collections of objects
type Bucket struct {
	ID         string
	BucketName string
	OwnerID    string
}

func init() {

	var err error
	fmt.Printf("database url %s\n", env.Config.Database.URL)
	db, err = sql.Open("postgres", env.Config.Database.URL)

	if err != nil {

		log.Printf("error to connect database %v", err)
	}

	if err = db.Ping(); err != nil {

		log.Printf("error to ping the database %v", err)
	}
}

// InsertMetaData store the meta of object in database
func InsertMetaData(id, filename, bucketID, ownerID string) error {

	var name string

	selectQuery := `SELECT name FROM buckets WHERE id = $1 AND owner_id = $2;`

	slog.Info("select query", "bucketId", bucketID, "owner id", ownerID)
	err := db.QueryRow(selectQuery, bucketID, ownerID).Scan(&name)

	if err != nil {
		return err
	}

	query := `INSERT INTO objects (id, name, bucket_name, owner_id) VALUES
	 ($1, $2, $3, $4)`

	_, err = db.Exec(query, id, filename, name, ownerID)

	if err != nil {
		return err
	}

	return nil
}

// IsMyBucketID check wheather bucketID owned or not
func IsMyBucketID(bucketID, ownerID string) (bool, error) {

	var bucketName string
	query := `SELECT bucketName FROM bucket WHERE bucketId = $1 AND owner_id = $2`

	err := db.QueryRow(query, bucketID, ownerID).Scan(
		&bucketName,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// GetMyObject allow authorised people access their objects
func GetMyObject(objectName, bucketName, ownerID string) (string, error) {

	var objectID string
	query := `SELECT id FROM objects WHERE name = $1 AND bucket_name = $2 AND owner_id = $3;`

	err := db.QueryRow(query, objectName, bucketName, ownerID).Scan(
		&objectID,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", nil
		}
		return objectID, err
	}
	return objectID, nil
}

// CreateBucket used to create new unique bucketName there is sql constraint for creating name
func CreateBucket(bucketName, ownerID string) (*Bucket, error) {

	var object Bucket

	query := `INSERT INTO buckets(name, owner_id) 
	VALUES($1, $2)
	RETURNING id, name, owner_id`

	err := db.QueryRow(query, bucketName, ownerID).Scan(
		&object.ID,
		&object.BucketName,
		&object.OwnerID,
	)

	if err != nil {
		return nil, err
	}

	return &object, nil
}

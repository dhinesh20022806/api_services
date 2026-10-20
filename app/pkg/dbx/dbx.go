package dbx

import (
	"database/sql"
	"errors"
)

var db *sql.DB

type Bucket struct {
	Id         string
	BucketName string
	OwnerId    string
}

func init() {

	var err error

	db, err = sql.Open("postgres", "")

	if err != nil {

	}

	if err = db.Ping(); err != nil {

	}
}

func InsertMetaData(id, filename, bucketId, types string) error {

	query := `INSERT INTO object (id, filename, bucketId, types) VALUES
	 ($1, $2, $3)`

	_, err := db.Exec(query, id, filename, bucketId, types)

	if err != nil {
		return err
	}

	return nil
}

func IsMyBucketId(bucketId, owner_id string) (bool, error) {

	var bucketName string
	query := `SELECT bucketName FROM bucket WHERE bucketId = $1 AND owner_id = $2`

	err := db.QueryRow(query, bucketId, owner_id).Scan(
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

func CreateBucket(bucketName, owner_id string) (*Bucket, error) {

	var object Bucket

	query := `INSERT INTO bucket(bucket_name, owner_id) 
	VALUES($1, $2)
	RETURNING id, bucket_name, owner_id`

	err := db.QueryRow(query, bucketName, owner_id).Scan(
		&object.Id,
		&object.bucketName,
		&object.OwnerId,
	)

	if err != nil {
		return nil, err
	}

	return &object, nil
}

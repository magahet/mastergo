package quotes

import (
	"fmt"

	"github.com/coreos/bbolt"
)

// DB is a quote database.
type DB struct {
	db *bolt.DB
}

const (
	quoteBucket = "standard"
)

// Open opens the database file at path and returns a DB or an error.
func Open(path string) (*DB, error) {
	db, err := bolt.Open(path, 0600, nil)
	return &DB{db}, err
}

func (d *DB) Close() error {
	return d.db.Close()
}

// Create takes a quote and saves it to the database, using the author name
// as the key. If the author already exists, Create returns an error.
func (d *DB) Create(q *Quote) error {
	err := d.db.Update(func(tx *bolt.Tx) error {

		b, err := tx.CreateBucketIfNotExists([]byte(quoteBucket))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}

		v := b.Get([]byte(q.Author))
		if v != nil {
			return fmt.Errorf("quote already exsists")
		}

		serializedQuote, err := q.Serialize()
		if err != nil {
			return fmt.Errorf("serialize quote: %s", err)
		}

		err = b.Put([]byte(q.Author), serializedQuote)
		if err != nil {
			return fmt.Errorf("put quote: %s", err)
		}

		return nil
	})

	return err
}

// Get takes an author name and retrieves the corresponding quote from the DB.
func (d *DB) Get(author string) (*Quote, error) {
	q := &Quote{}
	err := d.db.View(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte(quoteBucket))
		if b == nil {
			return fmt.Errorf("no bucket")
		}

		v := b.Get([]byte(author))
		if v == nil {
			return fmt.Errorf("no quote by that author")
		}

		return q.Deserialize(v)
	})

	if err != nil {
		return nil, err
	}
	return q, nil
}

// List lists all records in the DB.
func (d *DB) List() ([]*Quote, error) {
	// The database returns byte slices that we need to de-serialize
	// into Quote structures.
	structList := []*Quote{}

	// We use a View as we don't update anything.
	err := d.db.View(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte(quoteBucket))
		if b == nil {
			return fmt.Errorf("no bucket")
		}

		err := b.ForEach(func(k, v []byte) error {
			q := &Quote{}
			err := q.Deserialize(v)
			if err != nil {
				return err
			}

			structList = append(structList, q)
			return nil
		})

		return err
	})

	if err != nil {
		return nil, err
	}
	return structList, nil
}

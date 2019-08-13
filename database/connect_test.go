package database

import (
	"testing"
	"time"
)

func TestMonitor(t *testing.T) {
	t.Parallel()

	t.Run("runs and dies on connection close", func(t *testing.T) {
		t.Parallel()

		// connects
		db, err := Connect(Config{
			Name:     "autograph",
			User:     "myautographdbuser",
			Password: "myautographdbpassword",
			Host:     "127.0.0.1:5432",
		})
		if err != nil {
			t.Fatal(err)
		}

		quit := make(chan bool)
		go db.Monitor(5*time.Millisecond, quit)

		// should not error for initial monitor run
		err = db.CheckConnection()
		if err != nil {
			t.Fatalf("db.CheckConnection failed when it should not have with error: %s", err)
		}
		time.Sleep(10 * time.Millisecond)

		// error for failing checks
		db.Close()
		err = db.CheckConnection()
		if err == nil {
			t.Fatalf("db.CheckConnection did not fail for a closed DB")
		}

		quit <- true
	})
}
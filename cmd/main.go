package main

import (
	"fmt"
	"github.com/abdymazhit/go-test-task/pkg/handler"
	"github.com/abdymazhit/go-test-task/pkg/repository"
	"github.com/abdymazhit/go-test-task/pkg/service"
	"github.com/boltdb/bolt"
	"github.com/valyala/fasthttp"
	"log"
)

func main() {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	var productsBucket, productIndexBucket *bolt.Bucket
	if err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("product"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		productsBucket = b

		b, err = tx.CreateBucketIfNotExists([]byte("productIndex"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		productIndexBucket = b

		return nil
	}); err != nil {
		log.Fatal(err)
	}

	repos := repository.NewRepository(productsBucket, productIndexBucket)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	r := handlers.InitRoutes()
	log.Fatal(fasthttp.ListenAndServe(":8080", r.Handler))
}

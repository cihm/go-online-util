package redis

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

func Testflow(url string) {

	client := redis.NewClient(&redis.Options{
		Addr:     url,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>

	//nano
	err = client.Set("key", "123", 7*time.Second).Err()
	if err != nil {
		panic(err)
	}

	err = client.Set("key2", "456", 13*time.Second).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist

	interval := 5 * time.Second
	go func() {
		c := time.Tick(interval)
		loop := true
		for loop {
			select {
			case <-c:
				val, err := client.Get("key").Result()
				if err != nil {
					fmt.Println("key", "expire")
				}
				fmt.Println("key", val)

				val2, err2 := client.Get("key2").Result()
				if err2 != nil {
					fmt.Println("key2", "expire")
				}
				fmt.Println("key2", val2)

			}
		}
	}()

}

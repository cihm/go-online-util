package elsearch

import (
	"context"
	"fmt"

	"github.com/olivere/elastic"
)

func Updateflow(serverUrl string) {

	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(serverUrl))
	if err != nil {
		// Handle error
		fmt.Println(err.Error())
		panic(err)
	}

	//Update exist
	e1 := Employee{"Jane10", "Smith", 333, "I like to collect rock albums 33", []string{"music33"}}
	res, err := client.Update().
		Index("db1").
		Type("employee").
		Id("3").
		Doc(e1).
		Do(context.Background())
	if err != nil {
		println(err.Error())
	}
	fmt.Printf("update age %s\n", res.Result)

	exist, err := client.Exists().Index("db1").Type("employee").Id("10").Do(context.Background())
	if err != nil {
		println(err.Error())
	}
	if exist {
		fmt.Printf("Got document 10\n")
	} else {
		//not exist
		e1 = Employee{"Jane10", "Smith", 10, "I like to collect rock albums 10", []string{"music10"}}
		res, err := client.Index().
			Index("db1").
			Type("employee").
			Id("10").
			BodyJson(e1).
			Do(context.Background())
		if err != nil {
			println(err.Error())
		}
		fmt.Printf("update age %s\n", res.Result)
	}
}

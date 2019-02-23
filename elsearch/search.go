package elsearch

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/olivere/elastic"
)

func Searchflow(serverUrl string) {
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(serverUrl))
	if err != nil {
		// Handle error
		fmt.Println(err.Error())
		panic(err)
	}

	get1, err := client.Get().Index("db1").Type("employee").Id("2").Do(context.Background())
	if err != nil {
		panic(err)
	}
	if get1.Found {
		fmt.Printf("Got document %s in version %d from index %s, type %s\n", get1.Id, get1.Version, get1.Index, get1.Type)
	}

	size := 2
	page := 1
	//字段相等
	for {
		q := elastic.NewQueryStringQuery("last_name:Smith")
		res, err := client.Search("db1").
			Type("employee").
			Query(q).
			Size(size).
			From((page - 1) * size).
			Do(context.Background())
		if err != nil {
			println(err.Error())
		}
		printEmployee(res, err)
		fmt.Println("!!!", res.Hits.Hits)
		page++
		if len(res.Hits.Hits) < (size) || page < 1 {
			fmt.Printf("done")
			break
		}
	}

	q := elastic.NewMultiMatchQuery("嘉中英文試題爆侮辱蔡總統").
		FieldWithBoost("about", 2).
		FieldWithBoost("last_name", 1)
	//FieldWithBoost("first_name", 2).
	res, err := client.Search("db1").
		Type("employee").
		Query(q).
		Do(context.Background())
	if err != nil {
		println(err.Error())
	}
	printEmployee(res, err)
	fmt.Println("!total:", res.Hits.TotalHits)
}

func printEmployee(res *elastic.SearchResult, err error) {

	if res.Hits.TotalHits > 0 {
		fmt.Printf("Found a total of %d Employee \n", res.Hits.TotalHits)

		for _, hit := range res.Hits.Hits {

			var t Employee
			err := json.Unmarshal(*hit.Source, &t) //另外一种取数据的方法
			if err != nil {
				fmt.Println("Deserialization failed")
			}

			fmt.Printf("Employee name %s : %s\n", t.FirstName, t.LastName)
		}
	} else {
		fmt.Printf("Found no Employee \n")
	}

	if err != nil {
		print(err.Error())
		return
	}
	var typ Employee
	for _, item := range res.Each(reflect.TypeOf(typ)) { //从搜索结果中取数据的方法
		t := item.(Employee)
		fmt.Printf("%#v\n", t)
	}
}

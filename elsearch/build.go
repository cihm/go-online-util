package elsearch

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/olivere/elastic"
)

func Buildflow(serverUrl string) {

	ctx := context.Background()
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(serverUrl))
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	// Ping the Elasticsearch server to get e.g. the version number
	info, code, err := client.Ping(serverUrl).Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	// Getting the ES version number is quite common, so there's a shortcut
	esversion, err := client.ElasticsearchVersion(serverUrl)
	if err != nil {
		// Handle error
		panic(err)
	}

	fmt.Printf("Elasticsearch version %s\n", esversion)

	//Clear if exixst
	exists, err := client.IndexExists("db1").Do(ctx)
	if err != nil {
		log.Println("IndexExists" + err.Error())
	}
	if exists {
		res, err := client.DeleteIndex().
			Index([]string{"db1"}).
			Do(context.Background())
		if err != nil {
			panic(err)
		}
		fmt.Printf("delete  \n", res)
	}

	//Create index
	createIndex, err := client.CreateIndex("db1").Do(ctx)
	if err != nil {
		log.Println("CreateIndex" + err.Error())

	}
	if !createIndex.Acknowledged {
		// Not acknowledged
		log.Println("create index:db1" + ", not Ack")
	}

	//Mapping for ik analyzer
	//https://github.com/olivere/elastic/issues/755
	typeMapping := `{
		"employee":{
			"properties":{
				"about":{
					"analyzer": "ik_smart",
					"type": "text"		
				},"last_name":{
					"analyzer": "ik_smart",
					"type":"text"
				}
			}
		}
	}`
	putresp, err := client.PutMapping().Index("db1").Type("employee").BodyString(typeMapping).Do(context.TODO())
	if err != nil {
		log.Println("error NewIndicesCreateService" + err.Error())
	}
	if !putresp.Acknowledged {
		// Not acknowledged
		log.Println("!putresp.Acknowledgeds NewIndicesCreateService")
	}

	var a [10]string
	a[0] = "羞辱蔡英文？英文考題惹議 嘉中命題老師道歉"
	a[1] = "國立嘉義高中日前一份高二英文測驗考題近來在校園引起討論，學生質疑老師命題泛政治化，藉考題辱罵總統蔡英文。命題張姓老師澄."
	a[2] = "嘉中英文試題爆侮辱蔡總統"
	a[3] = "總統府發言人黃重諺昨天透過臉書暗喻高雄市長韓國瑜"
	a[4] = "韓國瑜回擊：狗嘴吐不出象牙"
	a[5] = "回嗆，這是國家發言人講話的高度"
	a[6] = "高雄市長韓國瑜前日深夜在高雄市酒吧開直播"
	a[7] = "元宵新吃法 小湯圓變拔絲、熱壓湯圓吐司"
	a[8] = "元宵節即將到來，很多超市都忙著補貨抓準這波商機，沒想到大家最喜歡的還是傳統滋味，芝麻、花生、鮮肉占了前三名"
	a[9] = "市府表揚考核績優醫院"

	var b [10]string
	b[0] = "回嗆"
	b[1] = "高雄市長"
	b[2] = "元宵"
	b[3] = "鮮肉"
	b[4] = "市府表揚考"
	b[5] = "羞辱蔡英文"
	b[6] = "國立"
	b[7] = "李鈺新"
	b[8] = "總統府發言人"
	b[9] = "狗嘴吐不出象牙"

	for i := 0; i < 10; i++ {
		si := strconv.Itoa(i)
		e1 := Employee{"Jane" + si, b[i], 32 + i, a[i], []string{"music" + si}}
		put1, err := client.Index().
			Index("db1").
			Type("employee").
			Id(si).
			BodyJson(e1).
			Do(context.Background())
		if err != nil {
			panic(err)
		}
		fmt.Printf("Indexed tweet %s to index s%s, type %s\n", put1.Id, put1.Index, put1.Type)
	}

}

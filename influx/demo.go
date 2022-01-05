package  main

//func writesPoints()  {
//	client := influxdb2.NewClient("http://localhost:8086","ThN5VZEOlt8RGoKJjFb0O011FY5bFTdig9wjlzfHAGsEmVitFmVRaajRLGZXna5WLoyzu8dOm-GpzIRSLEC2og==")
//	writeAPI := client.WriteAPIBlocking("q_org","q_bucket")
//	p := influxdb2.NewPoint("stat",
//		map[string]string{"unit":"temperature"},
//		map[string]interface{}{"avg":24.5,"max":45.0},
//		time.Now(),
//	)
//
//	writeAPI.WritePoint(context.Background(),p)
//	line := fmt.Sprintf("stat,unit=temperature avg=%f,max=%f",23.5,45.0)
//	writeAPI.WriteRecord(context.Background(),line)
//
//	queryAPI := client.QueryAPI("q_org")
//
//	result, err := queryAPI.Query(context.Background(),`from(bucket:"my-bucket")|> range(start: -1h) |> filter(fn: (r) => r._measurement == "stat")`)
//	if err ==nil{
//		for result.Next(){
//			if result.TableChanged(){
//				fmt.Printf("table: %s\n",result.TableMetadata().String())
//			}
//			fmt.Printf("row: %s\n",result.Record().String())
//		}
//		if result.Err() != nil{
//			fmt.Printf("Query error : %s \n",result.Err().Error())
//		}
//	}
//
//	client.Close()
//}
//
//func main()  {
//	writesPoints()
//	fmt.Println("----done-----!")
//}
package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"sync"
)

func main()  {
	consumer ,err := sarama.NewConsumer([]string{"127.0.0.1:9092"},nil)
	if err != nil{
		fmt.Printf("fail to start consumer,err:%v\n",err)
		return
	}

	partitionList, err:=consumer.Partitions("web_log")

	if err != nil{
		fmt.Printf("fail to get list of partition: err%v\n",err)
		return
	}

	fmt.Println(partitionList)
	var wg sync.WaitGroup
	for partition := range partitionList{
		pc, err := consumer.ConsumePartition("web_log", int32(partition), sarama.OffsetNewest)
		if err !=nil{
			fmt.Println("error")
			return
		}
		defer pc.AsyncClose()
		wg.Add(1)
		go func(sarama.PartitionConsumer) {
			for msg :=range pc.Messages(){
				fmt.Printf("Partiton:%d offset:%d key:%s value:%s\n",msg.Partition,msg.Offset,msg.Key,msg.Value)
			}
		}(pc)
	}
	wg.Wait()
}

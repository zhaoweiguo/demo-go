package main

import (
	sls "github.com/aliyun/aliyun-log-go-sdk"
	"log"
	"time"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	AccessKeyID := "LTAI4GCXuwsdwu68KwLsAvFa"
	AccessKeySecret := "48obgTFkKNw9C1ub3OvJqanflWxw6v"
	Endpoint := "cn-beijing.log.aliyuncs.com"
	client := sls.CreateNormalInterface(Endpoint, AccessKeyID, AccessKeySecret, "")

	//project := "k8s-log-ce32cf0d5d77a42f188f25cc6ab98ab73"
	//logstore := "lock-install-daily-api"
	project := "iot-ol-engine-log"
	logstore := "octopus-logs-api"

	// 查看shard列表
	shards, err := client.ListShards(project, logstore)
	if err != nil {
		log.Println(err)
	}

	totalLogCount := 0
	for _, shard := range shards {
		log.Printf("[shard] %d begin\n", shard.ShardID)
		beginCursor, err := client.GetCursor(project, logstore, shard.ShardID, "begin")
		if err != nil {
			panic(err)
		}
		endCursor, err := client.GetCursor(project, logstore, shard.ShardID, "end")
		if err != nil {
			panic(err)
		}
		nextCursor := beginCursor
		for nextCursor != endCursor {
			gl, nc, err := client.PullLogs(project, logstore, shard.ShardID, nextCursor, endCursor, 10)
			if err != nil {
				log.Printf("pull log error : %s\n", err)
				time.Sleep(time.Second)
				continue
			}
			nextCursor = nc
			log.Printf("now count %d \n", totalLogCount)
			if gl != nil {
				for _, lg := range gl.LogGroups {
					log.Println(len(lg.LogTags))
					log.Println(len(lg.Logs))
					log.Println(len(lg.Logs[0].Contents))
					log.Println(lg.String())

				}
			}
		}
	}
}

package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/zhaoweiguo/demo-go/pkg/net/http/demo/db"
	"log"
)

const InnerServer = "Server suffers an error"

func main()  {
	host := ""
	dbname := "video"
	table := "video"
	db.Conn(host, dbname, table)
	router := gin.Default()
	router.GET("/v1.0/video/list", getVideoList)
	router.POST("/v1.0/video", addVideo)
	router.PUT("/v1.0/video", updateVideo)
	router.DELETE("/v1.0/video", deleteVideo)
}


func getVideoList(c *gin.Context) {
	userName := c.GetString("user_name")
	video := db.New(userName, "", "")
	list := video.GetVideoList()
	jsonList, err := json.Marshal(list)
	if err != nil {
		log.Println(err.Error())
		returnErrResponse(c, InnerServer)
	}
	returnSuccResponse(c, jsonList)
}

func addVideo(c *gin.Context) {
	userName := c.GetString("user_name")
	videoName := c.GetString("video_name")
	videoType := c.GetString("video_type")

	video := db.New(userName, videoType, videoName)
	isSuccess := video.AddVideo()
	if isSuccess != true {
		returnErrResponse(c, InnerServer)
	}
	returnSuccResponse(c, []byte("success"))
}
func updateVideo(c *gin.Context) {
	id := c.GetInt("id")

	videoName := c.GetString("video_name")
	video := db.GetVideo(id)
	video.SetVideoName(videoName)
	returnSuccResponse(c, []byte("success"))
}

func deleteVideo(c *gin.Context) {
	id := c.GetInt("id")
	video := db.GetVideo(id)
	video.DelVideo()
	returnSuccResponse(c, []byte("success"))
}


// ============
// inner func
// ============
func returnErrResponse(c *gin.Context, msg string)  {
	c.JSON(200, gin.H{"msg": msg})
}

func returnSuccResponse(c *gin.Context, msg []byte)  {
	c.JSON(200, gin.H{"code" : 200, "msg" : string(msg)})
}
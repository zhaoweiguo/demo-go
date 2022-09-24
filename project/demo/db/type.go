package db

import (
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

type Video struct {
	Id        int       `bson:"_id"`
	UserName  string    `bson:"user_name"`
	VideoName string    `bson:"video_name"`
	VideoType string    `bson:"video_type"`
	IsDeleted bool      `bson:"isdeleted"`
	UpdatedAt time.Time `bson:"updated_at"`
	CreatedAt time.Time `bson:"created_at"`
}

func New(userName string, videoType string, videoName string) Video {
	return Video{
		UserName:  userName,
		VideoName: videoName,
		VideoType: videoType,
	}
}
func GetVideo(id int) Video {
	video := Video{
		Id: id,
	}
	return video
}

func (video *Video) SetVideoName(videoName string) error {
	video.VideoName = videoName
	return collection.UpdateId(video.Id, video)
}

func (video *Video) GetVideoList() *Video {
	query := bson.M{"user_name": video.UserName}
	err := collection.Find(query).One(video)
	if err != nil {
		log.Panic(err)
	}
	return video
}

func (video *Video) AddVideo() bool {
	err := collection.Insert(video)
	if err != nil {
		log.Panic(err)
	}
	return true
}

func (video *Video) DelVideo() bool {
	err := collection.RemoveId(video.Id)
	if err != nil {
		log.Panic(err)
	}
	return true
}

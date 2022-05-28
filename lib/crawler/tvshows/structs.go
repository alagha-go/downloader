package tvshows

import "go.mongodb.org/mongo-driver/bson/primitive"


type TvShow struct {
	ID											primitive.ObjectID						`json:"_id,omitempty" bson:"_id,omitempty"`
	Code										string									`json:"code,omitempty" bson:"code,omitempty"`
	Title										string									`json:"title,omitempty" bson:"title,omitempty"`
	Type										string									`json:"type,omitempty" bson:"type,omitempty"`
    Available                                   bool                                    `json:"available,omitempty" bson:"available,omitempty"`
    PageUrl										string									`json:"page_url,omitempty" bson:"page_url,omitempty"`
    ImageUrl									string									`json:"image_url,omitempty" bson:"image_url"`
}


type Season struct {
	ID											primitive.ObjectID						`json:"_id,omitempty" bson:"_id,omitempty"`
	Name										string									`json:"name,omitempty" bson:"name,omitempty"`
    Index                                       int                                     `json:"index,omitempty" bson:"index,omitempty"`
    Code										string									`json:"code,omitempty" bson:"code,omitempty"`
}

type Episode struct {
	ID											primitive.ObjectID						`json:"_id,omitempty" bson:"_id,omitempty"`
	Name										string									`json:"name,omitempty" bson:"name,omitempty"`
    Index                                       int                                     `json:"index,omitempty" bson:"index,omitempty"`
    Available                                   bool                                    `json:"available,omitempty" bson:"available,omitempty"`
    Code										string									`json:"code,omitempty" bson:"code,omitempty"`
	Server										*Server									`json:"server,omitempty" bson:"server,omitempty"`
	Servers										[]Server								`json:"servers,omitempty" bson:"servers,omitempty"`
}

type Server struct {
	ID											primitive.ObjectID						`json:"_id,omitempty" bson:"_id,omitempty"`
    Name                                        string                               	`json:"name,omitempty" bson:"name,omitempty"`
    Id                                          string                               	`json:"id,omitempty" bson:"id,omitempty"`
    WatchID                                     string                               	`json:"watch_id,omitempty" bson:"watch_id,omitempty"`
    Url                                         string                               	`json:"url,omitempty" bson:"url,omitempty"`
}


type LinkResponse struct {
	Type                                        string                                   `json:"type,omitempty" bson:"type,omitempty"`   
	Link                                        string                                   `json:"link,omitempty" bson:"link,omitempty"`   
	Sources                                     []interface{}                            `json:"sources,omitempty" bson:"sources,omitempty"`
	Tracks                                      []interface{}                            `json:"tracks,omitempty" bson:"tracks,omitempty"` 
	Title                                       string                                   `json:"title,omitempty" bson:"title,omitempty"`  
}
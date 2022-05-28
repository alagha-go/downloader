package movies

import "go.mongodb.org/mongo-driver/bson/primitive"

type Movie struct {
	ID											primitive.ObjectID						`json:"_id,omitempty" bson:"_id,omitempty"`
	Code										string									`json:"code,omitempty" bson:"code,omitempty"`
	Title										string									`json:"title,omitempty" bson:"title,omitempty"`
    PageUrl										string									`json:"page_url,omitempty" bson:"page_url,omitempty"`
    ImageUrl									string									`json:"image_url,omitempty" bson:"image_url"`
}

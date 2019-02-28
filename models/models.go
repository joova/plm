package models

import "github.com/mongodb/mongo-go-driver/bson/primitive"

// UnitOfMeasure UnitOfMeasure model
type UnitOfMeasure struct {
	Code string             `bson:"code,omitempty" json:"code,omitempty"`
	Name string             `bson:"name" json:"name"`
	Org  primitive.ObjectID `bson:"org,omitempty" json:"org,omitempty"`
}

// ProductType ProductType model
type ProductType struct {
	Code string             `bson:"code,omitempty" json:"code,omitempty"`
	Name string             `bson:"name" json:"name"`
	Org  primitive.ObjectID `bson:"org,omitempty" json:"org,omitempty"`
}

// ProductCategory ProductCategory model
type ProductCategory struct {
	Code string             `bson:"code,omitempty" json:"code,omitempty"`
	Name string             `bson:"name" json:"name"`
	Org  primitive.ObjectID `bson:"org,omitempty" json:"org,omitempty"`
}

// Product product model
type Product struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Code     string             `bson:"code" json:"code"`
	Name     string             `bson:"name" json:"name"`
	UOM      string             `bson:"uom" json:"uom"`
	Type     string             `bson:"type" json:"type"`
	Category string             `bson:"category" json:"category"`
	Org      primitive.ObjectID `bson:"org,omitempty" json:"org,omitempty"`
}

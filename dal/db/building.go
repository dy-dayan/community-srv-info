package db

type Building struct {
	ID   string    `bson:"_id"`
	Name string    `bson:"name"`
	Loc  []float32 `bson:"loc"`
}

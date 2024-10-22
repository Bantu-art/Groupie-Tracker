package types

type Relation struct {
	Id       int                 `json:"id"`
	Concerts map[string][]string `json:"datesLocations"`
}

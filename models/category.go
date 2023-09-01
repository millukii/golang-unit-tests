package models

type CategoryCountry struct {
	Id       int
	Country  string
	Category int
	Enabled  bool
}
type Category struct {
	Id   int
	Name string
}

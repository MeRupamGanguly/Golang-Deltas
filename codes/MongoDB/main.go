package main

type Status string

const (
	Pending  Status = "Pending"
	Approved Status = "Approved"
	Rejected Status = "Rejected"
)

type Artist struct {
	Id   string
	Name string
}
type Release struct {
	PlatformId string
	Status     Status
	Artists    []Artist
}
type Album struct {
	Id       string
	Name     string
	LabelId  string
	Releases []Release
}
type Song struct {
	Id      string
	Name    string
	AlbumId string
}

type Label struct {
	Id   string
	Name string
}

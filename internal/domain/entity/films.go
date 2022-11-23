package entity

type Movie struct {
	Name        string
	ReleaseDate string
	Assessments map[User]string
}

func NewMovie(name string, releaseDate string, assessments map[User]string) *Movie {
	return &Movie{Name: name, ReleaseDate: releaseDate, Assessments: assessments}
}

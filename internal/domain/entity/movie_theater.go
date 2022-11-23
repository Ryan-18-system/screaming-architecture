package entity

type MovieTheater struct {
	Name  string
	Films []Movie // slice de filmes
}

func NewMovieTheater(name string, films ...Movie) *MovieTheater {
	return &MovieTheater{Name: name, Films: films}
}

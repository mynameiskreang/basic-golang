package movie

import (
	"basic-golang/chapter-four/models"
	"basic-golang/chapter-four/structures/film"
	"github.com/sirupsen/logrus"
)

func GetMovies() (movies film.Movies, err error) {
	rows, errQuery := models.PgxConn.Query("select * from film limit 10")
	if errQuery != nil {
		defer logrus.WithFields(logrus.Fields{
			"database": "postgress",
			"action":   "GetMovies",
		}).Error(errQuery)
		return movies, errQuery
	}
	for rows.Next() {
		movie := film.Movie{}

		errScan := rows.Scan(&movie.FilmID, &movie.Title, &movie.Description, &movie.ReleaseYear, &movie.LanguageID, &movie.RentalDuration, &movie.RentalRate, &movie.Length, &movie.ReplacementCost, &movie.Rating, &movie.LastUpdate, &movie.SpecialFeatures, &movie.Fulltext)
		if errScan != nil {
			logrus.WithFields(logrus.Fields{
				"database": "postgress",
				"action":   "GetMovies",
			}).Error(errScan)
		}

		movies = append(movies, movie)
	}
	return movies, rows.Err()
}

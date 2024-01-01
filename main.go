package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type Movie struct {
	ImdbID      string  `json:"imdbID"`
	Title       string  `json:"title"`
	Year        int     `json:"year"`
	Rating      float32 `json:"rating"`
	IsSuperHero bool    `json:"isSuperHero"`
}

var movie = []Movie{
	{
		ImdbID:      "tt4154796",
		Title:       "Avengers: Endgame",
		Year:        2019,
		Rating:      8.4,
		IsSuperHero: true,
	},
	{
		ImdbID:      "tt4154756",
		Title:       "Avengers: Infinity War",
		Year:        2018,
		Rating:      8.4,
		IsSuperHero: true,
	},
}

func getAllMovieHandler(c echo.Context) error {
	year := c.QueryParam("year")
	id, err := strconv.Atoi(year)
	if year == "" {
		return c.JSON(http.StatusOK, movie)
	}

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	ms := []Movie{}
	for _, val := range movie {
		if id == val.Year {
			ms = append(ms, val)
		}
	}
	return c.JSON(http.StatusOK, ms)
}

func getAllMovieByHandler(c echo.Context) error {
	id := c.Param("id")
	for idx, v := range movie {
		if v.ImdbID == id {
			return c.JSON(http.StatusOK, movie[idx])
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"message": "not found"})
}

func crateMoviesHandler(c echo.Context) error {
	m := new(Movie)
	err := c.Bind(m)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	/*
		mov := Movie{
			ImdbID:      m.ImdbID,
			Title:       m.Title,
			Year:        m.Year,
			Rating:      m.Rating,
			IsSuperHero: m.IsSuperHero,
		}
	*/

	movie = append(movie, *m)
	return c.JSON(http.StatusCreated, m)
}

func main() {
	e := echo.New()
	port := "2565"
	log.Println("starting... port:", port)

	e.POST("/movies", crateMoviesHandler)

	e.GET("/movies", getAllMovieHandler)
	e.GET("/movies/:id", getAllMovieByHandler)

	log.Fatal(e.Start(":" + port))
}

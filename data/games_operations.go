package data

import (
	"database/sql"
	structure "example/user/Luis/structures"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func InsertGame(db *sql.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		//Create empty game
		var newGame structure.Game

		//BindJSON to bind the received JSON to newGame
		if err := c.BindJSON(&newGame); err != nil {
			panic(err)
		}
		insert, err := db.Prepare("INSERT INTO `games`(`idgame`,`Land1`,`Land2`,`Datum`, `Ergebnis`, `Gruppe`)VALUES(?, ?, ?, ?, ?, ?)")
		if err != nil {
			log.Fatal(err)
		}

		_, insertErr := insert.Exec(&newGame.Id, &newGame.Land1, &newGame.Land2, &newGame.Date, &newGame.Result, &newGame.Group)
		if insertErr != nil {
			log.Fatal(insertErr)
		}

		c.IndentedJSON(200, newGame)
	}
	return gin.HandlerFunc(fn)
}

func GetAllGames(db *sql.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		rows, err := db.Query("select * from games")
		if err != nil {
			fmt.Printf("Error: cant insert into users")
			panic(err.Error())
		}

		var games []structure.Game
		for rows.Next() {
			var game structure.Game
			if userErr := rows.Scan(&game.Id, &game.Land1, &game.Land2, &game.Date, &game.Result, &game.Group); userErr != nil {
				log.Fatal(userErr)
			}
			games = append(games, game)
		}

		c.IndentedJSON(200, games)
	}

	return gin.HandlerFunc(fn)
}

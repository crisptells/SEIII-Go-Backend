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
			fmt.Printf("Error: cant insert into games")
			panic(err.Error())
		}

		var games []structure.Game
		for rows.Next() {
			var game structure.Game
			if gameErr := rows.Scan(&game.Id, &game.Land1, &game.Land2, &game.Date, &game.Result, &game.Group); gameErr != nil {
				log.Fatal(gameErr)
			}
			games = append(games, game)
		}

		c.IndentedJSON(200, games)
	}

	return gin.HandlerFunc(fn)
}

func UpdateGame(db *sql.DB) gin.HandlerFunc  {
	fn := func(c *gin.Context) {
		//Create empty game
		var game structure.Game

		//BindJSON to bind the received JSON to game
		if err := c.BindJSON(&game); err != nil {
			panic(err)
		}
		_, err := db.Query("UPDATE games SET idgames = ?, Land1 = ?, Land2 = ?, Datum = ?, Ergebnis = ?, Gruppe = ? WHERE idgames = ?", game.Id, game.Land1, game.Land2, game.Date, game.Result, game.Group, game.Id)
		if err != nil {
			fmt.Printf("Error: cant update game")
			panic(err.Error())
		}

		c.IndentedJSON(200, game)
	}

	return gin.HandlerFunc(fn)
}

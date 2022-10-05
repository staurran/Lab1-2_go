package app

import (
	"awesomeProject/internal/app/ds"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (a *Application) StartServer() {
	log.Println("Server start up")
	log.Println("Server start up")
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		id := c.Query("id") // получаем из запроса query string
		if id != "" {
			log.Printf("id recived %s\n", id)
			intID, err := strconv.Atoi(id) // пытаемся привести это к чиселке
			if err != nil {                // если не получилось
				log.Printf("cant convert id %v", err)
				err = c.Error(err)
				return
			}
			log.Println(intID, id)
			product, err := a.repo.GetProductByID(uint(intID))
			log.Println(product)
			if err != nil { // если не получилось
				log.Printf("cant get product by id %v", err)
				err = c.Error(err)
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"product_price": product.Price,
			})
			return
		}
		create := c.Query("create")
		log.Printf("create recived %s\n", create)
		create_bool, err := strconv.ParseBool(create)
		if err != nil {
			log.Println("cant convert create")
			return
		}
		if create_bool {
			rand.Seed(time.Now().UnixNano())
			product := ds.Product{Code: "new_product", Price: uint(rand.Intn(999))}
			err = a.repo.CreateProduct(&product)
			if err != nil {
				log.Println("cant create new product")
				return
			}
			log.Println("new row was added")
			c.JSON(http.StatusOK, gin.H{
				"product_price": product.Price,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.LoadHTMLGlob("templates/*")

	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})
	r.GET("/bdlist", func(c *gin.Context) {
		c.HTML(http.StatusOK, "bdlist.tmpl", gin.H{
			"wishlist": []string{"relax", "calm", "safety"},
			"title":    "Wishlist",
		})
	})
	r.GET("/test", func(c *gin.Context) {
		c.HTML(http.StatusOK, "test.tmpl", gin.H{
			"title": "Main website",
			"test":  []string{"a", "b"},
		})
	})

	r.Static("/image", "./resources")

	err := r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		log.Println("Run failed")
	}
	log.Println("Server down")
}

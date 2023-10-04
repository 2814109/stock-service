package api_client

import (
	"app/db"
	"app/types"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type Item struct {
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
}

func Main(ctx context.Context) {
	qiita_user_name := os.Getenv("QIITA_USER_NAME")
	log.Printf("user name is %v", qiita_user_name)
	resp, err := http.Get(fmt.Sprintf("http://qiita.com/api/v2/users/%s/items?page=1&per_page=10", qiita_user_name))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data []Item // nil slice
	// data := make([]Item, 0) のように要素数0の slice としても良い

	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatal(err)
	}
	db := db.OpenDB()

	for _, item := range data {
		fmt.Printf("%s %s\n", item.CreatedAt, item.Title)

		post := &types.Post{Content: item.Title}
		_, err := db.NewInsert().Model(post).Exec(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}
}

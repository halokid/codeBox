package main 

/**
这个推荐系统就是找出两者的异常，然后把各自不同的东西推荐给对方吗？？
**/

import (
  "github.com/muesli/regommend"
  "fmt"
)

func main() {
  books := regommend.Table("books")
  
  booksChrisRead := make(map[interface{}]float64)
  booksChrisRead["1984"] = 5.0
  booksChrisRead["Robinson Crusoe"] = 4.0
  booksChrisRead["Moby-Dick"] = 3.0
  books.Add("Chris", booksChrisRead)

  booksJayRead := make(map[interface{}]float64)
  booksJayRead["1984"] = 5.0
  booksJayRead["Robinson Crusoe"] = 4.0
  booksJayRead["Gulliver's Travels"] = 4.5
  books.Add("Jay", booksJayRead)

 /**  
  recs, _ := books.Recommend("Chris")
  // recs, _ := books.Recommend("Jay")
  for _, rec := range recs {
    fmt.Println("Recommending", rec.Key, "with score:", rec.Distance)
  }
  **/
  
  nbs, _ := books.Neighbors("Chris")
  for _, nb := range nbs {
    fmt.Println("Recommending", nb.Key, "with score:", nb.Distance)
  }
  
  recs, _ := books.Recommend("Chris")
  for _, rec := range recs {
    fmt.Println("Recommending", rec.Key, "with score:", rec.Distance)
  }
  
}
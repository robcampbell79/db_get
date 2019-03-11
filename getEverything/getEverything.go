package getEverything

import(
  "encoding/json"
  "os"
  "database/sql"
  "fmt"
  _ "github.com/go-sql-driver/mysql"
)

type Randonums struct {
  Id int      `json:Id`
  Number int  `json:Number`
}

func Everything() []Randonums {
  fmt.Println("Connecting to db ...")

  db, err := sql.Open("mysql", "samurai:toor@/nippon")

  fmt.Println("Connected to db nippon ...")

  if err != nil {
    panic(err.Error())
  }

  defer db.Close()

  selectAll, err := db.Query("SELECT id, number FROM random_nums")

  if err != nil {
    panic(err.Error())
  }


  defer selectAll.Close()

  result := []Randonums{}

  for selectAll.Next() {
    ran := Randonums{}
    err := selectAll.Scan(&ran.Id, &ran.Number)
    if err != nil {
      fmt.Println(err)
    }
    result = append(result, ran)
  }

  json.NewEncoder(os.Stdout).Encode(result)

  return result

}

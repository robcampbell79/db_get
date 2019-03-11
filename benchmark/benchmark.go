package benchmark

import(
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  "time"
)

func Benchmark() time.Duration {

  start := time.Now()

  db, err := sql.Open("mysql", "samurai:toor@/nippon")

  if err != nil {
    panic(err.Error())
  }

  defer db.Close()

  selectAll, err := db.Query("SELECT id, number FROM random_nums")

  if err != nil {
    panic(err.Error())
  }


  defer selectAll.Close()

  end := time.Since(start)

  return end
}

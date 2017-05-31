package dataAccess

import (
  "log"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

type Test struct {
  Id int
  Value string
}

func ReadAll() []Test {
  db, err := sql.Open("mysql", "root:mariadb@/test")

  rows, err := db.Query("SELECT id, test_value FROM test1")
  if err != nil {
    log.Fatal(err)
  }

  defer rows.Close()

  res := make([]Test, 0)

  for rows.Next() {
    var  (
      value string
      id int
    )

    err := rows.Scan(&id, &value)
    if err != nil {
      log.Fatal(err)
    }

    res = append(res, Test{Id: id, Value: value})

  }

  err = rows.Err()
  if err != nil {
    log.Fatal(err)
  }

  // res := make([]Test, 2)
  // res[0] = Test{Id: 1, Value: "test1"}
  // res[1] = Test{Id: 2, Value: "test2"}



  return res
}

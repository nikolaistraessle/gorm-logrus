# gorm-logrus
Logrus logger for gorm v2

Based on https://github.com/onrik/gorm-logrus

```go
package main

import (
  "gorm.io/gorm"
  "gorm.io/driver/sqlite"
  "github.com/nikolaistraessle/gorm-logrus"
)

func main() {
  db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
    Logger: gormlogrus.New(),
  })
  if err != nil {
    panic("failed to connect database")
  }
}
```

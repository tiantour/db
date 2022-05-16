# db

sqlx wrapper for mysql

# how to use

## conf

conf.toml

    [db]
    IP = "your ip default 127.0.0.1"
    Port = "your port default :3306"
    Uname = "your uname default null"
    Passwd = "your passwd default null"
    DB = "your db default null"

## read

1. List

    db.NewRead().List(query, args)

2. ListNamed

    db.NewRead().ListNamed(query, args)

3. ListStruct

    ```golang
    package main

    import (
        "database/sql"
        "fmt"
        "log"

        "github.com/tiantour/db"
    )

    type Person struct {
        Name string
        Age  int
    }

    func main() {
        var person []*Person
        query := "SELECT name, age FROM `demo` WHERE age = ? LIMIT 10;"
        err := db.NewRead().ListStruct(&person, query, 18)
        if err!=nil{
            log.Panic(err)
        }
        fmt.Println(person)
    }
    ```

4. ListStructNamed

    ```golang
        package main

        import (
            "database/sql"
            "fmt"
            "log"

            "github.com/tiantour/db"
        )

        type Person struct {
            Name string
            Age  int
        }

        func main() {
            args := Person{
                Age: 18,
            }

            var person []*Person
            query := "SELECT name, age FROM `demo` WHERE age = :age LIMIT 10;"
            err := db.NewRead().ListStruct(&person, query, person)
            if err!=nil{
                log.Panic(err)
            }
            fmt.Println(person)
        }
        ```

5. Item

    db.NewRead().Item(query, args)

6. ItemNamed

    db.NewRead().ItemNamed(query, args)

7. ItemStruct

    db.NewRead().ItemStruct(dest, query, args)

8. ItemStructNamed

    db.NewRead().ItemStructNamed(dest, query, args)

# write

1. List

    db.NewWrite().List(query, args)

2. ListNamed

    db.NewWrite().List(query, args)

3. Item

    db.NewWrite().Item(query, args)

4. ItemNamed

    db.NewWrite().ItemNamed(query, args)

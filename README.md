# db

sqlx wrapper for mysql

## read

1. List

    db.NewRead().List(query, args)

2. ListNamed

    db.NewRead().ListNamed(query, args)

3. ListStruct

4. ListStructNamed

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

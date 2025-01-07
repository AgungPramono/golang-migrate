## GOLANG MIGRATION ##
1.  Contoh perintah menginstall golang migration tools

```shell
go install -tags 'postgres,mysql,mongodb' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

2.  Perintah untuk membuat database migration

```sql
migrate create -ext sql -dir db/migrations create_table_category
```
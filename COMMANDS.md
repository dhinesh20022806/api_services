# migration commands

```bash
export DATABASE_URL=<DATABASE_URL>
```

create migration

```bash
migrate create -ext sql -dir migrations -seq < FILENAME >
```

run a migrations

```bash
migrate -path migrations -database "$DATABASE_URL" up
```

down a migrations

```bash
migrate -path migrations -database "$DATABASE_URL" down 1
```

get the version

```bash
migrate -path migrations -database "$DATABASE_URL" version
```

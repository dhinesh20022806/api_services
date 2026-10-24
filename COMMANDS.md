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

```bash

curl -X GET  "http://localhost:8080/my-bucket/01_Verb_Tenses+(1).pdf"  -H "Authorization: Bearer $JWT_TOKEN"  -o sample.pdf

```

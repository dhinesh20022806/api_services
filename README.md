### This is S3 rest api services

```bash
curl -X POST http://localhost:8080/bucket \
  -H "Authorization: Bearer $JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"name": "my-new-bucket"}'

  echo
```

```bash
curl -X PUT http://localhost:8080/my-new-bucket/HelloWorld.mp4 \
  -H "Authorization: Bearer $JWT_TOKEN" \
  -F "file=@./HelloWorld.mp4"

  echo
```

```bash
curl -X GET http://localhost:8080/my-new-bucket/HelloWorld.mp4 \
  -H "Authorization: Bearer $JWT_TOKEN" \
  -o s3objectHelloWorld.mp4

  echo
```

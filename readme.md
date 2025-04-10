## Pokedex api

- THIS IS POKEDEX API

## Learnt

- In order to export, need to create a folder not a file
- **const** keyword supports only "basic" types
- **make** keyword for initializing empty object. Cannot directly assign it
- If struct type doesn't contain props, even if api return with all values, returns struct type without the one forgot to add
  - i.g)
    struct {Name string `json:"name"`}
    api returned type {name : 'toni', height:180}
    unmarshaled data will be {name : 'toni'} only

### Dynamo

- Schemaless DB

  - Don't need to specify each fields such as "cries", "sprites" etc...

- AttributeDefinitions

  - The AttributeDefinitions field specifies the data types of the attributes used in the table’s key schema or secondary indexes.
    - You only need to define attributes that are part of the key schema or used in indexes.

- The KeySchema defines the primary key for the table, which consists of:

  - A partition key (also called the “HASH” key).
  - An optional sort key (also called the “RANGE” key).

- AWS.string type

  - AttributeDefinition

    - "N" stands for number
    - "S" stands for string

  - KeySchemaElement
    - "HASH" is for partition key
    - "RANGE" is for sort key

### Main package

- Can have multiple main packages in different folders, as long as each folder represents a separate program. For example:

### AWS

- **~/.aws/credentials** contains account info

```js
[default]
aws_access_key_id = YOUR_DEFAULT_ACCOUNT_ACCESS_KEY
aws_secret_access_key = YOUR_DEFAULT_ACCOUNT_SECRET_KEY

[account1]
aws_access_key_id = ACCOUNT1_ACCESS_KEY
aws_secret_access_key = ACCOUNT1_SECRET_KEY

[account2]
aws_access_key_id = ACCOUNT2_ACCESS_KEY
aws_secret_access_key = ACCOUNT2_SECRET_KEY
```

- **~/.aws/config** specify regions for each profiles

- Use a specific profile
  ```js
      aws s3 ls --profile account1
  ```

### Query vs Scan

- https://dynobase.dev/dynamodb-scan-vs-query/

### BatchWrite size

- Apparently, the size of each writes should be less than 400KB

```
Error writing batch: ValidationException: Item size has exceeded the maximum allowed size
```

### Debugging configuration

- **${workspaceFolder}**: Represents the root folder of your workspace—the folder you opened in VS Code.

```
"program": "${workspaceFolder}",

```

### Docker

- To build image with tag

```
  docker build --tag pokedex-api .
```

- To run container

```
docker run -p 8080:8080 --network=pokedex-network --name pokedex-backend pokedex-api
```

# Finally made it work

1. Initial Docker Compose Setup

Issue:

Docker Compose showed warnings about the deprecated version field.

Backend failed to start due to: exec: "docker-pokedex-api": executable file not found in $PATH

What we did:

Removed the version: key (optional but cleaner).

Investigated the Go backend Dockerfile.

Realized that the binary docker-pokedex-api wasn’t executable or not present.

Updated the Dockerfile with:

RUN CGO_ENABLED=0 GOOS=linux go build -o docker-pokedex-api
RUN chmod +x docker-pokedex-api

✅ 2. API Routing & Frontend Integration

Issue:

Frontend could not access backend APIs: ERR_NAME_NOT_RESOLVED when trying to hit http://backend:8080.

What we did:

Confirmed that in Docker Compose, the backend service name is pokedex-api, not backend.

Set VITE_URL_PATH to "/pokemon" and used proxying via NGINX.

Updated nginx.conf in frontend to route:

location /pokemon/ {
proxy_pass http://pokedex-api:8080;
}

✅ 3. Docker Build Errors from Alpine/Golang Images

Issue:

Couldn’t pull alpine:latest or golang:1.24-alpine due to registry errors.

What we did:

Added a Docker Engine config with Google mirror:

{
"registry-mirrors": ["https://mirror.gcr.io"]
}

Restarted Docker and pulled images successfully.

✅ 4. Go Backend: .env & AWS Credential Issues

Issue:

App failed to load .env file inside container

Then, failed to connect to DynamoDB with:

NoCredentialProviders

UnrecognizedClientException

What we did:

Ensured .env was copied at build time if needed, but preferred using Compose environment variables instead.

Removed the local DynamoDB DYNAMO_ENDPOINT when switching to AWS DynamoDB Cloud.

Passed real AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY via docker-compose.yml.

✅ 5. DynamoDB Table Errors

Issue:

Backend connected to AWS DynamoDB, but got:
ResourceNotFoundException: Cannot do operations on a non-existent table

What we did:

Realized that the table existed only in AWS Console, not in DynamoDB Local.

Switched to using real DynamoDB cloud.

Verified table existence, and used correct region and credentials.

✅ 6. Frontend Not Showing / Port Refused

Issue:

Frontend wasn't visible on localhost:7777

Got browser error: This site can’t be reached

What we did:

Confirmed container was exposing port: 0.0.0.0:7777->80

Confirmed NGINX was running and logs were clean

Found that nginx.conf had invalid characters (``` backticks)

Fixed syntax and ensured it routes correctly to backend

Rebuilt image and finally got frontend working

✅ 7. Final Working Setup

Go backend running at: localhost:8080

React frontend served via NGINX at: localhost:7777

API calls from React to /pokemon/... are successfully proxied to backend

Backend connects to real AWS DynamoDB

Everything managed with Docker Compose

🔗 What You Learned Today

Multi-stage Docker builds (Go + Node/React)

Docker Compose networking and service linking

NGINX config for single-page apps + proxying APIs

AWS SDK integration with proper credentials in containers

Debugging container logs like a champ

Fixing subtle build issues, permissions, and port binding

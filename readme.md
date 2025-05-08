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

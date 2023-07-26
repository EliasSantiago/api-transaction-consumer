# API Transaction Consumer

## About the project

- An api for consumer queue

## How to run the api?

**Step 1: Clone the project, run the following commands:**

**Step 2: Create file config.json**
{
  "database": {
    "url": "://postgres:postgres@172.21.0.2:5432/postgres"
  },
  "server": {
    "port": "3001"
  },
  "secret_key": "HsZwxbcwuK/yZBbNeUjTAA=="
}

To get the database ip, use...
docker ps -> check container id postgres, then...
docker inspect container_id
Exemple; "IPAddress": "172.21.0.2"

**Step 3: Run docker(at the root of the project)**
- docker-composer up -d

**Step 4: go mod tidy(at the root of the project)**
- go mod tidy

**Step 5: run consumer**
- go run adapter/cmd/main.go
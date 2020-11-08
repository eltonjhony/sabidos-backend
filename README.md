# sabidos-backend
Notes:
docker rm $(docker ps -a -q)
docker run --name sabidos-db  -p 27017:27017    -e MONGO_INITDB_DATABASE=sabidos    mongo

Golang backend for sabidos

 - docker-compose up --build --force-recreate

 - POST: http://localhot:8080/v1/account

 {
        "name": "Hulk",
        "nickname": "Smash",
        "avatar": {
            "defaultAvatarId": ""
        },
        "reputation": {
            "level": "",
            "stars": ""
        },
        "totalAnswered": "",
        "totalHits": ""
    
}

- GET: http://localhot:8080/v1/account/1000

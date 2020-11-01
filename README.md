# sabidos-backend
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

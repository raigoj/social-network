##TO RUN
docker-compose up

##TO ACCESS
http://localhost:8080

##TO STOP
docker-compose down

##Run without dockercompose

docker build -t rtforum .
docker run -p 8009:8009 -d --name rtforum rtforum

docker build -t vue-app .
docker run -p 8080:8080 -d --name vue-app vue-app

##Stop without dockercompose

docker stop rtforum
docker stop vue-app

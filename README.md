# msdb5
The MSDB5 app

# manage via docker - ITA
docker build -t msdb5-it .
docker run -p 4000:8080 msdb5-it
or
docker run -d -p 4000:8080 msdb5-it

# manage via docker - ENG
docker build -f Dockerfile-en -t msdb5-en .
docker run -d -p 4000:8080 msdb5-en

# list containers and stop them
docker container ls -all
docker stop <container-id>
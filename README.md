# channels and go routines and encryption and such

# docker

<!-- to build -->
docker build -t peterjbishop/go-crypt:latest .
docker run -p 8080:8080 peterjbishop/go-crypt:latest
docker-compose up --build

# ngrok testing

<!-- to test live  -->
ngrok http 8080  
- or launch ngork through docker
FROM golang:latest

ENV REPO_URL =github.com/OleksandrVakhno/GitHubMicroservice/

ENV GO_PATH = /app

ENV APP_PATH=$GOPATH/src/$REPO_URL

ENV WORKPATH=$APP_PATH/src
COPY src $WORKPATH
WORKDIR $WORKPATH

RUN go build -o repos-api .

#Expose port 8080
EXPOSE 8080
CMD ["./repos-api"]
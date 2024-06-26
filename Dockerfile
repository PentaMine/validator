FROM code_runner

WORKDIR /usr/src/app

#COPY setup.sh .
#RUN sh setup.sh

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
# RUN go build -v -o /usr/local/bin/app ./...

EXPOSE 6969

CMD ["go", "run", "."]
FROM golang:alpine
WORKDIR '/app'
COPY *.go ./
COPY *.mod ./
COPY ./public/* ./public/
EXPOSE 8081
CMD ["go", "run", "."]
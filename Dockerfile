FROM golang
COPY . .
EXPOSE 8888
CMD ["go","run","main.go"]
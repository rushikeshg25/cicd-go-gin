ec
cd cicd-go-gin
git pull origin master
go mod tidy
go build -o main
./main
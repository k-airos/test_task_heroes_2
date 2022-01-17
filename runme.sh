if [ `whoami` != root ]; then
    echo Please run this script as root or using sudo
    exit
fi

docker run --name testmongodb -d -p 27017:27017 mongo
cd marvel
go build ./cmd/api/main.go
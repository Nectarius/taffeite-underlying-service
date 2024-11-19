# taffeite-underlying-service
taffeite underlying service

docker build -t taffeite .

docker run -p 2560:2560 taffeite
docker run -p 2560:2560 --rm -ti taffeite
docker run -d -p 2560:2560 taffeite

docker build -t taffeite -f Dockerfile ./taffeite

ssh-keygen -t rsa -b 4096 -C "nectarius@outlook.com"
xclip -sel clip < ~/.ssh/taffeite_key.pub

git push git@github.com:Nectarius/taffeite-underlying-service.git

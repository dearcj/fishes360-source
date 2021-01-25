containers=`docker ps -a -q`
docker stop $containers
docker container prune --force
#docker rm $(docker ps -a -q)
#docker rmi $(docker images -q)
git pull --force
cp -a /wbserv/src/github.com/dearcj/golangproj/vendor/. /wbserv/src/
docker build -t game .
docker run --publish 80:80 --name game --rm --net=host game
#~/Prometheus/prometheus-2.1.0.linux-amd64/prometheus
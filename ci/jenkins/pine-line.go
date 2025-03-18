package jenkins

/*
cd /var/www/bit-web/deployment/production/bit-web

sudo chmod -R 777 .
git reset --hard
git clean -fd .
git checkout deployment/production
git fetch
git pull
git status
cd ..
sudo chmod -R 777 .

docker-compose up -d --build bit-web

docker image prune --force
docker container prune --force
*/

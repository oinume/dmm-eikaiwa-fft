# lekcije
Follow your favorite teachers in DMM Eikaiwa and receive notification when favorite teachers open lessons.

# Install dependencies

```bash
brew cask install dockertoolbox
make setup
npm install
```

## For developers
```bash
brew install chromedriver
```

# Develop on your local machine

### Run MySQL server

```
eval "$(docker-machine env default)"
docker-compose up
```

### Run server
```
reflex -R node_modules -R vendor -r '\.go$' -s make serve
```

### Run frontend
```
npm start
```

### Connect to mysql on Docker

```
mysql -uroot -proot -h $(docker-machine ip default) -P 13306 lekcije
```

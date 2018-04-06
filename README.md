# Vuetify-Admin

## Prepare

### Install Node

```
curl -o- https://raw.githubusercontent.com/creationix/nvm/v0.31.1/install.sh | bash

source ~/.bashrc

nvm install v9.8.0
```

### Install yarn

```
curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | sudo apt-key add -

echo "deb https://dl.yarnpkg.com/debian/ stable main" | sudo tee /etc/apt/sources.list.d/yarn.list

sudo apt-get update && sudo apt-get install yarn
```

### Install go 

```
wget https://dl.google.com/go/go1.8.3.linux-amd64.tar.gz

tar -C /usr/local -xzf go1.8.3.linux-amd64.tar.gz

export GOPATH=/usr/local/go/bin
```

### Install glide

```
curl https://glide.sh/get | sh
```

### Config Port

docker-compose.yml

```
 xxx-admin:

    ...

    ports:
      - "<your web port>:80"

    ...
    
  
 xxx-admin-api:
    
    ...

    ports:
      - "<your api port>:8000"
    
    ...

```

### Config Port

xxx-admin/src/config/index.js.prod

```
export default {
  apiBaseURL: 'http://localhost:8000/console', // your api base url
};
```

## Build

```
sh ./build.sh
```
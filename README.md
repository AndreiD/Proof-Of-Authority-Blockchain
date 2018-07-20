# parity-poa-playground

Setup your first Parity PoA network with 3 authorities and 3 members.

### Why Use This :fire:

- Easy way to get started with POA
- Full Byzantium-fork compatibility & experimental support for WebAssembly

### Setup

0. Install [docker](https://docs.docker.com/engine/installation/) and [docker-compose](https://docs.docker.com/compose/install/)

Initial Setup Example:

~~~~
# Need curl before....
sudo apt install -y curl

# Git
sudo apt-add-repository ppa:git-core/ppa

# Add Docker repository key to APT keychain
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
# Update where APT will search for Docker Packages
echo "deb [arch=amd64] https://download.docker.com/linux/ubuntu ${CODENAME} stable" | \
    sudo tee /etc/apt/sources.list.d/docker.list
sudo apt-get update

sudo apt install -y git

# Configure git
git config --global core.autocrlf false
git config --global core.longpaths true

# Verifies APT is pulling from the correct Repository
sudo apt-cache policy docker-ce

sudo apt-get -y install docker-ce
sudo usermod -aG docker $(whoami)
sudo curl -L "https://github.com/docker/compose/releases/download/1.13.0/docker-compose-$(uname -s)-$(uname -m)" \
    -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
~~~~

1. Run `git clone https://github.com/orbita-center/parity-poa-playground.git && cd parity-poa-playground`
2. Run `docker-compose up -d`

#### Quick docker-compose commands cheatsheet:

- You can stop with: docker-compose kill, remove them with docker-compose rm
- Restart them with: docker-compose up (add -d for detach)
- List them with: docker ps
- Check what is started with: docker inspect  -f "{{.Name}} {{.Config.Cmd}}" $(docker ps -a -q)

You will have an account unlocked "0x00bd138abd70e2f00903268f3db08f2d25677c9e"


### Access the [ethstats](https://github.com/cubedro/eth-netstats) dashboard.
A nice dashboard is already configured and connected with all the nodes.
Find it at [http://127.0.0.1:3001](http://127.0.0.1:3001).

### Accounts
There is already an account with an empty password that has enough ether:

```
0x6B0c56d1Ad5144b4d37fa6e27DC9afd5C2435c3B
```

And another who is broke:
```
0x00E3d1Aa965aAfd61217635E5f99f7c1e567978f
```

You may also want to change the list of prefunded accounts in `parity/config/chain.json`.

Add JSON-formatted ethereum accounts to `parity/keys`.


### Access JSON RPC 
Talk to JSON RPC at [http://127.0.0.1:8545](http://127.0.0.1:8545) with your favorite client.

Be kind and send the poor an ether!

```
curl --data '{"jsonrpc":"2.0","method":"personal_sendTransaction","params":[{"from":"0x6B0c56d1Ad5144b4d37fa6e27DC9afd5C2435c3B","to":"0x00E3d1Aa965aAfd61217635E5f99f7c1e567978f","value":"0xde0b6b3a7640000"}, ""],"id":0}' -H "Content-Type: application/json" -X POST localhost:8545
```

### Extras:

#### Your own blockchain explorer so you can quickly check transaction hashes, accounts etc.

```
git clone https://github.com/gobitfly/etherchain-light --recursive
Install all dependencies: npm install
Rename config.js.example into config.js and adjust the file to your local environment
Start the explorer: npm start
Browse to http://localhost:3000
```

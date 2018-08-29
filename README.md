# Proof of Authority Blockchain 

### This is *AURA* Parity Version. 

Setup your first Parity PoA network with 3 authorities and 3 members.


### Why Use This :fire:

- Beginner friendly! (wow, that's rare in blockchain!)
- Easy way to get started with POA
- Full Byzantium-fork compatibility & experimental support for WebAssembly
- A block explorer provided
- Netstats provided
- Faucet provided 

### Setup

You need git, docker, nodejs etc.
Install [docker](https://docs.docker.com/engine/installation/) and [docker-compose](https://docs.docker.com/compose/install/)

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

1. Run `git clone git@github.com:AndreiD/Proof-Of-Authority-Blockchain.git && cd proof-of-authority-blockchain`
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

Example trasfering 1 ether from one account to another:
```
curl --data '{"jsonrpc":"2.0","method":"personal_sendTransaction","params":[{"from":"0x6B0c56d1Ad5144b4d37fa6e27DC9afd5C2435c3B","to":"0x00E3d1Aa965aAfd61217635E5f99f7c1e567978f","value":"0xde0b6b3a7640000"}, ""],"id":0}' -H "Content-Type: application/json" -X POST localhost:8545
```

### Extras:

#### Your own blockchain explorer so you can quickly check transaction hashes, accounts etc.

- Setup a nodejs & npm environment
- Install the latest version of the Parity Ethereum client
- Start parity using the following options: parity --chain=<yourchain> --tracing=on --fat-db=on --pruning=archive
- Clone this repository to your local machine: git clone https://github.com/poanetwork/chain-explorer --recursive (Make sure to include - recursive in order to fetch the solc-bin git submodule)
- Install all dependencies: npm install
- Rename config.js.example into config.js and adjust the file to your local environment
- Start the explorer: npm start
- Browse to http://server_ip:3000

### About

This docker compose project originally started on https://github.com/orbita-center/parity-poa-playground, after making some pull requests and explaining why they should include some changes, which were never approved, I decided to make my own repo. 

### Thanks

Thanks to [Grzegorz Bytniewski](https://github.com/bytniak) for his suggestions



### License

    DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE 
                    Version 2, December 2004 

 Copyright (C) 2004 Sam Hocevar <sam@hocevar.net> 

 Everyone is permitted to copy and distribute verbatim or modified 
 copies of this license document, and changing it is allowed as long 
 as the name is changed. 

            DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE 
   TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION 

  0. You just DO WHAT THE FUCK YOU WANT TO.

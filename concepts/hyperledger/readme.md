
# 🚀 Install & Run Hyperledger Fabric on Mac (with FabCar)

---

## 🧰 1. Install Prerequisites

### ✅ Docker Desktop (latest)

* 📥 [Download Docker Desktop](https://www.docker.com/products/docker-desktop)
* After installing, verify it’s running:

```bash
docker --version
docker-compose --version
```

---

### ✅ Homebrew (if not already installed)

```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```

---

### ✅ Install Required CLI Tools

```bash
brew install curl git go
brew install node  # if you want Node.js SDK later
```

---

## 📦 2. Download Hyperledger Fabric Samples

### ✅ Clone the official Fabric samples repository

```bash
git clone https://github.com/hyperledger/fabric-samples.git
cd fabric-samples
```

---

## 📥 3. Use the Official Install Script

### ✅ Step 1: Download the install script

```bash
curl -sSLO https://raw.githubusercontent.com/hyperledger/fabric/main/scripts/install-fabric.sh
chmod +x install-fabric.sh
```

---

### ✅ Step 2: Run the script

Install Docker images, binaries, and samples:

```bash
./install-fabric.sh docker samples binary
```

**What this does:**

* Downloads Fabric binaries (peer, orderer, cryptogen, configtxgen, etc.)
* Pulls Fabric Docker images
* Sets up the `fabric-samples` directory for development

---

### ✅ Step 3: Add Fabric binaries to your `PATH`

```bash
export PATH="$PATH:$(pwd)/bin"
```

Add this line to your `~/.zshrc` or `~/.bashrc` for persistence.

---

## 🧪 4. Start the Test Network

### ✅ Navigate to the test network

```bash
cd test-network
```

### ✅ Bring up the network

```bash
./network.sh up
```

This will start:

* `orderer.example.com` (port 7050)
* `peer0.org1.example.com` (port 7051)
* `peer0.org2.example.com` (port 9051)

---

## 📦 5. Deploy the FabCar Chaincode

### ✅ Deploy FabCar

```bash
./network.sh deployCC -ccn fabcar -ccp ../chaincode/fabcar/go -ccl go
```

**Flags:**

* `-ccn` → chaincode name (`fabcar`)
* `-ccp` → chaincode path
* `-ccl` → language (`go`, `node`, or `java`)

---

## 🔍 6. Interact with the Network

> ⚠️ Make sure you are inside the CLI container or have exported env vars for `CORE_PEER_*` if running directly on your terminal.

---

### ✅ Query all cars

```bash
peer chaincode query -C mychannel -n fabcar -c '{"Args":["queryAllCars"]}'
```

---

### ✅ Create a new car

```bash
peer chaincode invoke -o localhost:7050 \
  --ordererTLSHostnameOverride orderer.example.com \
  --tls \
  --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem \
  -C mychannel -n fabcar \
  -c '{"Args":["createCar","CAR10","Honda","Civic","blue","Alice"]}'
```

---

### ✅ Verify again

```bash
peer chaincode query -C mychannel -n fabcar -c '{"Args":["queryAllCars"]}'
```

Look for:

```json
{"Key":"CAR10", "Record":{"make":"Honda", "model":"Civic", "colour":"blue", "owner":"Alice"}}
```

---

## ✅ Summary

| Step               | Command / Tool                        |
| ------------------ | ------------------------------------- |
| Install Prereqs    | Docker, Homebrew, curl, git, go, node |
| Download Fabric    | `git clone` & `install-fabric.sh`     |
| Add to PATH        | `export PATH=...`                     |
| Start network      | `./network.sh up`                     |
| Deploy chaincode   | `./network.sh deployCC ...`           |
| Query chaincode    | `peer chaincode query ...`            |
| Invoke transaction | `peer chaincode invoke ...`           |

---

🎉 **Done!** You now have a full **Fabric network running locally** with the **FabCar chaincode** deployed and tested on your Mac 🚀



# Hyperledger Fabric — Reference Notes

## What is Hyperledger Fabric?
- **Permissioned blockchain platform** for business networks (private/consortium).
- **No mining** (no PoW/PoS). Instead:
  - Transactions validated via **endorsement policies**.
  - Ordered into blocks by an **ordering service**.
- **Configurable privacy**: channels and private data collections.
- **Pluggable components**: endorsement, ordering, consensus, databases (CouchDB/LevelDB).

---

## High-level Properties
- **Permissioned** → only participants with valid identities (via MSP/CA) can join.
- **Private data** → visibility restricted using channels or collections.
- **No miners/rewards** → ordering service manages consensus.
- **Modular** → pluggable consensus, databases, and membership services.

---

## Core Terminology
- **Asset** → digital representation of value (e.g., car part, shipment).
- **Network** → collection of organizations, peers, orderers, channels.
- **Organization (Org)** → company/entity with peers and identities.
- **Peer** → maintains ledger, endorses & commits transactions.
- **Orderer** → orders transactions into blocks (consensus).
- **Chaincode** → smart contract / business logic.
- **Ledger** → immutable blockchain + current world state (DB).
- **MSP (Membership Service Provider)** → manages identities & crypto.
- **Channel** → private subnet for a group of orgs.
- **World state DB** → current state (LevelDB = key-value, CouchDB = JSON/rich queries).

---

## Architecture (Simplified)

```mermaid
flowchart TD
    subgraph App["Client Application"]
        UI["UI / Web App"]
        SDK["Fabric SDK (Node / Go / Java)"]
        UI --> SDK
    end

    subgraph Org1["Organization 1"]
        Peer1["Peer (endorses, commits)"]
        CC1["Chaincode (Smart Contract)"]
        DB1["World State DB"]
        CA1["CA (Certificate Authority)"]

        Peer1 --> DB1
        Peer1 --> CC1
    end

    subgraph Org2["Organization 2"]
        Peer2["Peer (endorses, commits)"]
        CC2["Chaincode (Smart Contract)"]
        DB2["World State DB"]
        CA2["CA (Certificate Authority)"]

        Peer2 --> DB2
        Peer2 --> CC2
    end

    Orderer["Ordering Service"]

    SDK --> Peer1
    SDK --> Peer2
    Peer1 <--> Orderer
    Peer2 <--> Orderer

````

```mermaid
flowchart TD
    subgraph Org1["Organization 1"]
        Peer1["Peer"]
        DB1["World State DB (Replica)"]
        Peer1 --> DB1
    end

    subgraph Org2["Organization 2"]
        Peer2["Peer"]
        DB2["World State DB (Replica)"]
        Peer2 --> DB2
    end

    Orderer["Ordering Service"]

    Peer1 <--> Orderer
    Peer2 <--> Orderer
```

* **Chaincode**: deployed to peers, executes business logic.
* **SDK/API**: submits proposals → gets endorsements → sends to orderer → block created → peers validate + commit.

---

## Supply Chain Example (Car Parts)

Flow of assets:

```
Car Part → Collector → Assembler → Producer → Seller → Buyer
```

* Each step updates the **asset** on the ledger.
* Provenance is immutable and traceable.
* Participants only see allowed data (channels/private data).
* Replaces siloed databases → shared ledger across orgs.

---

## Deployment & Infrastructure

* Each **Org** runs:

  * One or more **peers** (Docker containers).
  * Optionally a **CA** (certificate authority).
* **Orderers**: separate containers.
* **Chaincode**: runs in its own container per peer.
* Deployment options:

  * Docker Compose (dev/test).
  * Kubernetes (prod).
  * Managed services.

---

## Chaincode (Smart Contract)

* **Implements business logic** (CRUD on assets).
* Written in **Go, Node.js, or Java**.
* Runs in isolated container.
* Lifecycle:

  1. Package
  2. Install on peers
  3. Approve for org
  4. Commit to channel
* **Endorsement policies**: define which orgs must sign off a transaction.

---

## Steps to Build a PoC

1. **Plan network**: orgs, peers, channels, policies.
2. **Setup MSP & CA**: generate identities/certs.
3. **Create network**: bring up peers & orderers, create channels.
4. **Develop chaincode**: define assets + transaction functions.
5. **Deploy chaincode**: package, install, approve, commit.
6. **Write client app**: use Fabric SDK (Node/Go/Java) to interact.
7. **Test**: asset creation, transfer, provenance queries.
8. **Add privacy**: channels or private data collections.
9. **Monitor & maintain**: logs, metrics, snapshots.

---

## Common Misconceptions

* ❌ **“Fabric has no consensus”** → It does: ordering service + endorsement policies.
* ❌ **“All data is public”** → Channels & private data collections restrict access.
* ❌ **“Fabric replaces DBs entirely”** → Ledger is shared, but each peer maintains a local world state DB.

---

## Practical Tips

* Start with **Fabric test-network** from official samples.
* Use **CouchDB** for rich JSON queries.
* Keep **large files off-chain**; store only hashes/refs.
* Balance **endorsement policies** for trust vs performance.
* Carefully design **asset models** and event handling.



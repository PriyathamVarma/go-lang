# 📘 Hyperledger Fabric Notes for Fabcar

## 6.1 Introduction

Hyperledger Fabric is one of the leading platforms under the Hyperledger framework (initiated by the Linux Foundation) to create permissioned blockchain applications.

It provides:

* A solution for enterprise blockchain use cases.
* Modular and extensible architecture.
* Support for modeling **car ownership tracking** as a case study.
* Tools for application development and environment setup.

> 💡 **Smart Fact:**
> The Linux Foundation is a non-profit technology consortium founded in 2000. It promotes Linux and open-source projects globally.
> Source: [linuxfoundation.org](https://www.linuxfoundation.org)

---

## 6.2 Use Case – Car Ownership Tracking

The use case demonstrates a blockchain network that tracks car ownership from **manufacturing → sale to customer**.

### Simplified stakeholders:

* Manufacturer
* Dealer
* Customer

### Key questions:

1. **Does this involve multiple participants sharing assets/data?**
   ✅ Yes. Car is shared across manufacturer, dealer, and customer.

2. **Are intermediaries involved?**
   ✅ Yes. Dealers and agents are intermediaries (though blockchain can reduce reliance on them).

3. **Does the solution require shared write access?**
   ✅ Yes. Car ownership data must be updated across participants.

4. **Can the solution work without delete?**
   ✅ Yes. Blockchain is append-only. Ownership changes are tracked, not deleted.

5. **Is it required to store non-transactional data?**
   ❌ No. Only relevant car data is digitized on the blockchain. Non-essential data can stay off-chain.

6. **Does the solution require small group control?**
   ✅ Yes. Controlled by participating parties.

7. **Does it involve high-performance transactions?**
   ✅ Yes. Transactions span multiple participants, often across days/weeks.

> 💡 **Smart Fact:**
> Lamborghini announced in Nov 2019 that it uses blockchain to track cars in the secondary market to ensure authenticity of cars and parts.

---

## 6.3 Hyperledger Fabric

Hyperledger Fabric is an **open-source distributed ledger technology (DLT) platform** designed for enterprise use.

### Key features:

* **Permissioned** (unlike Bitcoin/Ethereum public chains).
* Supports **different consensus mechanisms**.
* Enables storing data in **multiple formats**.
* Allows private communication channels for extra privacy.
* Does **not** have built-in tokens (like Ethereum).

### Fabric Architecture Components (Figure 6.1)

* **CA (Certificate Authority)**

  * Registration of identities
  * Manage certificates

* **Peer**

  * Endorses transactions
  * Simulates and commits transactions

* **Orderer**

  * Verifies consensus
  * Creates blocks

* **Ledger**

  * Blockchain + World state

* **Channels**

  * Private subsets for smart contracts
  * Peers can have multiple channels

* **Smart Contracts (Chaincode)**

  * Chaincode compilation
  * Endorsement & read/write logic

* **Other Concepts**

  * Assets
  * Transactions
  * Endorsement policies
  * Gossip protocol


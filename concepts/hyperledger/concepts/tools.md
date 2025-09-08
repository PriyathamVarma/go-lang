# Hyperledger Ecosystem

Hyperledger is not a single blockchain but an **umbrella project** under the Linux Foundation.  
It consists of **frameworks** (blockchain platforms) and **tools** (utilities that support networks).

---

## 🔹 Frameworks
Frameworks are blockchain platforms tailored for different use cases.

- **Hyperledger Fabric**  
  - Modular, permissioned blockchain platform.  
  - Supports pluggable consensus, private channels, and chaincode (smart contracts in Go, Node.js, Java).  
  - Most widely adopted in supply chain, trade finance, healthcare.

- **Hyperledger Indy**  
  - Focused on **decentralized identity (DID)**.  
  - Used for verifiable credentials, self-sovereign identity (SSI).  
  - Provides privacy-preserving identity management.

- **Hyperledger Iroha**  
  - Simple, lightweight blockchain framework.  
  - Mobile-friendly and easy to integrate with existing systems.  
  - Often used in IoT, digital asset management, and identity.

- **Hyperledger Sawtooth**  
  - Modular blockchain with support for different consensus mechanisms.  
  - Introduced **PoET (Proof of Elapsed Time)** consensus.  
  - Good for enterprise supply chains and IoT.

- **Hyperledger Burrow**  
  - Permissioned Ethereum smart contract engine.  
  - Supports EVM (Ethereum Virtual Machine).  
  - Used when Ethereum-style contracts are needed in a permissioned setting.

---

## 🔹 Tools
Tools support the development, deployment, monitoring, and operation of Hyperledger frameworks.

- **Hyperledger Composer**  
  - (⚠️ Deprecated) Framework for rapidly building blockchain applications on Fabric.  
  - Provided modeling language + REST server.  
  - Simplified Fabric app dev but now replaced by direct SDK usage.

- **Hyperledger Explorer**  
  - Web-based tool to **view blockchain data** (blocks, transactions, smart contracts, channels, network).  
  - Provides dashboards for admins and users.

- **Hyperledger Cello**  
  - Blockchain-as-a-Service (BaaS) tool.  
  - Automates deployment and management of blockchains on cloud or bare metal.  
  - Provides a console for operators to provision/manage nodes.

---

## ✅ Quick Summary
- **Frameworks** = blockchain platforms (Fabric, Indy, Iroha, Sawtooth, Burrow).  
- **Tools** = supporting utilities (Composer, Explorer, Cello).  
- Together, they form the **Hyperledger ecosystem** for enterprise blockchain development.

# Inventory Management Service (IMS)

<p align="center">
  <img src="resources/OMS-IMS%20WorkFlow%20Diagram.png" alt="OMS-IMS Workflow Diagram" width="100%">
</p>

---

## 🧠 Overview

IMS is a microservice that communicates with the Order Management System (OMS). It manages and stores inventory, product, seller, and tenant data using PostgreSQL. Redis is used for caching valid `hub_id` and `sku_id` values to optimize validation performance.

---

## 🛠 Technologies Used

- **PostgreSQL** – for storing inventory-related data (products, sellers, tenants, etc.)
- **Redis** – for caching valid `hub_id` and `sku_id` values

---

## 📁 Project Directory Structure

```
IMS/
├── configs/
│   └── config.yaml
├── init/
│   └── initialize.go
├── migrations/
│   ├── 1_schema.down.sql
│   ├── 1_schema.up.sql
│   └── migration.go
├── mycontext/
│   └── mycontext.go
├── pkg/
│   ├── controllers/
│   ├── integrations/
│   ├── db/
│   ├── redis/
│   ├── middlewares/
│   ├── models/
│   └── routes/
├── resources/
│   └── OMS-IMS WorkFlow Diagram.png
├── go.mod
├── go.sum
└── main.go
```

---

## 🔌 API Overview

### 🔹 Hub + SKU APIs
- CRUD operations for hubs and SKUs
- Filter SKUs by tenant, seller, or SKU codes

### 🔸 Inventory APIs
- Atomic upsert for inventory updates
- View inventory based on hub and SKUs
- Missing entries default to quantity `0`

---

## 🔁 Interservice Communication

### ✅ Validation Workflow

- The **CSV Processor** in OMS checks Redis and PostgreSQL via IMS to validate `hub_id` and `sku_id`.
- A **Kafka Consumer** in OMS calls the `checkInventory` API of IMS to:
  - Validate if the current `sku_id` and `hub_id` have sufficient quantity to fulfill the order.
  - If sufficient, the quantity is reduced in PostgreSQL.
  - IMS responds with `status: true`, indicating the inventory can fulfill the order.
  - OMS then updates the order state to `NEW_ORDER`.

---

## ▶️ Run Locally

```bash
$env:CONFIG_SOURCE = "local"
go run main.go
```

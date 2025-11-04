# ⚡ Go Backend Service – Gin, GORM, Redis, ELK, Prometheus

A robust, production-ready backend built with **Go** and the **Gin** web framework.  
It provides a scalable RESTful API with **PostgreSQL**, **GORM**, **Redis**, structured logging via **Zap**, real-time monitoring with **Prometheus & Grafana**, and full observability through the **ELK Stack (Elasticsearch, Logstash, Kibana)**.  
The project also includes **OTP-based authentication**, **Swagger documentation**, and a **configurable setup** via YAML/JSON config files.

---

## ✨ Features

- 🚀 **High-performance API** built with [Gin](https://github.com/gin-gonic/gin)
- 🗄️ **Database ORM** using [GORM](https://gorm.io) with **PostgreSQL**
- 🔐 **OTP Authentication** via phone number, backed by **Redis**
- 🧾 **Centralized logging** with [Zap](https://github.com/uber-go/zap)
- 📊 **Monitoring** with **Prometheus** & **Grafana**
- 📈 **Log pipeline** via **ELK Stack** + **Filebeat**
- ⚙️ **Configurable environment** using structured config file (YAML/JSON)
- 🧠 **API documentation** via **Swagger**
- 🧩 **Graceful shutdown**, **middleware layers**, and clean architecture principles

---

## 🧱 Tech Stack

| Layer | Technology |
|-------|-------------|
| **Language** | Go 1.22+ |
| **Framework** | [Gin](https://github.com/gin-gonic/gin) |
| **Database** | PostgreSQL |
| **ORM** | GORM |
| **Cache / OTP** | Redis |
| **Logger** | Zap (JSON structured logs) |
| **Log Pipeline** | Filebeat → Elasticsearch → Kibana |
| **Monitoring** | Prometheus + Grafana |
| **Docs** | Swagger |
| **Config** | YAML |

---

🧩 Logging & Observability
🔹 Zap Logger

Structured and high-performance JSON logging.

Different log levels (debug, info, warn, error, fatal).

Output streams to both stdout and a file.

🔹 ELK + Filebeat

Filebeat forwards logs to → Elasticsearch for indexing.

Kibana provides dashboards for search and visualization.

---

📊 Monitoring (Prometheus + Grafana)

Exposes metrics endpoint at /metrics

Collects:

Request latency and throughput

DB query performance

Goroutine count

Redis operations

Integrate with Grafana dashboards for visualization.

---
🧩 Middleware Highlights

Request logging (Zap)

Error recovery

Authentication (JWT or Session)

Request rate limiting (optional)

CORS & security headers
---
🪪 License

This project is released under the MIT License — free to use, modify, and distribute.
---
🙌 Acknowledgements

Gin Web Framework

GORM ORM

Zap Logger

Swaggo

Prometheus
 + Grafana

Elastic Stack (ELK)

Redis

## 📂 Project Structure


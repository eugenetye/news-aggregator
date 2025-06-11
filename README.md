# 📰 Personalized News Aggregator

A cloud-native, AI-powered microservice platform that scrapes Reddit and tech/sports news, summarizes content using GPT-4, stores results in Redis, and serves them via a public API and Alexa Flash Briefing.

## ✨ Features

- ✅ Reddit and RSS scraping (Go microservices)
- ✅ GPT-4 summarization via OpenAI API
- ✅ Daily scheduled aggregation (GKE CronJob)
- ✅ Redis-based caching with TTL support
- ✅ API Gateway using Go + Gin on Cloud Run
- ✅ Alexa Flash Briefing integration
- ✅ Deployed with Terraform and Kubernetes on GCP

## 🧱 Architecture

- **Go** for all backend services
- **Kubernetes** (GKE Autopilot) for CronJobs and Redis
- **Redis** as a shared data store (internal LoadBalancer)
- **Cloud Run** to serve API Gateway with autoscaling
- **Terraform** for infrastructure-as-code
- **GitHub Actions (optional)** for CI/CD
- **OpenAI GPT-4** for text summarization

## 🚀 Deployment

### Cloud Infrastructure

Provisioned via Terraform:

- GKE Autopilot Cluster
- Cloud Run Service
- VPC Connector
- Firewall Rules
- Redis Service

## 📡 API Endpoints

- `GET /feed/today`: Returns today's summarized news
- `GET /alexa/briefing`: Returns Alexa-compatible feed


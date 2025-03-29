# 🚀 Deployment Instructions

## 📦 Local Development

### Prerequisites
- Docker & Docker Compose installed

### Run Locally
```bash
docker-compose up --build
```
Visit the backend API at: [http://localhost:8080](http://localhost:8080)

---

## 🌐 Deployment to Render (Free Tier)

### Steps:
1. Push your project to GitHub.
2. Go to [Render.com](https://render.com) and sign in.
3. Create a **new Web Service**:
   - Select your repo
   - Use Go build command: `go build -o main .`
   - Use Start command: `./main`
   - Set environment variable: `DATABASE_URL` (you can use a hosted PostgreSQL like Supabase or Render PostgreSQL)
4. Add a Render PostgreSQL service and connect the URL to `DATABASE_URL`.

---

## 🌐 Deployment to Railway

1. Go to [Railway.app](https://railway.app)
2. Create a new project from your GitHub repo.
3. Add PostgreSQL plugin.
4. Add environment variable:
   - `DATABASE_URL` = Railway PostgreSQL connection string
5. Railway will auto-deploy your Go server.

Happy Hosting! 🌠

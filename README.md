# CCPP Full-Stack Application

A modern full-stack application built with Go + Gin + GORM backend and Vue.js + Tailwind CSS frontend, following clean architecture principles.

## 🏗️ Architecture

### Backend (Go)
- **Framework**: Gin (HTTP web framework)
- **ORM**: GORM (Object-Relational Mapping)
- **Database**: PostgreSQL
- **Architecture**: Clean Architecture with separation of concerns
  - Domain layer (models, repositories, services)
  - Infrastructure layer (handlers, database, external services)

### Frontend (Vue.js)
- **Framework**: Vue.js 3 with Composition API
- **Styling**: Tailwind CSS
- **Routing**: Vue Router
- **State Management**: Pinia
- **Build Tool**: Vite

## 📁 Project Structure

```
ccpp-site/
├── backend/                    # Go backend application
│   ├── cmd/                   # Application entry points
│   │   └── server/
│   │       └── main.go        # Main application file
│   ├── internal/              # Private application code
│   │   ├── domain/           # Domain layer (business logic)
│   │   │   ├── models/       # Domain models
│   │   │   ├── repositories/ # Repository interfaces
│   │   │   └── services/     # Business logic services
│   │   └── infrastructure/   # Infrastructure layer
│   │       ├── database/     # Database connection
│   │       ├── handlers/     # HTTP handlers
│   │       ├── repositories/ # Repository implementations
│   │       └── routes/       # Route definitions
│   ├── go.mod                # Go module file
│   └── env.example           # Environment variables example
├── frontend/                 # Vue.js frontend application
│   ├── src/
│   │   ├── components/      # Vue components
│   │   ├── views/          # Page components
│   │   ├── services/       # API services
│   │   ├── router/         # Vue Router configuration
│   │   ├── App.vue         # Root component
│   │   └── main.js         # Application entry point
│   ├── package.json        # Node.js dependencies
│   ├── vite.config.js      # Vite configuration
│   └── tailwind.config.js  # Tailwind CSS configuration
├── Makefile                # Build and development commands
└── README.md              # This file
```

## 🚀 Getting Started

### Prerequisites

- Go 1.21 or higher
- Node.js 18 or higher
- npm or yarn
- PostgreSQL 13 or higher

### Installation & Development

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd ccpp-site
   ```

2. **Install all dependencies**
   ```bash
   make install
   ```

3. **Set up PostgreSQL database**
   ```bash
   # Install PostgreSQL if not already installed
   # Create database: ccpp_db
   # Default credentials: user=postgres, password=password
   make db-setup
   ```

4. **Start development servers**
   ```bash
   make dev
   ```

   This will start:
   - Backend server at `http://localhost:8080`
   - Frontend server at `http://localhost:3000`

### Individual Commands

- **Install backend dependencies**: `make install-backend`
- **Install frontend dependencies**: `make install-frontend`
- **Start backend only**: `make dev-backend`
- **Start frontend only**: `make dev-frontend`
- **Build for production**: `make build`
- **Clean build artifacts**: `make clean`

## 🛠️ Available Commands

| Command | Description |
|---------|-------------|
| `make help` | Show available commands |
| `make install` | Install all dependencies |
| `make dev` | Start development servers |
| `make build` | Build for production |
| `make clean` | Clean build artifacts |
| `make test-backend` | Run backend tests |
| `make test-frontend` | Run frontend tests |

## 📡 API Endpoints

### Users
- `GET /api/v1/users` - Get all users
- `GET /api/v1/users/:id` - Get user by ID
- `POST /api/v1/users` - Create new user
- `PUT /api/v1/users/:id` - Update user
- `DELETE /api/v1/users/:id` - Delete user

### Health Check
- `GET /health` - API health status

## 🔧 Configuration

### Backend Configuration
Create a `.env` file in the `backend` directory:
```env
PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=ccpp_db
DB_SSLMODE=disable
GIN_MODE=debug
```

### Frontend Configuration
The frontend is configured to proxy API requests to the backend server running on port 8080.

## 🧪 Testing

- **Backend tests**: `make test-backend`
- **Frontend tests**: `make test-frontend`

## 🏗️ Build & Deployment

### Production Build
```bash
make build
```

This creates:
- Binary executable in `bin/server`
- Frontend build in `frontend/dist/`

### Running Production Build
```bash
# Start backend
./bin/server

# Serve frontend (use any static file server)
# Example with Python:
cd frontend/dist && python -m http.server 8000
```

## 🐳 Docker Deployment

### Using Docker Compose (Recommended)
```bash
# Start all services (PostgreSQL, backend, frontend)
docker-compose up -d

# View logs
docker-compose logs -f

# Stop all services
docker-compose down
```

### Manual Docker Setup
```bash
# Build and run PostgreSQL
docker run --name postgres-ccpp -e POSTGRES_DB=ccpp_db -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password -p 5432:5432 -d postgres:15-alpine

# Build and run backend
cd backend && docker build -t ccpp-backend .
docker run --name ccpp-backend --link postgres-ccpp -p 8080:8080 -e DB_HOST=postgres-ccpp ccpp-backend

# Build and run frontend
cd frontend && docker build -t ccpp-frontend .
docker run --name ccpp-frontend -p 3000:80 ccpp-frontend
```

## 🎨 Features

- ✅ User management (CRUD operations)
- ✅ RESTful API design
- ✅ Clean Service-Oriented Architecture
- ✅ Responsive UI with Tailwind CSS
- ✅ Modern Vue.js 3 with Composition API
- ✅ PostgreSQL database with GORM migrations
- ✅ CORS enabled for frontend-backend communication
- ✅ Environment-based configuration
- ✅ Development and production build scripts

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## 📝 License

This project is licensed under the MIT License.

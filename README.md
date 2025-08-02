# 🛒 Shopping Cart App (Fullstack)

A fullstack shopping cart application built with:

- **Backend**: Go (Gin), GORM, SQLite
- **Frontend**: ReactJS
- **Auth**: JWT
- **Database**: SQLite (lightweight file-based)
- **API Testing**: Postman Collection included

---

## 🚀 How to Run the Project

### 📂 Backend Setup

1. Install Go: https://go.dev/doc/install  
2. Install GCC for SQLite support:  
   ➤ [Install MinGW-w64 for Windows](https://www.mingw-w64.org/downloads/)
3. Navigate to backend folder:

```bash
cd shopping-cart-backend
go mod tidy
go run main.go
```

Your backend API will run at: [http://localhost:9090](http://localhost:9090)

---

### 💻 Frontend Setup

1. Install Node.js: https://nodejs.org/

2. Navigate to frontend folder:

```bash
cd shopping-cart-frontend
npm install
npm start
```

Frontend runs at: [http://localhost:3000](http://localhost:3000)

---

## 🧪 Postman Collection

Import `shopping-cart.postman_collection.json` (included in this repo) into Postman.

Test APIs like:

- `POST /users` → Create new user
- `POST /users/login` → Login and get JWT
- `GET /items` → List available items
- `POST /carts` → Add item to cart
- `POST /orders` → Place order

---

## 🔐 Login Example

**Username**: `adars`  
**Password**: `123456`  
→ Use the token from login response in `Authorization: Bearer <token>` header.

---

## 📁 Folder Structure

```
shopping-cart/
│
├── shopping-cart-backend/
│   ├── main.go
│   ├── controllers/
│   ├── database/
│   └── models/
│
├── shopping-cart-frontend/
│   ├── src/
│   └── public/
│
├── shopping-cart.postman_collection.json
└── README.md
```

---

## 📬 Submission

- ✅ Code hosted on GitHub
- ✅ README included
- ✅ Postman collection attached
- ✅ Works locally with SQLite

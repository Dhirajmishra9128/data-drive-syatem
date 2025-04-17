# data-drive-syatem
The project is build data drive system—similar to Google Drive—using Go with the Gin web framework.
# 📁 Data Drive System

A simple Google Drive–like system built with **Go (Gin)**, **MySQL**, and **HTML/CSS** that supports user authentication and CRUD operations for files and folders — including nested folders.

---

## 🚀 Features

- User Registration & Login
- Token-based Authorization
- CRUD Operations on Files & Folders
- Nested Folder Support (infinite depth)
- Clean HTML/CSS UI
- Postman support for API testing

---

## 🛠 Tech Stack

- [Go](https://golang.org/) (with Gin framework)
- [MySQL](https://www.mysql.com/)
- [GORM](https://gorm.io/) (ORM for Go)
- HTML/CSS (for simple UI)

---

## 📦 Folder Structure

```
data-drive-system/
├── main.go
├── config/          # DB config
├── models/          # User and File models
├── controllers/     # Auth & File handlers
├── middleware/      # Token validation
├── routes/          # All API routes
├── templates/       # HTML pages
├── static/          # CSS styling
└── README.md
```

---

## ✅ Setup Instructions

### 1. Install Dependencies
```bash
sudo apt update
sudo apt install golang mysql-server
```

### 2. Start MySQL and Create DB
```bash
sudo systemctl start mysql
sudo mysql
```
```sql
CREATE DATABASE datadrive;
```

### 3. Clone / Copy the Project
```bash
git clone https://github.com/Dhirajmishra9128/data-drive-syatem.git  
cd data-drive-system
go mod init data-drive-system
go get ./...
```

### 4. Run the App
```bash
go run main.go
```

Open browser: [http://localhost:8080](http://localhost:8080)

---

## 🔐 Token-based Authorization
- After login, you get a token (user ID).
- Use this token in `Authorization` header for protected APIs.

---

## 📮 API Endpoints

| Method | Endpoint            | Description              |
|--------|---------------------|--------------------------|
| POST   | `/register`         | Create new user          |
| POST   | `/login`            | Login and get token      |
| GET    | `/dashboard`        | List user files/folders  |
| POST   | `/drive/create`     | Create file/folder       |
| POST   | `/drive/update/:id` | Rename file/folder       |
| POST   | `/drive/delete/:id` | Delete file/folder       |

---

## 📬 Testing with Postman

**Header:**
```
Authorization: 1
```
Replace `1` with your user ID (token returned after login).

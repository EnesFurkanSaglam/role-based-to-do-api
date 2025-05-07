# 🗂️ Role-Based TO-DO API (Golang)

## 🚀 Başlatmak için

🔗 **Canlı API Adresi:** [https://role-based-to-do-api-production.up.railway.app](https://role-based-to-do-api-production.up.railway.app)

Bu proje, role-based kullanıcı yönetimi ile çalışan bir yapılacaklar (TO-DO) listesinin **backend API** tarafını içermektedir.  
Proje, Golang diliyle yazılmıştır ve `MVC + Clean Architecture` desenine uygun geliştirilmiştir.  

## 📌 Özellikler

- 🔐 **JWT tabanlı kullanıcı doğrulama**
- 👥 **Role-Based Yetkilendirme** (admin / basic)
- ✅ TO-DO list & adım (step) yönetimi
- 🧱 **Soft Delete** (veri gerçekten silinmez)
- ⏱️ `UpdatedAt` güncelleme mantığı
- 📊 Step tamamlandıkça **% tamamlanma oranı**
- 📂 Katmanlı yapı: `controller`, `service`, `repository`, `model`, `middleware`, `util`

## 👤 Kullanıcı Bilgileri

| Username | Password    | Role   |
|----------|-------------|--------|
| `enes`   | `1234`      | basic  |
| `admin`  | `adminpass` | admin  |

## 🚀 Başlatmak için

### 1. Çevre Değişkenleri `.env`
```
JWT_SECRET_KEY=senin_jwt_keyin
```

### 2. Terminal:
```bash
go run cmd/main.go
```

## 🧪 API Endpointleri

### 🔑 /login
```http
POST /login
{
  "username": "enes",
  "password": "1234"
}
```
🟢 Yanıt: `{ "token": "..." }`

### 📁 /lists (Listeler)
| Endpoint             | Açıklama                          |
|----------------------|-----------------------------------|
| `POST /lists`        | Yeni liste oluşturur              |
| `GET /lists`         | Liste(leri) getirir               |
| `POST /lists/delete` | Listeyi soft siler                |
| `POST /lists/update` | Liste adını günceller             |

**Örnek Liste Ekleme:**
```json
{
  "name": "Final Proje Görevleri"
}
```

### 📄 /steps (Adımlar)
| Endpoint                | Açıklama                             |
|-------------------------|--------------------------------------|
| `POST /steps`           | Adım (görev) ekler                   |
| `GET /steps?list_id=1`  | Belirli listenin adımlarını listeler |
| `POST /steps/complete`  | Adımı tamamlandı olarak işaretler    |
| `POST /steps/delete`    | Adımı soft siler                     |
| `POST /steps/update`    | Adım içeriğini günceller             |

**Örnek Step Ekleme:**
```json
{
  "list_id": 1,
  "content": "Yazılım PDF raporunu hazırla"
}
```

## 🛡️ Role Bazlı Erişim Mantığı

- `basic` kullanıcı: yalnızca **kendi listelerini ve adımlarını** görebilir
- `admin` kullanıcı: **tüm kullanıcıların** verilerine erişebilir
- Tüm işlemler `JWT` ile yetkilendirilir, `Authorization: Bearer <token>` header'ı gereklidir

## 📁 Proje Yapısı

```
role-based-to-do-api/
├── cmd/
│   └── main.go
├── internal/
│   ├── controller/
│   ├── service/
│   ├── repository/
│   ├── middleware/
│   ├── model/
│   └── util/
├── .env
└── go.mod
```




# 🗂️ Role-Based TO-DO API (Golang)

## 🚀 Getting Started

🔗 **Live API URL:** [https://role-based-to-do-api-production.up.railway.app](https://role-based-to-do-api-production.up.railway.app)

This project contains the **backend API** of a TO-DO list that works with role-based user management.  
It is written in Golang and developed following the `MVC + Clean Architecture` pattern.

## 📌 Features

- 🔐 **JWT-based user authentication**
- 👥 **Role-Based Authorization** (admin / basic)
- ✅ TO-DO list & step management
- 🧱 **Soft Delete** (data is not physically deleted)
- ⏱️ `UpdatedAt` update logic
- 📊 **% completion rate** based on steps completed
- 📂 Layered structure: `controller`, `service`, `repository`, `model`, `middleware`, `util`

## 👤 User Credentials

| Username | Password    | Role   |
|----------|-------------|--------|
| `enes`   | `1234`      | basic  |
| `admin`  | `adminpass` | admin  |

## 🚀 Getting Started

### 1. Environment Variables `.env`
```
JWT_SECRET_KEY=your_jwt_key
```

### 2. Terminal:
```bash
go run cmd/main.go
```

## 🧪 API Endpoints

### 🔑 /login
```http
POST /login
{
  "username": "enes",
  "password": "1234"
}
```
🟢 Response: `{ "token": "..." }`

### 📁 /lists

| Endpoint             | Description                       |
|----------------------|-----------------------------------|
| `POST /lists`        | Creates a new list                |
| `GET /lists`         | Retrieves list(s)                 |
| `POST /lists/delete` | Soft-deletes a list               |
| `POST /lists/update` | Updates a list name               |

**Example Add List:**
```json
{
  "name": "Final Project Tasks"
}
```

### 📄 /steps

| Endpoint                | Description                            |
|-------------------------|----------------------------------------|
| `POST /steps`           | Adds a step (task)                     |
| `GET /steps?list_id=1`  | Lists steps of a specific list         |
| `POST /steps/complete`  | Marks a step as completed              |
| `POST /steps/delete`    | Soft-deletes a step                    |
| `POST /steps/update`    | Updates step content                   |

**Example Add Step:**
```json
{
  "list_id": 1,
  "content": "Prepare software PDF report"
}
```

## 🛡️ Role-Based Access Logic

- `basic` user: can only see **their own lists and steps**
- `admin` user: can access **all users' data**
- All actions require `JWT` authorization with `Authorization: Bearer <token>` header

## 📁 Project Structure

```
role-based-to-do-api/
├── cmd/
│   └── main.go
├── internal/
│   ├── controller/
│   ├── service/
│   ├── repository/
│   ├── middleware/
│   ├── model/
│   └── util/
├── .env
└── go.mod
```
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
|   
├── internal/
│   ├── controller/
│   ├── service/
│   ├── repository/
│   ├── middleware/
│   ├── model/
│   └── util/
├── .env
└── go.mod
└── main.go
```

---

## 🔐 /login

### Genel Bilgiler

| Özellik          | Değer              |
|------------------|--------------------|
| Yöntem (Method)  | `POST`             |
| URL              | `/login`           |
| Kimlik Doğrulama | Gerekmez           |
| İçerik Tipi      | `application/json` |

### İstek
```json
{
  "username": "enes",
  "password": "1234"
}
```

### Yanıtlar

**200 OK**
```json
{
  "token": "JWT_Token_Bilgisi"
}
```

**401 Unauthorized**
```json
{
  "error": "invalid credentials"
}
```

### Notlar

- Başarılı giriş sonrası dönen token, tüm korumalı isteklerde `Authorization: Bearer <token>` olarak kullanılmalıdır.

---

## 📁 /lists

### 🔸 GET /lists

| Özellik          | Değer              |
|------------------|--------------------|
| Yöntem           | `GET`              |
| Kimlik Doğrulama | Gerekli (`JWT`)    |
| Açıklama         | Kullanıcının (veya admin ise tüm kullanıcıların) listelerini getirir |

**Yanıt:**
```json
[
  {
    "id": 1,
    "name": "Final Projesi",
    "owner": "enes",
    "is_deleted": false,
    "updated_at": "2024-05-07T15:00:00Z"
  }
]
```

---

### 🔸 POST /lists

| Özellik          | Değer              |
|------------------|--------------------|
| Yöntem           | `POST`             |
| Kimlik Doğrulama | Gerekli (`JWT`)    |
| Açıklama         | Yeni bir liste oluşturur |

**İstek:**
```json
{
  "name": "Yeni Liste Adı"
}
```

**Yanıt:**
```json
{
  "id": 2,
  "name": "Yeni Liste Adı",
  "owner": "enes"
}
```

**400 Bad Request:**
```json
{
  "error": "list name required"
}
```

---

### 🔸 POST /lists/delete

| Özellik          | Değer              |
|------------------|--------------------|
| Yöntem           | `POST`             |
| Kimlik Doğrulama | Gerekli (`JWT`)    |
| Açıklama         | Listeyi soft delete yapar |

**İstek:**
```json
{
  "list_id": 1
}
```

**Yanıt:**
```json
{
  "message": "List deleted"
}
```

**404 Not Found:**
```json
{
  "error": "list not found"
}
```

---

### 🔸 POST /lists/update

| Özellik          | Değer              |
|------------------|--------------------|
| Yöntem           | `POST`             |
| Kimlik Doğrulama | Gerekli (`JWT`)    |
| Açıklama         | Listenin adını günceller |

**İstek:**
```json
{
  "list_id": 1,
  "new_name": "Güncellenmiş Liste Adı"
}
```

**Yanıt:**
```json
{
  "message": "List updated"
}
```

---

## 📄 /steps

### 🔸 GET /steps?list_id=X

| Özellik          | Değer              |
|------------------|--------------------|
| Yöntem           | `GET`              |
| Kimlik Doğrulama | Gerekli (`JWT`)    |
| Açıklama         | Belirli bir listeye ait adımları getirir |

**Yanıt:**
```json
[
  {
    "id": 1,
    "list_id": 1,
    "content": "Rapor hazırla",
    "is_completed": false
  }
]
```

---

### 🔸 POST /steps

| Özellik          | Değer              |
|------------------|--------------------|
| Yöntem           | `POST`             |
| Kimlik Doğrulama | Gerekli (`JWT`)    |
| Açıklama         | Yeni bir adım (step) ekler |

**İstek:**
```json
{
  "list_id": 1,
  "content": "Sunumu yap"
}
```

**Yanıt:**
```json
{
  "id": 3,
  "list_id": 1,
  "content": "Sunumu yap",
  "is_completed": false
}
```

---

### 🔸 POST /steps/complete

| Özellik          | Değer              |
|------------------|--------------------|
| Yöntem           | `POST`             |
| Kimlik Doğrulama | Gerekli (`JWT`)    |
| Açıklama         | Adımı tamamlandı olarak işaretler |

**İstek:**
```json
{
  "step_id": 3
}
```

**Yanıt:**
```json
{
  "message": "Step completed"
}
```

---

### 🔸 POST /steps/delete

| Özellik          | Değer              |
|------------------|--------------------|
| Yöntem           | `POST`             |
| Kimlik Doğrulama | Gerekli (`JWT`)    |
| Açıklama         | Adımı soft delete yapar |

**İstek:**
```json
{
  "step_id": 2
}
```

**Yanıt:**
```json
{
  "message": "Step deleted"
}
```

---

### 🔸 POST /steps/update

| Özellik          | Değer              |
|------------------|--------------------|
| Yöntem           | `POST`             |
| Kimlik Doğrulama | Gerekli (`JWT`)    |
| Açıklama         | Adım içeriğini günceller |

**İstek:**
```json
{
  "step_id": 1,
  "new_content": "Yeni içerik"
}
```

**Yanıt:**
```json
{
  "message": "Step updated"
}
```


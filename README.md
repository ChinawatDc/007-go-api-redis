# 007-go-api-redis

ใช้ Redis สำหรับ Caching และ Session Management ด้วย Go

## 📦 ใช้แพ็กเกจ

- Gin สำหรับ HTTP routing
- Redis (go-redis v8) สำหรับ caching
- godotenv สำหรับโหลด config จาก `.env`

## 📂 โครงสร้างไฟล์โปรเจกต์

```bash
007-go-api-redis/
├── config/
├── controllers/
│   └── cache.go
├── routes/
│   └── router.go
├── .env
├── go.mod
├── go.sum
├── main.go
├── README.md
```

## 🔧 การติดตั้ง Redis Client
ติดตั้งไลบรารี Redis ด้วยคำสั่ง:
```bash
go get github.com/go-redis/redis/v8
```

## ⚙️ วิธีโหลด config Redis จากไฟล์ .env

สร้างไฟล์ .env ในโฟลเดอร์โปรเจ็ค และเพิ่มค่า:

```bash
REDIS_ADDR=localhost:6379
REDIS_PASSWORD=
```

## 🚀 วิธีใช้งาน

```bash
go run main.go
```

## 📂 โฟลเดอร์

- main.go – Entry point
- config/ – สำหรับ Redis config (อนาคต)
- controllers/ – สำหรับแยก handler
- routes/ – สำหรับแยก routing

📧 ติดต่อ
ถ้ามีคำถามหรือข้อเสนอแนะ ติดต่อได้เลยครับ!
---

📄 License  
MIT License – ใช้งานฟรีเพื่อการศึกษาและพัฒนาต่อได้เต็มที่

🙋‍♂️ ผู้พัฒนา  
พัฒนาโดยคุณ [ปลื้ม ชินวัตร]  
📫 ติดต่อ: chinawat.dc@gmail.com

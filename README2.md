# API Wilayah Kecamatan Klaten

Tiga API sederhana untuk menyimpan data wilayah ke DynamoDB menggunakan:

- Python (Flask)
- Node.js (Express)
- Golang

## Struktur

```
python-api/
nodejs-api/
golang-api/
```

## Menjalankan API

### Python
```bash
cd python-api
pip install -r requirements.txt
python3 app.py
```

### Node.js
```bash
cd nodejs-api
npm install
node app.js
```

### Golang
```bash
cd golang-api
go run main.go
```

### Test
Gunakan Postman:
- Method: `POST`
- URL: `http://<ec2-ip>:3000/kecamatan` (Python), `3001` (Node), `3002` (Golang)
- Body (JSON):
```json
{
  "id": "01",
  "nama": "Wedi"
}
```

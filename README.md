# klaten-api
**satu per satu** secara **sangat jelas** dan **berurutan**.

---

## 🌐 Bagian 1: Infrastruktur Dasar (VPC dan Jaringannya)

### 1. **Buat VPC**
- Masuk ke AWS Console → cari dan buka **VPC**.
- Klik **Your VPCs** → **Create VPC**
- Pilih **VPC only**.
- Masukkan nama: `KlatenVPC`
- IPv4 CIDR block: `10.0.0.0/16`
- Klik **Create VPC**

---

### 2. **Buat Subnet**
- Klik **Subnets** → **Create subnet**
- VPC: Pilih `KlatenVPC`
- Subnet name: `PublicSubnet1`
- Availability Zone: Misal `us-east-1a`
- IPv4 CIDR block: `10.0.1.0/24`
- Klik **Create subnet**

---

### 3. **Buat Internet Gateway**
- Klik **Internet Gateways** → **Create internet gateway**
- Name: `KlatenIGW`
- Klik **Create internet gateway**
- Setelah selesai, klik **Attach to VPC** → pilih `KlatenVPC`

---

### 4. **Buat Route Table**
- Klik **Route Tables** → **Create route table**
- Name: `PublicRouteTable`
- VPC: `KlatenVPC`
- Klik **Create**
- Klik tab **Routes** → **Edit routes**
  - Tambahkan route:
    - Destination: `0.0.0.0/0`
    - Target: Internet Gateway (`KlatenIGW`)
- Klik tab **Subnet Associations** → **Edit subnet associations**
  - Centang `PublicSubnet1`
  - Klik **Save associations**

---

### 5. **Buat Security Group**
- Masuk ke **Security Groups** → klik **Create Security Group**
- Name: `KlatenSG`
- VPC: `KlatenVPC`
- Inbound rules:
  - SSH (port 22) – Source: your IP (atau 0.0.0.0/0 untuk tes)
  - HTTP (port 80)
  - Custom TCP (port 3000 untuk Python, 3001 untuk Node.js, 3002 untuk Go)
- Klik **Create security group**

---

## 🗃️ Bagian 2: Database DynamoDB

### 6. **Buat DynamoDB Table**
- Masuk ke AWS Console → cari dan buka **DynamoDB**
- Klik **Create table**
- Table name: `KecamatanKlaten`
- Partition key: `id` (type: String)
- Klik **Create Table**

---

## 📁 Bagian 3: Siapkan API Folder (Lokal)

### 7. **Struktur Folder Lokal**
```
klaten-api/
├── api-python/
│   ├── app.py
│   └── requirements.txt
├── api-nodejs/
│   ├── app.js
│   └── package.json
└── api-golang/
    └── main.go
```

---

### 8. **Contoh Isi File**
#### ✅ Python (`api-python/app.py`)
```python
from flask import Flask, request, jsonify
import boto3

app = Flask(__name__)
dynamodb = boto3.resource('dynamodb', region_name='us-east-1')
table = dynamodb.Table('KecamatanKlaten')

@app.route('/kecamatan', methods=['POST'])
def add_data():
    data = request.json
    table.put_item(Item=data)
    return jsonify({"message": "Data berhasil ditambahkan"})

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=3000)
```

#### ✅ Node.js (`api-nodejs/app.js`)
```js
const express = require('express');
const AWS = require('aws-sdk');
const app = express();
app.use(express.json());

AWS.config.update({ region: 'us-east-1' });
const docClient = new AWS.DynamoDB.DocumentClient();
const tableName = 'KecamatanKlaten';

app.post('/kecamatan', (req, res) => {
    const data = req.body;
    const params = {
        TableName: tableName,
        Item: data
    };
    docClient.put(params, (err) => {
        if (err) res.status(500).send(err);
        else res.send({ message: "Data berhasil ditambahkan" });
    });
});

app.listen(3001, () => console.log('Node.js API running on port 3001'));
```

#### ✅ Golang (`api-golang/main.go`)
```go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Kecamatan struct {
	ID   string `json:"id"`
	Nama string `json:"nama"`
}

func main() {
	sess := session.Must(session.NewSession(&aws.Config{Region: aws.String("us-east-1")}))
	svc := dynamodb.New(sess)

	http.HandleFunc("/kecamatan", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			var data Kecamatan
			json.NewDecoder(r.Body).Decode(&data)

			item := map[string]*dynamodb.AttributeValue{
				"id":   {S: aws.String(data.ID)},
				"nama": {S: aws.String(data.Nama)},
			}

			_, err := svc.PutItem(&dynamodb.PutItemInput{
				TableName: aws.String("KecamatanKlaten"),
				Item:      item,
			})

			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			json.NewEncoder(w).Encode(map[string]string{"message": "Data berhasil ditambahkan"})
		}
	})

	http.ListenAndServe(":3002", nil)
}
```

---

## 🧬 Bagian 4: Upload ke GitHub

### 9. **Upload Folder**
- Buka terminal lokal kamu:
```bash
git init
git remote add origin https://github.com/Dcyii/klaten-api.git
git add .
git commit -m "Initial commit"
git push -u origin master
```

---

## ☁️ Bagian 5: EC2 + Jalankan API

### 10. **Buat EC2 Instance**
- Buka AWS Console → EC2 → Launch Instance
- AMI: Amazon Linux 2023
- Instance type: t2.micro
- Network: Pilih `KlatenVPC`, subnet: `PublicSubnet1`
- Security Group: Pilih `KlatenSG`
- Klik **Launch**

---

### 11. **Clone Repo dan Jalankan API**
- Login SSH ke EC2
- Install Git & tools:
```bash
sudo dnf update -y
sudo dnf install -y git python3-pip nodejs golang
```

- Clone project:
```bash
git clone https://github.com/Dcyii/klaten-api.git
cd klaten-api
```

- Jalankan satu-satu:
```bash
# Python
cd api-python
pip install -r requirements.txt
python3 app.py &

# Node.js
cd ../api-nodejs
npm install
node app.js &

# Golang
cd ../api-golang
go run main.go &
```

---

## 📬 Bagian 6: Testing API

### 12. **Gunakan Postman**
- Kirim `POST` ke:
  - http://<EC2_IP>:3000/kecamatan (Python)
  - http://<EC2_IP>:3001/kecamatan (Node.js)
  - http://<EC2_IP>:3002/kecamatan (Golang)

- Format body (JSON):
```json
{
  "id": "001",
  "nama": "Kecamatan Wedi"
}
```

---

## 🔁 Bagian 7: Auto Scaling

### 13. **Buat Launch Template & Auto Scaling Group**
- Masuk ke EC2 → Launch Templates → Create
- Pilih konfigurasi yang sama seperti instance yang tadi

- Masuk ke Auto Scaling Group → Create Auto Scaling Group
  - Pilih Launch Template
  - Pilih VPC/Subnet
  - Attach Load Balancer (jika sudah dibuat)

---

const express = require("express");
const AWS = require("aws-sdk");

const app = express();
app.use(express.json());

AWS.config.update({ region: "us-east-1" });
const dynamodb = new AWS.DynamoDB.DocumentClient();
const tableName = "KecamatanKlaten";

app.post("/kecamatan", (req, res) => {
  const item = req.body;

  const params = {
    TableName: tableName,
    Item: item
  };

  dynamodb.put(params, (err, data) => {
    if (err) {
      res.status(500).json({ error: "Gagal menambahkan data" });
    } else {
      res.json({ message: "Data berhasil ditambahkan" });
    }
  });
});

app.listen(3001, () => {
  console.log("Server berjalan di port 3001");
});

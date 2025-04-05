package main

import (
    "encoding/json"
    "fmt"
    "log"
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
    sess := session.Must(session.NewSession(&aws.Config{
        Region: aws.String("us-east-1"),
    }))
    svc := dynamodb.New(sess)

    http.HandleFunc("/kecamatan", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != "POST" {
            http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
            return
        }

        var kecamatan Kecamatan
        json.NewDecoder(r.Body).Decode(&kecamatan)

        input := &dynamodb.PutItemInput{
            TableName: aws.String("KecamatanKlaten"),
            Item: map[string]*dynamodb.AttributeValue{
                "id":   {S: aws.String(kecamatan.ID)},
                "nama": {S: aws.String(kecamatan.Nama)},
            },
        }

        _, err := svc.PutItem(input)
        if err != nil {
            http.Error(w, "Gagal menyimpan data", http.StatusInternalServerError)
            return
        }

        w.Write([]byte(`{"message":"Data berhasil ditambahkan"}`))
    })

    fmt.Println("Server berjalan di port 3002")
    log.Fatal(http.ListenAndServe(":3002", nil))
}

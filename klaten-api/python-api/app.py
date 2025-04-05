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

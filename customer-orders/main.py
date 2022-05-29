from flask import Flask, jsonify, current_app
from waitress import serve
from time import sleep
from random import uniform


app = Flask(__name__)

@app.route("/customers/<customer_id>/orders")
def customerOrders():
    sleep(uniform(0.15, 0.25))
    return jsonify([
        {"WashMaster 3000": "1"},
        {"eyePhone 4FS": "1"}
    ])

@app.route("/health")
def healthCheck():
    return "ok"


if __name__ == '__main__':
    serve(app, host='0.0.0.0', port=8080)

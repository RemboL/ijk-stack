from flask import Flask, jsonify, current_app
from waitress import serve
from time import sleep
from random import uniform


app = Flask(__name__)

@app.route("/verifyProductsAvailability", methods = ['POST'])
def verifyProductsAvailability():
    sleep(uniform(0.15, 0.25))
    return jsonify({})

@app.route("/health")
def healthCheck():
    return "ok"


if __name__ == '__main__':
    serve(app, host='0.0.0.0', port=8080)

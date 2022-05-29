from flask import Flask, jsonify, current_app
from waitress import serve
from time import sleep
from random import uniform


app = Flask(__name__)

@app.route("/special-offers")
def specialOffers():
    sleep(uniform(0.05, 0.15))
    return jsonify({"eyePhone 4FS": "-2000"})

@app.route("/health")
def healthCheck():
    return "ok"


if __name__ == '__main__':
    serve(app, host='0.0.0.0', port=8080)

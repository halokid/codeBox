from flask import Flask, jsonify
from flask_cors import cross_origin

app = Flask(__name__)


@app.route('/pizza')
@cross_origin()
def root():
  t = [
    {"id": 1, "name": "Cheese pizza", "description": "very cheesy"},
    {"id": 2, "name": "Al Tono pizza", "description": "lots of tuna"}
  ]

  # t = {
  #   'a': 1,
  #   'b': 2,
  #   'c': [3, 4, 5]
  # }
  return jsonify(t)


if __name__ == '__main__':
  app.debug = True
  app.run()

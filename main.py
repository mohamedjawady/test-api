from flask import Flask, jsonify, request, abort
from flask_jwt_extended import JWTManager, create_access_token, jwt_required
from flask import make_response

app = Flask(__name__)
app.config['SECRET_KEY'] = 'supersecretkey'
app.config['JWT_SECRET_KEY'] = 'secret'
jwt = JWTManager(app)

# Dummy data
cities = [
    {"name": "Tunis", "country": "Tunisia", "population": 1200000, "latitude": 36.8065, "longitude": 10.1815, "language": "Arabic"},
    {"name": "New York", "country": "USA", "population": 8500000, "latitude": 40.7128, "longitude": -74.0060, "language": "English"},
    {"name": "Paris", "country": "France", "population": 2100000, "latitude": 48.8566, "longitude": 2.3522, "language": "French"},
    {"name": "Tokyo", "country": "Japan", "population": 13929286, "latitude": 35.6895, "longitude": 139.6917, "language": "Japanese"},
    {"name": "London", "country": "UK", "population": 8982000, "latitude": 51.5074, "longitude": -0.1278, "language": "English"},
    # Add other cities as needed
]

# Routes
@app.route('/login', methods=['POST'])
def login():
    req = request.get_json()
    username = req.get('username')
    password = req.get('password')
    
    # Validate credentials
    if username != "admin" or password != "password":
        abort(403, 'Invalid credentials')
    
    # Create JWT token
    token = create_access_token(identity={'username': username})
    
    return jsonify({'token': token})

@app.route('/api/cities', methods=['GET'])
@jwt_required()
def get_cities():
    city_name = request.args.get('city')
    
    if city_name:
        city = next((city for city in cities if city['name'] == city_name), None)
        if city:
            return jsonify(city)
        else:
            abort(404, 'City not found')
    
    return jsonify(cities)

if __name__ == '__main__':
    app.run(debug=True, port=8000)


# pip install Flask Flask-JWT-Extended

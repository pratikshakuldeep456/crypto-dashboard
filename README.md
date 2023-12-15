# crypto-dashboard

# Clone the repository
git clone https://github.com/pratiksha456/crypto-dashboard.git
cd crypto-dashboard

# Set up environment variables for MongoDB (replace with your actual values)
# For MongoDB Atlas
export MONGO_DB_URI="your-mongodb-atlas-connection-string"

# OR for local MongoDB
export MONGO_DB_URI="mongodb://localhost:27017"

# Install Go dependencies
go get -u github.com/gofiber/fiber/v2
go get -u go.mongodb.org/mongo-driver/mongo
go get -u go.mongodb.org/mongo-driver/mongo/options

# Run the Go application
go run main.go

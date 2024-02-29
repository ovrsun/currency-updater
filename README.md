# Currency Exchange Rate Checking Service

This service provides functionality to update currency exchange rates, retrieve exchange rates by identifier, and get the latest exchange rate for a specified currency.

## Key Features

1. **Update Exchange Rate**

   Updates the exchange rate for the specified currency. Upon successful update, returns the update identifier.

   **Example Request:**
   ```
   POST /updates
   {
       "code": "EUR/MXN"
   }
   ```
   **Example Response:**
   ```
   {
       "id": 1
   }
   ```

2. **Get Exchange Rate by Identifier**

   Retrieves the exchange rate for the specified update identifier.

   **Example Request:**
   ```
   GET /updates/1
   ```
   **Example Response:**
   ```
   {
       "id": 1,
       "code": "EUR/MXN",
       "updated": "2024-02-28T20:07:27.874786Z",
       "rate": 18.5126
   }
   ```

3. **Get Latest Exchange Rate**

   Retrieves the latest exchange rate for the specified currency.

   **Example Request:**
   ```
   GET /updates/?code=EUR/RUB
   ```
   **Example Response:**
   ```
   {
       "id": 38,
       "code": "EUR/RUB",
       "updated": "2024-02-29T10:00:54.661401Z",
       "rate": 99.3353
   }
   ```

## Running the Service

To run the service, follow these steps:

1. Clone the repository to your local machine.
2. Create a `.env` (mv .env.example .env) file and specify your API key for the external service exchangerate-api.com.
3. Start Docker Compose using the following command:
   ```
   docker-compose up
   ```

## Technologies Used

- Programming Language: Go
- Database: PostgreSQL
- External Data Source: exchangerate-api.com
- Docker Compose for containerization of the application

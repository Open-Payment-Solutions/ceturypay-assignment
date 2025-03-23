# Concurrent Money Transfer System

This is a simple concurrent money transfer system built in Go that allows users to transfer money between accounts atomically and safely.

## Features

- Transfer money between accounts
- Atomic updates to prevent race conditions
- Prevent overdrafts (users cannot send more money than they have)
- RESTful HTTP API for account management and transfers
- Concurrent processing of transfers

## Installation

1. Clone the repository:

```bash
git clone https://github.com/yourusername/centurypay.git
cd centurypay
go mod download
```

## Running the Application

Start the server:

```bash
go run cmd/server/main.go
```

The server will start on port 8000 with three pre-configured accounts:
- Mark with $100 USD
- Jane with $50 USD
- Adam with $0 USD

## API Endpoints

### Get All Accounts

```
GET /accounts
```

Returns a list of all accounts in the system.

### Get Account by ID

```
GET /accounts/:id
```

Returns details for a specific account.

### Create a Transfer

```
POST /transfers
```

Request body:
```json
{
  "fromAccountId": "acc-123...",
  "toAccountId": "acc-456...",
  "amount": 25.00,
  "currency": "USD"
}
```

This will create and process a transfer between the accounts.

For simplicity to test using I've added this line in `internal/services/accounts_service.go`:
```go 
account.ID = fmt.Sprintf("%d", len(s.accounts)+1) // uncomment this for api request testing 
```
I've made it to not mess with 32 chars long hash - during testing via api calls.

Feel free to comment it out if You want ID-s to be 32 chars long hashes.

### Get Transaction Status

```
GET /transactions/:id
```

Returns the status and details of a specific transaction.

## Architecture

The system consists of the following components:

1. **Account Service**: Manages accounts and their balances
2. **Transaction Service**: Manages money transfers between accounts
3. **Transaction Processor**: Processes transactions asynchronously (assumed to work with external payment gateway)
4. **API Layer**: HTTP endpoints for interacting with the system

### Concurrency Strategy

The system uses mutex locks to ensure thread safety:

1. **Read-Write Locks**: Used for account and transaction access to allow concurrent reads but exclusive writes
2. **Two-Phase Process**: Transfers happen in multiple phases (Created → Pending → Confirmed → Completed) with atomic updates at each stage
3. **Account Balance Locking**: When a transfer is initiated, the amount is immediately deducted from the sender's account and held until the transfer completes or fails

## Error Handling

The system handles several error conditions:
- Insufficient funds
- Transfers to the same account
- Invalid account IDs
- Expired transactions
- Different currencies

## Testing

Run the tests:

```bash
go test ./...
```
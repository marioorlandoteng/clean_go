# Golang Clean Architecture

## Requirements
Money Transfer APIs using Golang programming consist of 3 endpoints, namely:
1. Account Validation URL endpoint that used to validate account number and name of the bank account owner. To validate the account, use mock endpoint using https://mockapi.io/ or other preferred service that will act as a bank.
2. Transfer/Disbursement URL endpoint that used to transfer money to the destination account. To transfer the money, use mock endpoint similar to point 1 that acts as a bank.
3. Transfer/Disbursement Callback URL endpoint that used to receive callback status of transfer from the bank.

## Project Structure
This project use clean architecture from https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html.
The implementation is based on https://github.com/bxcodec/go-clean-arch.git

It has four layers
1. Model Layer: store model that being used as data transfer between layer, e.g. `moneytransfer/domain`
2. Repository Layer: connect with database or mockapi to get data, e.g. `moneytransfer/internal/repository`. Especially in mockapi the database acts as proxy and has concurrency logic.
3. Usecase Layer: hold business logic like status change, e.g. `moneytransfer/service`
4. Delivery Layer: hold delivery logic like rest, e.g. `moneytransfer/internal/handler`

![golang clean architecture](https://github.com/bxcodec/go-clean-arch/raw/master/clean-arch.png)

## How to Run
* Install postgres and run `moneytransfer.sql` in sql folder
* Install go and its dependencies, then run `go run main.go`

## Resources
* MockAPI URL to get users https://663eb5dbe3a7c3218a4b345f.mockapi.io/api/v1/users
* MockAPI URL to get bank transfer https://663eb5dbe3a7c3218a4b345f.mockapi.io/api/v1/banks

## cURL Example
1. Validate bank account and name
    Format
    ```
    curl --header "Content-Type: application/json" \
      --request GET \
      http://localhost:9090/api/v1/bank/account/[account_number]/[account_name]
    ```
    Example
    ```
    curl --header "Content-Type: application/json" \
      --request GET \
      http://localhost:9090/api/v1/bank/account/1/Jesus%20Langworth
    ```
    
2. Create transaction
    Example
    ```
    curl --header "Content-Type: application/json" \
      --request POST \
      --data '{"fromAccountNo":1,"toAccountNo":2,"amount":100}' \
      http://localhost:9090/api/v1/bank/transaction
    ```
    > Note: This endpoint will create ref_id that will be used in disbursement process
    
3. Disburse transaction
    Format
    ```
    curl --header "Content-Type: application/json" \
      --request PUT \
      http://localhost:9090/api/v1/bank/transaction/[ref_id]
    ```
    Example
    ```
    curl --header "Content-Type: application/json" \
      --request PUT \
      http://localhost:9090/api/v1/bank/transaction/0ab4d42e-f4f9-4b9f-afd6-43e0418b49b1
    ```

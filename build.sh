#!/bin/bash

docker build products -t products:0.0.1
docker build special-offers -t special-offers:0.0.1
docker build carts -t carts:0.0.1
docker build customer-consents -t customer-consents:0.0.1
docker build customer-orders -t customer-orders:0.0.1
docker build customers -t customers:0.0.1
docker build spedition -t spedition:0.0.1
docker build warehouse -t warehouse:0.0.1
docker build delivery -t delivery:0.0.1
docker build gift-cards -t gift-cards:0.0.1
docker build debtors -t debtors:0.0.1
docker build credit-cards -t credit-cards:0.0.1
docker build installments -t installments:0.0.1
docker build payments -t payments:0.0.1
docker build commerce -t commerce:0.0.4


#!/bin/bash


url=`minikube service commerce --url`

curl -v -X POST "$url/finalizePurchase" --header "content-type: application/json" --header "accept: application/json" --data '{"cartId":13,"customerId":41,"paymentInfo":{"creditCardId":1001,"paymentType":"TWELVE_INSTALLMENTS","giftCard":"ABC123"}}'

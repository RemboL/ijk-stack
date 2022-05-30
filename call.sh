#!/bin/bash

set -euCo pipefail

count=${1:-10}
url=`minikube service commerce --url`

for i in $(seq 1 $count)
do
	echo "calling #$i..."
	start=$(date +%s.%3N)
	curl -s -X POST "$url/finalizePurchase" --header "content-type: application/json" --header "accept: application/json" --data '{"cartId":13,"customerId":41,"paymentInfo":{"creditCardId":1001,"paymentType":"TWELVE_INSTALLMENTS","giftCard":"ABC123"}}' > /dev/null
	end=$(date +%s.%3N)
	duration=$( echo "$end - $start" | bc -l )
	echo "took ${duration}s"
done

package pl.tjug.rembol.commerce;

import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.client.RestTemplate;

import java.math.BigDecimal;
import java.util.Map;

@RestController
public class CommerceController {

    private RestTemplate restTemplate = new RestTemplate();

    @GetMapping("health")
    public String healthCheck() {
        return "ok";
    }

    @PostMapping("finalizePurchase")
    public Response finalizePurchase(@RequestBody Request request) {
        var cart = restTemplate.getForObject("http://carts:8080/carts/" + request.cartId, Cart.class);
        restTemplate.postForObject("http://delivery:8080/scheduleDelivery", cart.cart(), DeliveryStatus.class);
        restTemplate.postForObject("http://customers:8080/customers/"+request.customerId+"/registerOrder", cart, Void.class);
        restTemplate.postForObject("http://payments:8080/processPayment", request.paymentInfo(), Void.class);

        return new Response("ok");
    }

    @ExceptionHandler
    public ResponseEntity<Void> onException(Exception e) {
        return ResponseEntity.internalServerError().build();
    }

    public record DeliveryStatus(String status) {
    }

    public record Cart(Map<String, BigDecimal> cart, BigDecimal totalPrice){
    }

    public record Request(int cartId, int customerId, PaymentInfo paymentInfo) {
    }

    public record PaymentInfo(int creditCardId, PaymentType paymentType, String giftCard) {
    }

    public enum PaymentType {
        TWELVE_INSTALLMENTS, ONE_TIME
    }

    public record Response(String status) {
    }

}

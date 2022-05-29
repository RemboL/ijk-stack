const express = require('express');
const axios = require('axios');

const app = express();

const createCart = function(id) {
    return {
        'WashMaster 3000': (id % 9 + 1) % 3,
        'eyePhone 4FS': Math.floor((id % 9 + 1) / 3)
    }
};

const calculateTotalPrice = function(cart, products, specialOffers) {
    return Object.keys(cart)
        .map(item => (parseFloat(products[item] || "0") + parseFloat(specialOffers[item] || "0")) * parseInt(cart[item] || "0"))
        .reduce((a, b) => a+b);
};

app.route('/health').get(function(req, res) { 
    res.send('ok'); 
});
app.route('/carts/:id').get(function(req, res) { 
    axios.all([
        axios.get('http://products:8080/products'),
        axios.get('http://special-offers:8080/special-offers')
    ]).then(axios.spread((...responses) => {
        const products = responses[0].data;
        const specialOffers = responses[1].data;
        const cart = createCart(req.params.id);
        const totalPrice = calculateTotalPrice(cart, products, specialOffers);
        res.json({cart, totalPrice});
    })).catch(error => {
        res.status(500).send('');
    });
});


app.listen(8080);
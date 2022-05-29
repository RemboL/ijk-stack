const express = require('express');

const app = express();

const wait = function(millis) {
    var date = new Date();
    var curDate = null;
    do { curDate = new Date(); }
    while(curDate - date < millis);
}

const random = function(min, max) {
    min = Math.ceil(min);
    max = Math.floor(max);
    return Math.floor(Math.random() * (max - min)) + min;
}

app.route('/health').get(function(req, res) { 
    res.send('ok'); 
});

app.route('/customers/:id/consents').post(function(req, res) {
    wait(random(250, 450));
    res.json({status: 'ok'});
});

app.listen(8080);

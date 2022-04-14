var express = require('express');
var router = express.Router();

router.get('/', function(req, res, next) {
  res.render('index', { title: 'Home page' });
});

router.get('/register/', function(req, res, next) {
  res.render('account-staff/register', { title: 'Register', api_url: "http://localhost:80/" });
});

router.get('/login', function(req, res, next) {
  res.render('account-staff/login', { title: 'Login' });
});

router.get('/logout', function(req, res, next) {
  res.render('account-staff/logout', { title: 'Logout' });
});

module.exports = router;

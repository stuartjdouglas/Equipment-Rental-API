

// var domain = 'http://192.168.1.99/api'
// var domain = 'http://localhost:3000'
 var domain = 'http://localhost:3000'
//var domain = 'http://lemondev.xyz:3000'
var api = '/api';

var backend = domain + api;
data = domain;

angular.module('app.config', [])
.factory('Configuration', function() {
    return {
        backend: backend
    }
});

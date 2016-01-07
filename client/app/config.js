

// var domain = 'http://192.168.1.99/api'
var domain = 'http://localhost:3000/api'





var backend = domain;
angular.module('app.config', [])

.factory('Configuration', function() {
    return {
        backend: backend
    }
});

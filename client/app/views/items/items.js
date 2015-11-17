'use strict';

angular.module('app.items', ['ngRoute'])

.config(['$routeProvider',
    function($routeProvider) {
        $routeProvider.when('/p', {
            templateUrl: 'views/items/items.html',
            controller: 'itemsCtrl'
        });
    }
])
.controller('itemsCtrl', ['$rootScope', '$scope', '$http', function($rootScope, $scope, $http) {
    if ($rootScope.loggedIn) {
        $http({
            url: backend + "/p",
            method: 'GET',
        }).success(function(data, status, headers, config) {
            console.log(data);
            $scope.products = data;
        }).
        error(function(data, status, headers, config) {
            $scope.error = true;
        });
    } else {
        $scope.view = false;
    }


}]);

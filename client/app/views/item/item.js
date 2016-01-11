'use strict';

angular.module('app.item', ['ngRoute'])

.config(['$routeProvider',
    function($routeProvider) {
        $routeProvider.when('/listing/:id', {
            templateUrl: 'views/item/item.html',
            controller: 'itemCtrl'
        });
    }
])
.controller('itemCtrl', ['$rootScope', '$scope', '$http', '$routeParams', '$location', function($rootScope, $scope, $http, $routeParams, $location) {
        $http({
            url: backend + "/product/" + $routeParams.id,
            method: 'GET',
        }).success(function(data, status, headers, config) {
          if (data.items[0].owner.username === $rootScope.auth.username) {
            $location.path('/owner/listing/' + $routeParams.id);
          }
            $scope.product = data.items[0];
        }).
        error(function(data, status, headers, config) {
          console.log('error');
            $scope.error = true;
        });



}]);

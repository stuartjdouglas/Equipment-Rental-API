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
.controller('itemCtrl', ['$rootScope', '$scope', '$http', '$routeParams', function($rootScope, $scope, $http, $routeParams) {
  console.log( $routeParams.id);
        $http({
            url: backend + "/product/" + $routeParams.id,
            method: 'GET',
        }).success(function(data, status, headers, config) {
          console.log(data);
            $scope.product = data.item[0];
        }).
        error(function(data, status, headers, config) {
          console.log('error');
            $scope.error = true;
        });
    


}]);

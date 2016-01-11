'use strict';

angular.module('app.ownerItem', ['ngRoute'])

.config(['$routeProvider',
    function($routeProvider) {
        $routeProvider.when('/owner/listing/:id', {
            templateUrl: 'views/owner/item/item.html',
            controller: 'ownerItemCtrl'
        });
    }
])
.controller('ownerItemCtrl', ['$rootScope', '$scope', '$http', '$routeParams', function($rootScope, $scope, $http, $routeParams) {
        $http({
            url: backend + "/product/" + $routeParams.id,
            method: 'GET',
        }).success(function(data, status, headers, config) {
            $scope.product = data.items[0];
        }).
        error(function(data, status, headers, config) {
            $scope.error = true;
        });



}]);

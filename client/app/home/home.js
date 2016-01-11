'use strict';

angular.module('app.home', ['ngRoute'])

.config(['$routeProvider', function($routeProvider) {
  $routeProvider.when('/home/', {
    templateUrl: 'home/home.html',
    controller: 'HomeCtrl'
  });
}])

.controller('HomeCtrl', ['$rootScope', '$scope', '$http', 'authFactory', function($rootScope, $scope, $http, authFactory) {
	if ($rootScope.loggedIn) {
        	$scope.view = true;

            $http({
                url: backend + "/hello",
                method: 'GET',
                headers: {
                    'Content-Type': 'multipart/form-data',
                    'token': authFactory.getToken()
                }
            }).success(function(data, status, headers, config) {
            	$scope.message = data.message;
            }).
            error(function(data, status, headers, config) {

            });
        } else {
        	$scope.view = false;
        }
}]);

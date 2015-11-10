'use strict';

angular.module('app.home', ['ngRoute'])

.config(['$routeProvider', function($routeProvider) {
  $routeProvider.when('/home/', {
    templateUrl: 'home/home.html',
    controller: 'HomeCtrl'
  });
}])

.controller('HomeCtrl', ['$rootScope', '$scope', '$http', function($rootScope, $scope, $http) {
	if ($rootScope.loggedIn) {
        	$scope.view = true;

            $http({
                url: backend + "/hello",
                method: 'GET',
                headers: {
                    'Content-Type': 'multipart/form-data',
                    'token': window.sessionStorage.token
                }
            }).success(function(data, status, headers, config) {
            	$scope.message = data.message;
                console.log(data);
            }).
            error(function(data, status, headers, config) {
                console.log(data);



            });
        } else {
        	$scope.view = false;
        }
}]);

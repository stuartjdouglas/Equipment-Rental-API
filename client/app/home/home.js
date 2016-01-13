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
        $scope.showHello = false;
        sayHello();

        $scope.showTimeline = false;
        getTimeline();

    } else {
        $scope.view = false;
    }

    function getTimeline() {
        $http({
            url: backend + "/p/rent/current",
            method: 'GET',
            headers: {
                'Count': 5,
                'token': authFactory.getToken()
            }
        }).success(function(data, status, headers, config) {
            $scope.timeline = data;
            $scope.showTimeline = true;
        }).
        error(function(data, status, headers, config) {
            $scope.error = true;
        });
    }

    function sayHello() {
        $http({
            url: backend + "/hello",
            method: 'GET',
            headers: {
                'Content-Type': 'multipart/form-data',
                'token': authFactory.getToken()
            }
        }).success(function(data, status, headers, config) {
            $scope.message = data.message;
            $scope.showHello = true;
        }).
        error(function(data, status, headers, config) {

        });
    }
}]);

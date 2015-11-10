'use strict';

angular.module('app.users', ['ngRoute'])

    .config(['$routeProvider',
        function($routeProvider) {
            $routeProvider.when('/users', {
                templateUrl: 'views/users/users.html',
                controller: 'UsersCtrl'
            });
        }
    ])
    .controller('UsersCtrl', ['$rootScope', '$scope', '$http', function($rootScope, $scope, $http) {
        if ($rootScope.loggedIn) {
            $scope.view = true;

            $http({
                url: backend + "/users",
                method: 'GET'
            }).success(function(data, status, headers, config) {
                $scope.users = data;
            }).error(function(data, status, headers, config) {
                });
        } else {
            $scope.view = false;
        }
    }]);

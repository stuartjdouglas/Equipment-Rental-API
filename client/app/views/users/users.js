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
    .controller('UsersCtrl', ['$cookies', '$scope', '$http', function($cookies, $scope, $http) {
        if ($cookies.get('token')) {
            $scope.view = true;

            $http({
                url: backend + "/users",
                method: 'GET'
            }).success(function(data, status, headers, config) {
                console.log(data);
                $scope.users = data;
            }).
                error(function(data, status, headers, config) {
                    console.log(data);



                });
        } else {
            $scope.view = false;
        }
    }]);

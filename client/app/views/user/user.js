'use strict';

angular.module('app.user', ['ngRoute'])

    .config(['$routeProvider',
        function($routeProvider) {
            $routeProvider.when('/user/:user', {
                templateUrl: 'views/user/user.html',
                controller: 'UserCtrl'
            });
        }
    ])
    .controller('UserCtrl', ['$routeParams', '$scope', '$http', function($routeParams, $scope, $http) {

            $scope.view = true;
            $scope.query = $routeParams.user;
            $http({
                url: backend + "/user/" + $routeParams.user,
                method: 'GET'
            }).success(function(data, status, headers, config) {
                console.log(data);
                $scope.user = data[0];
            }).
                error(function(data, status, headers, config) {
                    console.log(data);



                });

    }]);

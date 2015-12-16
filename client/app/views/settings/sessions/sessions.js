'use strict';

angular.module('app.sessions', ['ngRoute'])

    .config(['$routeProvider',
        function($routeProvider) {
            $routeProvider.when('/profile/sessions', {
                templateUrl: 'views/settings/sessions/sessions.html',
                controller: 'SessionsCtrl'
            });
        }
    ])
    .controller('SessionsCtrl', ['$rootScope', '$scope', '$http', function($rootScope, $scope, $http) {
        if ($rootScope.loggedIn) {
            console.log("hello");
            $http({
                url: backend + "/profile/sessions",
                method: 'GET',
                headers: {
                    'token': window.sessionStorage.token
                }
            }).success(function(data, status, headers, config) {
                console.log(data);
                $scope.sessions = data;
            }).
            error(function(data, status, headers, config) {
                $scope.error = true;
            });


        } else {
            $scope.view = false;
        }


    }]);

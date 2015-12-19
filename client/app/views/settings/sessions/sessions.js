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
            getSessions();


        } else {
            $scope.view = false;
        }

        function getSessions() {
            $http({
                url: backend + "/profile/sessions",
                method: 'GET',
                headers: {
                    'token': window.sessionStorage.token
                }
            }).success(function(data, status, headers, config) {
                $scope.sessions = data;
            }).
            error(function(data, status, headers, config) {
                $scope.error = true;
            });
        }

        $scope.disable = function(index) {
            $http({
                url: backend + "/session",
                method: 'DELETE',
                headers: {
                    'token': window.sessionStorage.token,
                    'id': index
                }
            }).success(function(data, status, headers, config) {
                getSessions()
            }).
            error(function(data, status, headers, config) {
                $scope.error = true;
            });
        }


    }]);

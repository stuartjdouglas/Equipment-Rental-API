'use strict';

angular.module('app.profile', ['ngRoute'])

.config(['$routeProvider',
    function($routeProvider) {
        $routeProvider.when('/profile/', {
            templateUrl: 'views/profile/profile.html',
            controller: 'ProfileCtrl'
        });
    }
])
.controller('ProfileCtrl', ['$rootScope', '$scope', '$http', function($rootScope, $scope, $http) {
    if ($rootScope.loggedIn) {

        $http({
            url: backend + '/profile',
            method: 'GET',
            headers: {
                'Content-Type': 'multipart/form-data',
                'token': window.sessionStorage.token
            }
        }).success(function(data, status, headers, config) {
            $scope.profile = data.profile;
            $scope.view = true;
            $http({
                url: backend + "/identify/qr/user",
                method: 'GET',
                headers: {
                    'width': 300,
                    'height': 300,
                    'code': data.profile.username
                }
            }).success(function(data, status, headers, config) {

                $scope.qr = data;
            }).
            error(function(data, status, headers, config) {
                $scope.error = true;
            });
        }).
        error(function(data, status, headers, config) {
            $scope.error = true;
        });



    } else {
        $scope.view = false;
    }


}]);

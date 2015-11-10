'use strict';

angular.module('app.images', ['ngRoute'])

    .config(['$routeProvider',
        function($routeProvider) {
            $routeProvider.when('/images', {
                templateUrl: 'views/images/images.html',
                controller: 'ImagesCtrl'
            });
        }
    ])
    .controller('ImagesCtrl', ['$routeParams', '$scope', '$http', '$rootScope', function($routeParams, $scope, $http, $rootScope) {

        if ($rootScope.loggedIn) {
            $scope.view = true;

            $http({
                url: backend + "/images",
                method: 'GET',
                headers: {
                    'Content-Type': 'multipart/form-data',
                    'token': window.sessionStorage.token
                }
            }).success(function(data, status, headers, config) {
                $scope.images = data;
                console.log(data);
            }).
                error(function(data, status, headers, config) {
                    console.log(data);



                });
        } else {
            $scope.view = false;
        }
    }]);

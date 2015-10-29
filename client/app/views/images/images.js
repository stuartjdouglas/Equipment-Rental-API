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
    .controller('ImagesCtrl', ['$routeParams', '$scope', '$http', '$cookies', function($routeParams, $scope, $http, $cookies) {

        if ($cookies.get('token')) {
            $scope.view = true;

            $http({
                url: backend + "/images",
                method: 'GET',
                headers: {
                    'Content-Type': 'multipart/form-data',
                    'token': $cookies.get('token')
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

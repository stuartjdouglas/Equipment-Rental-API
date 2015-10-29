'use strict';

angular.module('app.image', ['ngRoute'])

    .config(['$routeProvider',
        function($routeProvider) {
            $routeProvider.when('/image/:filename', {
                templateUrl: 'views/image/image.html',
                controller: 'ImageCtrl'
            });
        }
    ])
    .controller('ImageCtrl', ['$routeParams', '$scope', '$http', function($routeParams, $scope, $http) {

        $scope.view = true;
        $scope.query = $routeParams.filename;
        $http({
            url: backend + "/image/" + $routeParams.filename,
            method: 'GET'
        }).success(function(data, status, headers, config) {
            console.log(data.image);
            $scope.images = data.image;
        }).
            error(function(data, status, headers, config) {
                console.log(data);



            });

    }]);

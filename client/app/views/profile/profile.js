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
.controller('ProfileCtrl', ['$cookies', '$scope', '$http', function($cookies, $scope, $http) {
	if ($cookies.get('token')) {
        	$scope.view = true;

            $http({
                url: backend + "/profile",
                method: 'GET',
                headers: {
                    'Content-Type': 'multipart/form-data',
                    'token': $cookies.get('token')
                }
            }).success(function(data, status, headers, config) {
            	$scope.profile = data.profile[0];
                console.log(data);
            }).
            error(function(data, status, headers, config) {
                console.log(data);



            });
        } else {
        	$scope.view = false;
        }
}]);

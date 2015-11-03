'use strict';

angular.module('app.logout', ['ngRoute'])

    .config(['$routeProvider', function($routeProvider) {
        $routeProvider.when('/logout', {
            templateUrl: 'views/logout/logout.html',
            controller: 'LogoutCtrl'
        });
    }])

    .controller('LogoutCtrl', ['$scope', '$rootScope', function($scope, $rootScope) {
        $scope.view = $rootScope.loggedIn;

        if ($rootScope.loggedIn) {
            window.sessionStorage.removeItem("token");
            $rootScope.loggedIn = false;
        }
    }]);
'use strict';

angular.module('app.logout', ['ngRoute'])

    .config(['$routeProvider', function($routeProvider) {
        $routeProvider.when('/logout', {
            templateUrl: 'views/logout/logout.html',
            controller: 'LogoutCtrl'
        });
    }])

    .controller('LogoutCtrl', ['$scope', '$rootScope', 'authFactory', function($scope, $rootScope, authFactory) {
        $scope.view = $rootScope.loggedIn;

        if ($rootScope.loggedIn) {
            // window.sessionStorage.remov0eItem("token");
            authFactory.killAuth();
            $rootScope.loggedIn = false;
        }
    }]);

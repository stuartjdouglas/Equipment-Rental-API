'use strict';

// Declare app level module which depends on views, and components
angular.module('app', [
  'ngRoute',
  'app.home',
  'app.login',
  'app.profile',
  'app.register',
    'app.users',
    'app.user',
    'app.image',
    'app.images',
  // Directives
  'app.loginPanel',
  'app.registerPanel',
  // Factories
  'app.config',
  'ngCookies'
]).
config(['$routeProvider','$locationProvider', function($routeProvider, $locationProvider) {
	$routeProvider.otherwise({redirectTo: '/home'});


	$locationProvider.html5Mode(false);
}])

.controller('AuthCtrl', ['$cookies', '$scope', '$rootScope', function($cookies, $scope, $rootScope) {
	$rootScope.loggedIn = $cookies.get('token') != undefined;
}]);

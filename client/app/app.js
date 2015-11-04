'use strict';

// Declare app level module which depends on views, and components
angular.module('app', [
  'ngRoute',
  'app.home',
  'app.login',
    'app.logout',
  'app.profile',
  'app.register',
    'app.users',
    'app.user',
    'app.image',
    'app.images',
    'app.imageupload',
  // Directives
  'app.loginPanel',
  'app.registerPanel',
    'app.imageUploadForm',
  // Factories
  'app.config',
  //  Dependencies
    'ngFileUpload',
    'ngStorage',
  'ngCookies',
    'angularMoment'
]).
config(['$routeProvider','$locationProvider', function($routeProvider, $locationProvider) {
	$routeProvider.otherwise({redirectTo: '/home'});
	$locationProvider.html5Mode(false);
}])

.controller('AuthCtrl', ['$sessionStorage', '$scope', '$rootScope', function($sessionStorage, $scope, $rootScope) {
	$rootScope.loggedIn = window.sessionStorage.token != undefined;
}]);

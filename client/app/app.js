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
  'app.item',
  'app.items',
  'app.sessions',
  'app.myItems',

  // Directives
  'app.loginPanel',
  'app.registerPanel',
  'app.imageUploadForm',
  'app.rlabel',
  'app.availability',
  'app.rentButton',
  // Factories
  'app.config',
  //  Dependencies
  'angularMoment',
  'ui-notification'
]).
config(['$routeProvider','$locationProvider', function($routeProvider, $locationProvider) {
	$routeProvider.otherwise({redirectTo: '/home'});
	$locationProvider.html5Mode(false);
}])

.controller('AuthCtrl', ['$scope', '$rootScope', function($scope, $rootScope) {
  $rootScope.loggedIn = window.sessionStorage.token !== undefined;
  $rootScope.auth = {
    username: window.localStorage.username,
    gravatar: window.localStorage.gravatar
  };

  $rootScope.site = {
    title: "Site name"
  };
}]);

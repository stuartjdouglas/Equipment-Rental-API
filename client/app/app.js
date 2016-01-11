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
  'app.creator',
  'app.ownerItems',
  'app.ownerItem',

  // Directives
  'app.loginPanel',
  'app.registerPanel',
  'app.imageUploadForm',
  'app.rlabel',
  'app.availability',
  'app.rentButton',
  'app.rentButtonOwner',
  'app.qrcode',
  'app.limage',
  // Factories
  'app.config',
  'app.auth',
  //  Dependencies
  'angularMoment',
  'ui-notification',
  'naif.base64'
]).
config(['$routeProvider','$locationProvider', function($routeProvider, $locationProvider) {
	$routeProvider.otherwise({redirectTo: '/home'});
	$locationProvider.html5Mode(false);
}])

.controller('AuthCtrl', ['$scope', '$rootScope', 'authFactory', function($scope, $rootScope, authFactory) {
  $rootScope.loggedIn = authFactory.getAuth() !== undefined;
  // console.log( authFactory.getToken);
  $rootScope.auth = authFactory.getAuth();

  $rootScope.site = {
    title: "Site name"
  };
}]);

'use strict';

angular.module('app.cookie', ['ngRoute'])

    .config(['$routeProvider', function($routeProvider) {
        $routeProvider.when('/cookie', {
            templateUrl: 'views/cookie/cookie.html',
            controller: 'cookieCtrl'
        });
    }])

    .controller('cookieCtrl', ['$scope', '$rootScope', 'authFactory', function($scope, $rootScope, authFactory) {
        $scope.disable = function() {
            authFactory.RemoveSessCookie();
        }


        $scope.content = {
            title: "cookie",
            content: [{
                "title": "What do we store?",
                "content": "We store your username, token and gravatar id in the cookie"
            }, {
                "title": "Why do you need this?",
                "content": "To be able to access your user data and perform authorized calls to our services we require you to have an access token to identify you"
            }, {
                "title": "Are the cookies used elsewhere",
                "content": "No, we don't use the cookies for anything else, nor do we accept or allow other cookies to be used on this site."
            }]
        };
    }]);

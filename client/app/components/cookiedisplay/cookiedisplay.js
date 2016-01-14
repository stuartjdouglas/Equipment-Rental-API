'use strict';

angular.module('app.cookiedisplay', ['app.config', 'app.auth'])
    .directive('cookiedisplay', function() {
        return {
            restrict: 'AEC',
            scope: {
                url: '@'
            },
            templateUrl: 'components/cookiedisplay/cookiedisplay.html',
            controller: function($scope, $colorThief, authFactory, $rootScope) {
                var cookie = authFactory.usingCookies();

                $rootScope.$watch('enableCookieSession', function() {
                    if ($rootScope.enableCookieSession) {
                        createCookie();
                    }
                });

                if (cookie && cookie.active) {
                    $scope.displaycookiebar = false;
                    $rootScope.noCookieUsage = true;
                    $rootScope.enableCookieSession = true;
                } else {
                    $scope.displaycookiebar = true;
                    $rootScope.noCookieUsage = false;
                    $rootScope.enableCookieSession = false;
                }

                $scope.agreeCookie = function() {
                    createCookie();
                }

                function createCookie() {
                    $scope.animation = 'slideOutDown';
                    authFactory.createCookieSess();
                    $rootScope.noCookieUsage = true;
                    setTimeout(function() {
                        $scope.displaycookiebar = false;
                    }, 1000);
                }

                $scope.declineCookie = function() {
                    $scope.animation = 'slideOutDown';
                    authFactory.RemoveSessCookie();
                    $rootScope.noCookieUsage = false;
                    setTimeout(function(){
                        $scope.displaycookiebar = false;
                    }, 1000);
                };
            },
            link: function(scope, elem, attrs, http, authFactory) {

            }
        };
    });

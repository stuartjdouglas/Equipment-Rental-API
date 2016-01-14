'use strict';

angular.module('app.loginPanel', ['app.config'])
    .directive('loginPanel', function() {
        return {
            restrict: 'AEC',
            scope: {
                datasource: '='
            },
            templateUrl: 'components/loginPanel/loginPanel.html',
            controller: function($scope, $http, $rootScope, $location, authFactory) {

                $rootScope.$watch('noCookieUsage', function() {
                    $scope.disable = !$rootScope.noCookieUsage;
                });

                $scope.enableCookie = function() {
                    console.log("enable cookies");
                    $rootScope.enableCookieSession = true;
                }

                $scope.login = function(user) {
                    $scope.showError = false;
                    if (user !== undefined && user !== "" && user.name !== "" && user.password !== "") {
                        var hash = CryptoJS.SHA512(user.password).toString();

                        $http({
                            url: backend + "/login",
                            method: 'POST',
                            data: {
                                'username': user.name,
                                'password': hash
                            },
                            headers: {
                                'Content-Type': 'multipart/form-data'

                            }
                        }).success(function(data, status, headers, config) {
                            authFactory.setAuth(data.token, data.username, data.gravatar, Date.parse(data.expiry))
                            $rootScope.auth = authFactory.getAuth();
                            $scope.error = false;
                            $rootScope.loggedIn = true;
                            $location.path( "/home");
                        }).
                        error(function(data, status, headers, config) {
                            $scope.showError = true;
                            $scope.error = data.message;
                        });
                    } else {
                        $scope.showError = true;
                        $scope.error = "You are missing values";
                    }


                };
            },
            link: function(scope, elem, attrs) {
                // Just for altering the DOM
            }
        };
    });

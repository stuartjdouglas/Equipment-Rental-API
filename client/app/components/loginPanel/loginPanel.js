'use strict';

angular.module('app.loginPanel', ['app.config'])
    .directive('loginPanel', function() {
        return {
            restrict: 'AEC',
            scope: {
                datasource: '='
            },
            templateUrl: 'components/loginPanel/loginPanel.html',
            controller: function($scope, $http, $rootScope, $location) {
                
                $scope.login = function(user) {

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
                        //$sessionStorage.token = data.token;
                        window.sessionStorage.token = data.token;
                        window.localStorage.username = data.username;
                        window.localStorage.gravatar = data.gravatar;
                        $scope.error = false;
                        $rootScope.loggedIn = true;
                        $location.path( "/home");
                    }).
                    error(function(data, status, headers, config) {
                        $scope.error = data.message;


                    });

                }
            },
            link: function(scope, elem, attrs) {
                // Just for altering the DOM
            }
        };
    });
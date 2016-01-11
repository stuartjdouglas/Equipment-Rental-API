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
                        console.log(data);
                        authFactory.setAuth(data.token, data.username, data.gravatar, Date.parse(data.expiry))
                        $rootScope.auth = authFactory.getAuth();
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

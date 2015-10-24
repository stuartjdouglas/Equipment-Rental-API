'use strict';

angular.module('app.loginPanel', ['app.config'])
    .directive('loginPanel', function() {
        return {
            restrict: 'AEC',
            scope: {
                datasource: '='
            },
            templateUrl: 'components/loginPanel/loginPanel.html',
            controller: function($scope, $http, $rootScope) {
                
                $scope.login = function(user) {
                    console.log(user);
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
                          document.cookie="token=" + data.token + "; expires" + data.expiry + " path=/";
                          $scope.error = false;
                        console.log(data);
                        $rootScope.loggedIn = true;
                    }).
                    error(function(data, status, headers, config) {
                        console.log(data);
                        $scope.error = data.message;


                    });

                }
            },
            link: function(scope, elem, attrs) {
                // Just for altering the DOM
            }
        };
    });
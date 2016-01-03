'use strict';

angular.module('app.registerPanel', ['app.config'])
    .directive('registerPanel', function() {
        return {
            restrict: 'AEC',
            scope: {
                datasource: '='
            },
            templateUrl: 'components/registerPanel/registerPanel.html',
            controller: function($scope, $http) {
              $scope.$watch(
                "user.email",
                function handleFooChange( newValue, oldValue ) {
                  var hash = CryptoJS.MD5(newValue);
                  $scope.gravatar = hash.toString();
                }
              );
                $scope.register = function(user) {
                    console.log(user);
                    if (user.username != "" && user.email != "" && user.password != "" &&
                    user.username != " " && user.email != " " && user.password.length < 6) {
                        var hash = CryptoJS.SHA512(user.password).toString();

                        $http({
                            url: backend + "/user/register",
                            method: 'POST',
                            data: {
                                'username': user.name,
                                'password': hash,
                                'email':user.email
                            },
                            headers: {
                                'Content-Type': 'multipart/form-data'

                            }
                        }).success(function(data, status, headers, config) {

                            $scope.error = false;
                            console.log(data);
                        }).
                        error(function(data, status, headers, config) {
                            console.log(data);
                            $scope.error = data.message;


                        });
                    }

                    console.log(user);


                }
            },
            link: function(scope, elem, attrs) {
                // Just for altering the DOM
            }
        };
    });

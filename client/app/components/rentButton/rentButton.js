'use strict';

angular.module('app.rentButton', ['app.config'])
    .directive('rentButton', function() {
        return {
            restrict: 'AEC',
            scope: {
                datasource: '='
            },
            templateUrl: 'components/rentButton/rentButton.html',
            controller: function($scope, $http, $rootScope, $location, $attrs, Notification) {
                //$scope.datasource =  $attrs.datasource;
                $scope.availability = "Loading.....";
                $scope.rentButtonClass = [];
                //$scope.rentButtonClass.push('fa fa-spinner');

                $scope.$watch(
                    "datasource",
                    function handleFooChange(oldValue, newValue) {
                        console.log($scope.datasource)
                        if ($scope.datasource != undefined) {
                            if ($scope.datasource.id != undefined) {
                                $scope.showLoading = false;
                                //console.log(">>> " + $scope.datasource)
                                // var headers = {};
                                // Check if we are a logged in user or not
                                getRentalStatus();


                            } else {
                                $scope.showLoading = true;
                            }
                        }
                    }
                );

                $scope.click = function(id) {
                    if ($scope.availability === 'Unavailable') {
                        Notification.error({message: '<i class="fa fa-exclamation-triangle"></i> ' + $scope.datasource.title + ' is not available. :(', positionY: 'bottom', positionX: 'center'});
                    } else {
                        rent(id);

                    }
                }

                function getRentalStatus() {
                    if ($rootScope.loggedIn) {
                        console.log("logged in user");
                        console.log( window.sessionStorage.token);
                        $http({
                            url: backend + '/p/' + $scope.datasource.id + '/availability',
                            method: 'GET',
                            headers: {
                                'token': window.sessionStorage.token
                            },
                        }).success(function (data, status, headers, config) {
                            $scope.result = data;
                            console.log(data);
                            $scope.gotRes = true;
                            $scope.rentButtonClass.splice("", 0);
                            if (data.available) {
                                if ($scope.gotRes) {
                                    $scope.availability = "Available";
                                    $scope.rentButtonClass.push('button-primary');
                                }
                            } else {
                                if (!data.available) {
                                    if ($scope.gotRes) {
                                        console.log(data.owner);
                                        if (data.owner) {
                                            $scope.owner = true;
                                            $scope.availability = 'You current own this';
                                        } else {
                                            $scope.availability = "Unavailable";
                                        }

                                    }
                                }
                            }
                        }).
                        error(function (data, status, headers, config) {
                            $scope.error = true;
                        });
                    } else {
                        console.log("unlogged in user");
                        $http({
                            url: backend + '/p/' + $scope.datasource.id + '/availability',
                            method: 'GET',
                            // header: headers,
                        }).success(function (data, status, headers, config) {
                            $scope.gotRes = true;
                            $scope.rentButtonClass.splice("", 0);
                            if (data.available) {
                                if ($scope.gotRes) {
                                    $scope.availability = "Available";
                                    $scope.rentButtonClass.push('button-primary');
                                }
                            } else {
                                if (!data.available) {
                                    if ($scope.gotRes) {
                                        $scope.availability = "Unavailable";

                                    }
                                }
                            }
                        }).
                        error(function (data, status, headers, config) {
                            $scope.error = true;
                        });
                    }
                }

                function rent(id) {
                    $http({
                        url: backend + '/p/' + $scope.datasource.id + '/rent',
                        method: 'POST',
                        headers: {
                            'token': window.sessionStorage.token
                        },
                    }).success(function (data, status, headers, config) {
                        Notification.success({message: '<i class="fa fa-paper-plane"></i> ' + $scope.datasource.title + ' has just been rented. :)', positionY: 'bottom', positionX: 'center'});
                        $scope.owner = true;
                        getRentalStatus();
                        $scope.rentButtonClass = [];
                    }).
                    error(function (data, status, headers, config) {
                        $scope.error = true;
                    });
                }

                $scope.return = function(id) {
                    console.log(id);
                    $http({
                        url: backend + '/p/' + $scope.datasource.id + '/return',
                        method: 'POST',
                        headers: {
                            'token': window.sessionStorage.token
                        },
                    }).success(function (data, status, headers, config) {
                        getRentalStatus();
                        $scope.owner = false;
                    }).
                    error(function (data, status, headers, config) {
                        $scope.error = true;
                    });
                }
            },

            link: function(scope, elem, attrs, http) {


            }
        };
    });

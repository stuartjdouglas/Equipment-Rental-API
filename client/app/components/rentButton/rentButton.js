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

                        if ($scope.datasource.id != undefined) {
                            $scope.showLoading = false;
                            //console.log(">>> " + $scope.datasource)
                            $http({
                                url: backend + '/p/' + $scope.datasource.id + '/availability',
                                method: 'GET',
                            }).success(function(data, status, headers, config) {
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
                            error(function(data, status, headers, config) {
                                $scope.error = true;
                            });
                        } else {
                            $scope.showLoading = true;
                        }
                    }
                );

                $scope.click = function(id) {
                    if ($scope.availability = 'Unavailable') {
                        Notification.error({message: '<i class="fa fa-exclamation-triangle"></i> ' + $scope.datasource.title + ' is not available. :(', positionY: 'bottom', positionX: 'center'});
                    }
                }
            },

            link: function(scope, elem, attrs, http) {


            }
        };
    });
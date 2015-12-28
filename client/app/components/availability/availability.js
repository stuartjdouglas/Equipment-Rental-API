'use strict';

angular.module('app.availability', ['app.config'])
    .directive('availability', function() {
        return {
            restrict: 'AEC',
            templateUrl: 'components/availability/availability.html',
            controller: function($scope, $http, $rootScope, $location, $attrs) {
                $scope.datasource =  $attrs.datasource;
                $scope.availability = "Loading...";
                $scope.$watch(
                    "datasource",
                    function handleFooChange( ) {
                        if ($attrs.datasource != "{{product.id}}") {
                            $scope.showLoading = false;
                            //$scope.datasource =  $attrs.datasource;
                            console.log('/p/' + $attrs.datasource + '/availability');
                            $http({
                                url: backend + '/p/' + $attrs.datasource + '/availability',
                                method: 'GET',
                            }).success(function(data, status, headers, config) {
                                if (data.available) {
                                    $scope.availability = "Available";
                                } else {
                                    $scope.availability = "Unavailable";
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




            },

            link: function(scope, elem, attrs, http) {

            }
        };
    });
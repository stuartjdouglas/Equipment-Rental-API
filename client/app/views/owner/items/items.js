'use strict';

angular.module('app.ownerItems', ['ngRoute'])

    .config(['$routeProvider',
        function($routeProvider) {
            $routeProvider.when('/owner/listing', {
                templateUrl: 'views/owner/items/items.html',
                controller: 'ownerItemListingCtrl'
            });
        }
    ])
    .controller('ownerItemListingCtrl', ['$rootScope', '$scope', '$http', 'authFactory', function($rootScope, $scope, $http, authFactory, $watch) {
        if (window.localStorage.getItem("product_count")) {
            $scope.count = parseInt(window.localStorage.getItem("product_count"));
        } else {
            window.localStorage.setItem("product_count", 2);
            $scope.count = 2;
        }


        $scope.start = 0;



        if ($rootScope.loggedIn) {
            updateResults();
        } else {
            $scope.view = false;
        }

        $scope.$watch("count", function(newValue) {
            window.localStorage.setItem("product_count", newValue);
            updateResults();
        });

        $scope.back = function() {
            //$scope.start + $scope.count > $scope.products.total

            if ($scope.start <= 0) {
                $scope.viewResults = false;
            } else {
                $scope.viewResults = true;
                $scope.start = $scope.start - $scope.count;
                updateResults();
            }
        }

        $scope.forward = function() {
            if ($scope.start >= $scope.products.total - $scope.count) {
                $scope.viewResults = false;
            } else {
                $scope.viewResults = true;
                $scope.start = $scope.start + $scope.count;
                updateResults();
            }
        }

        function updateResults() {
            $http({
                url: backend + "/owner/products",
                method: 'GET',
                headers: {
                    'token': authFactory.getToken(),
                    'Start':$scope.start,
                    'Count':$scope.count
                }
            }).success(function(data, status, headers, config) {
                $scope.products = data;
            }).
            error(function(data, status, headers, config) {
                $scope.error = true;
            });
        }


    }]);

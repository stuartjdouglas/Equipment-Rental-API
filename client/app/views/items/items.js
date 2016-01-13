'use strict';

angular.module('app.items', ['ngRoute'])

.config(['$routeProvider',
    function($routeProvider) {
        $routeProvider.when('/listing', {
            templateUrl: 'views/items/items.html',
            controller: 'itemsCtrl'
        });
    }
])
.controller('itemsCtrl', ['$rootScope', '$scope', '$http', '$location', function($rootScope, $scope, $http, $location, $watch) {
    if (window.localStorage.getItem("product_count")) {
        $scope.count = parseInt(window.localStorage.getItem("product_count"));
    } else {
        window.localStorage.setItem("product_count", 10);
        $scope.count = 10;
    }

    $scope.goto = function(id) {
        $location.path('/listing/' + id);
    }


    $scope.start = 0;



    if ($rootScope.loggedIn) {
        //updateResults();
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
            url: backend + "/p",
            method: 'GET',
            headers: {
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

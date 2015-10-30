'use strict';

angular.module('app.imageUploadForm', ['app.config'])
    .directive('imageUploadForm', function() {
        return {
            restrict: 'AEC',
            scope: {
                datasource: '='
            },
            templateUrl: 'components/imageUploadForm/imageUploadForm.html',
            controller: function($scope, $http, $rootScope, $cookies) {
                $scope.data = false;



                $scope.upload = function(file) {
                    var fd = new FormData();
                    fd.append('image', file);
                    $http({
                        url: backend + "/image/upload",
                        method: 'POST',
                        dataType: 'multipart/form-data',
                        data: fd,
                        transformRequest: angular.identity,
                        headers: {
                            'Content-Type': undefined,
                            'token': $cookies.get('token')
                        }
                    }).success(function(data, status, headers, config) {
                        $scope.success = true;
                    }).
                        error(function(data, status, headers, config) {

                            $scope.success = false;

                        });
                }
            },
            link: function(scope, elem, attrs) {
                // Just for altering the DOM
            }
        };
    });
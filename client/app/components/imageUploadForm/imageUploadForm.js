'use strict';

angular.module('app.imageUploadForm', ['app.config'])
    .directive('imageUploadForm', function() {
        return {
            restrict: 'AEC',
            scope: {
                datasource: '='
            },
            templateUrl: 'components/imageUploadForm/imageUploadForm.html',
            controller: function($scope, $http, $rootScope, $location) {
                $scope.data = false;



                $scope.upload = function(file) {
                    console.log(file);
                    //if ($rootScope.loggedIn) {
                    //    $scope.view = true;
                    //    $scope.uploading = true;
                    //    var fd = new FormData();
                    //    fd.append('image', file);
                    //    $http({
                    //        url: backend + "/image/upload",
                    //        method: 'POST',
                    //        dataType: 'multipart/form-data',
                    //        data: fd,
                    //        transformRequest: angular.identity,
                    //        headers: {
                    //            'Content-Type': undefined,
                    //            'token': window.sessionStorage.token
                    //        }
                    //    }).success(function(data, status, headers, config) {
                    //        $scope.success = true;
                    //        $location.path( "/image/" + data.image[0].location );
                    //    }).
                    //    error(function(data, status, headers, config) {
                    //
                    //        $scope.success = false;
                    //
                    //    });
                    //} else {
                    //    $scope.view = false;
                    //}

                }
            },
            link: function(scope, elem, attrs) {
                // Just for altering the DOM
            }
        };
    });
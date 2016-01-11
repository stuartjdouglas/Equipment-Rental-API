'use strict';

angular.module('app.creator', ['ngRoute'])

.config(['$routeProvider',
    function($routeProvider) {
        $routeProvider.when('/create', {
            templateUrl: 'views/item/creator/creator.html',
            controller: 'itemCreatorCtrl'
        });
    }
])
.controller('itemCreatorCtrl', ['$rootScope', '$scope', '$http', '$routeParams', function($rootScope, $scope, $http, $routeParams) {
  $scope.product = {};
  $scope.message = {
    button: 'Create',
    submit: 'loading',
    loading: 'processing',
    failed: 'failed',
    error: {
      image:{
        enable: false,
        text:'image cannot be empty'
      },
      title:{
        enable:false,
        text:'title cannot be empty'
      },
      description:{
        enable: false,
        text: 'description cannot be empty'
      },
      days:{
        enable: false,
        text: 'days cannot be empty'
      },
    }
  }
  $scope.uploadButtonClass = [];

    if ($rootScope.loggedIn) {
      $scope.fileChange = function(data) {
        var file = document.getElementById('fileUploader').files[0];
        var preview = document.getElementById('imagePreview');
        if (file) {
          var reader  = new FileReader();
          reader.onloadend = function () {
              preview.src = reader.result;
          }
          reader.readAsDataURL(file)
        }
      }

      $scope.create = function(product) {
        // console.log(product != null);
        // console.log(product.base64 != null);
        // console.log(product.title != null);
        // console.log(product.description != null);
        // console.log(product.days != null);

        if (!$scope.uploadSuccess) {


          if (!product.base64) {
            $scope.message.error.image.enable = true;
          } else {
            $scope.message.error.image.enable = false;
          }

          if (!product.title) {
            $scope.message.error.title.enable = true;
          } else {
            $scope.message.error.title.enable = false;
          }

          if (!product.description) {
            $scope.message.error.description.enable = true;
          } else {
            $scope.message.error.description.enable = false;
          }

          if (!product.days) {
            $scope.message.error.days.enable = true;
          } else {
            $scope.message.error.days.enable = false;
          }


          if (product) {
            if (product.base64 && product.title && product.description && product.days) {
              console.log(product);
              // $scope.uploadButtonClass.push('button-primary');
              $scope.message.button = $scope.message.loading;

              var fd = new FormData();
              fd.append('title', product.title);
              fd.append('description', product.description);
              fd.append('rental_period_limit', product.days);
              fd.append('image', product.base64[0].base64);
              fd.append('filetype', product.base64[0].filetype);

              $http({
                url: backend + "/p",
                method: 'POST',
                dataType: 'multipart/form-data',
                data: fd,
                transformRequest: angular.identity,
                headers: {
                  'Content-Type': undefined,
                  'token': window.sessionStorage.token,
                }
              }).success(function (data, status, headers, config) {
                // $scope.message.button = $scope.uploadSuccess = true;
                // $scope.success = true;
              }).error(function (data, status, headers, config) {

                // $scope.success = false;

              });
            }
          }
        }
      }
    } else {

    }


}]);

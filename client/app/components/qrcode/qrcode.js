'use strict';

angular.module('app.qrcode', ['app.config'])
  .directive('qrcode', function() {
    return {
      restrict: 'AEC',
      scope: {
        source: '@',
        width: '=',
        height: '=',
        type: '='
      },
      templateUrl: 'components/qrcode/qrcode.html',
      controller: function($scope, $http, $rootScope, $location, $attrs, Notification) {


        $scope.$watch(
          "source",
          function handleFooChange() {
            if ($attrs.source != "{{product}}" && $attrs.source != undefined) {
              $scope.showLoading = false;
              $http({
                url: backend + '/identify/qr/' + $attrs.type,
                method: 'GET',
                headers: {
                  'code': JSON.parse($attrs.source).id,
                  'height': $attrs.height,
                  'width': $attrs.width
                }
              }).success(function(data, status, headers, config) {
                $scope.qr = data;
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

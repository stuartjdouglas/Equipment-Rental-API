'use strict';

angular.module('app.limage', ['app.config'])
  .directive('limage', function() {
    return {
      restrict: 'AEC',
      scope: {
        width: '=',
        height: '=',
        image: '@',
        preview: '='
      },
      templateUrl: 'components/image/image.html',
      controller: function($scope, $colorThief) {

        $scope.tap = function() {
          if ($scope.preview) {
            $scope.tapped = !$scope.tapped;
            var img = new Image();

            img.onload = function() {
              setTimeout(function(){
                $scope.imageStyles['background-image'] = 'url(\"' + $scope.url + '\")';
                $scope.$apply();
              }, 1);
            }

           img.src = $scope.url;
           if (img.complete) img.onload();

          }
        }

        $scope.previewStyle = {
          'height': '100%',
          'width':'100%',
          'position':'fixed',
          'top':0,
          'left':0,
          'background-color': 'rgba(0, 0, 0, 0.2)',
          'background-size': 'cover',
          'z-index':100,
          'padding': '-20%'
        }
        $scope.imageStyles = {
          'height':'100%',
          'width':'100%',
          'background': 'url("") no-repeat center center',
          'background-size': 'contain'
        }

        $scope.imageStyle = {
          'height': $scope.height + 'px',
          'width': $scope.width + 'px',
          'background': 'url("") no-repeat center center',
          'background-size': 'cover',
          'border-radius':$scope.height + 'px',
        };
        $scope.$watch(
          "image",
          function handleFooChange() {
             console.log($scope.height)

            if ($scope.image !== '{{product}}' && $scope.image !== 'undefined' && $scope.image !== '') {

                $scope.data = JSON.parse($scope.image);
                if ($scope.data.images !== undefined) {
                    $scope.data = $scope.data.images.size;
                } else {
                    $scope.data = $scope.data.image.size;
                }

                $scope.url = data + $scope.data.large;
                var url = data + $scope.data.thumb;

              // $scope.imageStyles.push('background-image': 'url(' + url + ')');
              //console.log(url);
              $scope.imageStyle['background-image'] = 'url(\"https://d13yacurqjgara.cloudfront.net/users/12755/screenshots/1037374/hex-loader2.gif\")';

              var img = new Image();
            img.crossOrigin = "Anonymous"
              img.onload = function() {
                setTimeout(function(){
                  $scope.imageStyle['background-image'] = 'url(\"' + url + '\")';
                  $scope.imageStyles['background-image'] = 'url(\"' + $scope.url + '\")';
                    $scope.imageStyle.border = '3px solid rgb(' + $colorThief.getColor(img) +  ')';
                  $scope.$apply();
                }, 1);
              }

             img.src = url;
             if (img.complete) img.onload();
            }

          }
        );
      },

      link: function(scope, elem, attrs, http) {


      }
    };
  });

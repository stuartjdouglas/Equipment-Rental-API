'use strict';

angular.module('app.rlabel', ['app.config'])
    .directive('rlabel', function() {
        return {
            restrict: 'AEC',
            scope: {
                input: '='
            },
            templateUrl: 'components/rLabel/rLabel.html',
            controller: function() {
                debugger
            },
            link: function(scope, elem, attrs) {
                // Just for altering the DOM
            }
        };
    });
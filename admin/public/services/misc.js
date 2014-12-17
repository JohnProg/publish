var contextMenuServices = angular.module('contextMenuServices', ['ngResource']);

contextMenuServices.factory('ContextMenuService', function() {
    return {
      element: null,
      menuElement: null
    };
  })


var underscoreServices = angular.module('underscoreServices', ['ngResource']); 

underscoreServices.factory('_', ['$window', function ($window) { return $window._; }]);
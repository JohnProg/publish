'use strict';

// Declare app level module which depends on components
angular.module('myApp', [
  'ui.router',
  'adminControllers',
  'authInterceptorService',
  'authenticationService',
  'userControllers',
  'userServices',
  'contentControllers',
  'mediaControllers',
  'contentTypeControllers',
  'mediaTypeControllers',
  'directoryControllers',
  //'stylesheetControllers',
  'dataTypeControllers',
  'templateControllers',
  'nodeServices',
  'entityServices',
  'contextMenuServices',
  'ui.utils',
  'checklist-model'
  // 'underscoreServices'
])
.config(function($stateProvider,$urlRouterProvider,$locationProvider, $httpProvider) {
	$stateProvider
		.state('adminLogin', {
			url: '/admin/login',
			templateUrl: 'public/views/admin/admin-login.html',
			data: {
				access: { requiredAuthentication: false }
			} 
		})
		// .state('adminIndex', {
		// 	url: '/admin',
		// 	abstract: true,
		// 	templateUrl: 'public/views/admin/index.html',
		// 	data: {
		// 		access: { requiredAuthentication: true }
		// 	}     
		// })
		.state('adminDashboard', {
			url: '/admin/',
			templateUrl: 'public/views/admin/dashboard.html',
			data: {
				access: { requiredAuthentication: true }
			}     
		})
		.state('adminContent', {
			url: '/admin/content',
			// abstract: true,
			templateUrl: 'public/views/content/index.html',
			data: {
				access: { requiredAuthentication: true }
			}     
		})
		.state('adminContent.edit', {
			url: '/edit/:nodeId',
			templateUrl: 'public/views/content/edit.html',
			data: {
				access: { requiredAuthentication: true }
			}     
		})
		.state('adminContent.new', {
			url: '/new?node_type&content_type_node_id&parent_id',
			templateUrl: 'public/views/content/new.html',
			data: {
				access: { requiredAuthentication: true }
			}     
		})
		.state('adminMedia', {
			url: '/admin/media',
			templateUrl: 'public/views/media/index.html',
			data: {
				access: { requiredAuthentication: true }
			}     
		})
		.state('adminMedia.edit', {
			url: '/edit/:nodeId',
			templateUrl: 'public/views/media/edit.html',
			data: {
				access: { requiredAuthentication: true }
			}     
		})
		.state('adminMedia.new', {
			url: '/new?node_type&content_type_node_id&parent_id',
			templateUrl: 'public/views/media/new.html',
			data: {
				access: { requiredAuthentication: true }
			}     
		})
		.state('adminUsers', {
			url: '/admin/users',
			templateUrl: 'public/views/users/index.html',
			data: {
				access: { requiredAuthentication: true }
			}     
		})
		.state('adminSettings', {
			url: '/admin/settings',
			templateUrl: 'public/views/settings/index.html',
			data: {
				access: { requiredAuthentication: true }
			}     
		})
		.state('adminSettings.contentType', {
			url: '/content-type',
			// abstract: true,
			templateUrl: 'public/views/settings/content-type/index.html',
			data: {
				access: { requiredAuthentication: true }
			}     
		})
		.state('adminSettings.contentType.edit', {
			url: '/edit/:nodeId',
			templateUrl: 'public/views/settings/content-type/edit.html',
			data: {
				access: { requiredAuthentication: true }
			}     
		})
		.state('adminSettings.contentType.new', {
			url: '/new?type&parent',
			templateUrl: 'public/views/settings/content-type/new.html',
			data: {
				access: { requiredAuthentication: true }
			}     
		})
		.state('adminSettings.mediaType', {
			url: '/media-type',
			// abstract: true,
			templateUrl: 'public/views/settings/media-type/index.html',
			data: {
				access: { requiredAuthentication: true }
			}     
		})
		.state('adminSettings.mediaType.edit', {
			url: '/edit/:nodeId',
			templateUrl: 'public/views/settings/media-type/edit.html',
			data: {
				access: { requiredAuthentication: true }
			}     
		})
		.state('adminSettings.mediaType.new', {
			url: '/new?type&parent',
			templateUrl: 'public/views/settings/media-type/new.html',
			data: {
				access: { requiredAuthentication: true }
			}     
		})
		.state('adminSettings.dataType', {
			url: '/data-type',
			// abstract: true,
			templateUrl: 'public/views/settings/data-type/index.html',
			data: {
				access: { requiredAuthentication: true }
			}     
		})
		.state('adminSettings.dataType.edit', {
			url: '/edit/:nodeId',
			templateUrl: 'public/views/settings/data-type/edit.html',
			data: {
				access: { requiredAuthentication: true }
			}     
		})
		.state('adminSettings.dataType.new', {
			url: '/new',
			templateUrl: 'public/views/settings/data-type/new.html',
			data: {
				access: { requiredAuthentication: true }
			}     
		})
		.state('adminSettings.template', {
			url: '/template',
			// abstract: true,
			templateUrl: 'public/views/settings/template/index.html',
			data: {
				access: { requiredAuthentication: true }
			}     
		})
		.state('adminSettings.template.edit', {
			url: '/edit/:nodeId',
			templateUrl: 'public/views/settings/template/edit.html',
			data: {
				access: { requiredAuthentication: true }
			}     
		})
		.state('adminSettings.template.new', {
			url: '/new?parent',
			templateUrl: 'public/views/settings/template/new.html',
			data: {
				access: { requiredAuthentication: true }
			}     
		})
		.state('adminSettings.script', {
			url: '/script',
			// abstract: true,
			templateUrl: 'public/views/settings/script/index.html',
			data: {
				access: { requiredAuthentication: true },
				rootdir: 'script' 
			}     
		})
		.state('adminSettings.script.edit', {
			url: '/edit/:name',
			templateUrl: 'public/views/settings/script/edit.html',
			data: {
				access: { requiredAuthentication: true }
			}     
		})
		.state('adminSettings.script.new', {
			url: '/new?type&parent',
			templateUrl: 'public/views/settings/script/new.html',
			data: {
				access: { requiredAuthentication: true }
			}     
		})
		.state('adminSettings.stylesheet', {
			url: '/stylesheet',
			// abstract: true,
			templateUrl: 'public/views/settings/stylesheet/index.html',
			data: {
				access: { requiredAuthentication: true },
				rootdir: 'stylesheet'
			}     
		})
		.state('adminSettings.stylesheet.edit', {
			url: '/edit/:name',
			templateUrl: 'public/views/settings/stylesheet/edit.html',
			data: {
				access: { requiredAuthentication: true }
			}     
		})
		.state('adminSettings.stylesheet.new', {
			url: '/new?type&parent',
			templateUrl: 'public/views/settings/stylesheet/new.html',
			data: {
				access: { requiredAuthentication: true }
			}     
		});
	$locationProvider.html5Mode(true);
	$httpProvider.interceptors.push('authInterceptorService');
})

.constant('_', window._)

.run(['$rootScope', '$state', 'authenticationService', '$location', '$window', '$q', function ($rootScope, $state, authenticationService, $location, $window, $q) {
	$rootScope._ = window._;
	$rootScope.$on("$stateChangeSuccess",function (event, toState, toParams, fromState, fromParams) {
		$rootScope.state = toState;
	});
	$rootScope.$on("$stateChangeStart",function (event, toState, toParams, fromState, fromParams) {

		if (toState != null && toState.data.access != null && toState.data.access.requiredAuthentication && !$window.sessionStorage.token) {

			if(!(authenticationService.authenticate().isAuthenticated)){

				// $location.path("/login")
				$state.go('adminLogin', toParams, {notify: false}).then(function() {
					
				    $rootScope.$broadcast('$stateChangeSuccess', toState, toParams, fromState, fromParams);
				});
				event.preventDefault();
		    }
		}
    });
}])

.filter('unsafe', function($sce) {
    return function(val) {
        return $sce.trustAsHtml(val);
    };
})

.filter('capitalize', function() {
    return function(input, all) {
      return (!!input) ? input.replace(/([^\W_]+[^\s-]*) */g, function(txt){return txt.charAt(0).toUpperCase() + txt.substr(1).toLowerCase();}) : '';
    }
  })

.filter('pathToUrl', function(){
  return function(text){
  	text = text.replace(/\\/g, '/');
  	return text;
  }
})

// .filter('unique', function() {

//   return function (arr, field) {
//     var o = {}, i, l = arr.length, r = [];
//     for(i=0; i<l;i+=1) {
//       o[arr[i][field]] = arr[i];
//     }
//     for(i in o) {
//       r.push(o[i]);
//     }
//     return r;
//   };
// })

.directive('wrapInput', [function () {
   return {
      replace: true,
      transclude: true,
      //template: '<div>{{prop.data_type.Html}}</div>'
      template: '<div class="input-wrapper" ng-transclude></div>'
   };
}])

.directive('compile',function($compile, $timeout){
    return{
        restrict:'A',
        link: function(scope,elem,attrs){
            $timeout(function(){                
                $compile(elem.contents())(scope);    
            });
        }        
    };
})

.directive('ngContextMenu', function($parse, $compile) {
	var offset = {
        left: 40,
        top: -20
    }
    return function(scope, element, attrs) {


    	//console.log(scope.menuOptions);
    	var template = "<ul>\
    				<li ng-repeat='option in currentItem.nodes'>\
    					<a>{{ option.label }}\
    						<ul ng-if='option.nodes' style='padding: 1em 0; list-style-type: none;'>\
    							<li ng-repeat='child in option.nodes'>\
									<a>{{ child.label }}</a>\
								</li>\
							</ul>\
						</a>\
					</li>\
				</ul>";
    	//var lol = '<ul><li ng-repeat="option in currentItem.nodes"><a>{{ option.label }} <ul ng-if="option.nodes" style="padding: 1em 0; list-style-type: none;"><li ng-repeat="child in option.nodes"><a>{{ child.label }}</a></li></ul></a></li></ul>';
    	var $oLay = angular.element(document.getElementById('overlay'))
    	
        var fn = $parse(attrs.ngContextMenu);

        // scope.showOptions = function (item,$event) {       
        //     var overlayDisplay;
        //     if (scope.currentItem === item) {
        //         scope.currentItem = null;
        //          overlayDisplay='none'
        //     }else{
        //          scope.currentItem = item;
        //         overlayDisplay='block'
        //     }
          
        //     var overLayCSS = {
        //         left: $event.clientX + offset.left + 'px',
        //         top: $event.clientY + offset.top + 'px',
        //         display: overlayDisplay
        //     }

        //      $oLay.css(overLayCSS)
        // }

        element.bind('contextmenu', function(event) {
        	//alert(scope.currentItem);

        	$oLay = angular.element(document.getElementById('overlay'))
            scope.$apply(function() {
            	if(scope.getEntityInfo != undefined)
            		scope.getEntityInfo(scope.data);
                event.preventDefault();
                event.stopPropagation();
                //$oLay.html('<p>showing options for: {{currentItem.label}}</p>').show();
                
                fn(scope, {$event:event});
                // $oLay.html(template).show();
                // $compile($oLay.contents())(scope);
            //     if(scope.currentItem!= null)
	           //      if('nodes' in scope.currentItem)
	        			// console.log(scope.currentItem.nodes)
            });
        });
    };
})

// Would IsolateScope be better here? 
// Rather than relying on a controller function? 
// Would it make more like a self contained component?
.directive('fileInput', ['$parse', function($parse){
	return {
		restrict: 'A',
		link: function(scope, elm, attrs){
  			if(typeof(scope.test) == undefined){
		      scope.test = { "files": []}
		    }
		    if(typeof(scope.test.files) !== undefined){
		      scope.test["files"] =[]
		    }
			elm.bind('change', function(){

				$parse(attrs.fileInput)
				.assign(scope,elm[0].files)
				scope.$apply()
			})
		}
	}
}]);

// .directive('ngContextMenu', [
// 	'$parse',
//     '$document',
//     'ContextMenuService',
//     function($parse, $document, ContextMenuService) {

//       return {
//         restrict: 'A',
//         scope: {
//           'callback': '&contextMenu',
//           'disabled': '&contextMenuDisabled'
//         },
//         link: function($scope, $element, $attrs) {
//         	alert($scope.menuOptions);
// 	        var data = $parse($attrs.ngContextMenu)($scope);
// 	        console.log(data);

// 	        var opened = false;
//           	function open(event, menuElement) {
// 	            menuElement.addClass('open');

// 	            var doc = $document[0].documentElement;
// 	            var docLeft = (window.pageXOffset || doc.scrollLeft) -
// 	                          (doc.clientLeft || 0),
// 	                docTop = (window.pageYOffset || doc.scrollTop) -
// 	                         (doc.clientTop || 0),
// 	                elementWidth = menuElement[0].scrollWidth,
// 	                elementHeight = menuElement[0].scrollHeight;
// 	            var docWidth = doc.clientWidth + docLeft,
// 	              docHeight = doc.clientHeight + docTop,
// 	              totalWidth = elementWidth + event.pageX,
// 	              totalHeight = elementHeight + event.pageY,
// 	              left = Math.max(event.pageX - docLeft, 0),
// 	              top = Math.max(event.pageY - docTop, 0);

// 	            if (totalWidth > docWidth) {
// 	              left = left - (totalWidth - docWidth);
// 	            }

// 	            if (totalHeight > docHeight) {
// 	              top = top - (totalHeight - docHeight);
// 	            }

// 	            menuElement.css('top', top + 'px');
// 	            menuElement.css('left', left + 'px');
// 	            opened = true;
//           	}

// 			function close(menuElement) {
// 				menuElement.removeClass('open');
// 				opened = false;
// 			}


// 	        $element.bind('contextmenu', function(event) {
// 				if (ContextMenuService.menuElement !== null) {
// 					close(ContextMenuService.menuElement);
// 				}
// 				ContextMenuService.menuElement = angular.element(
// 					document.getElementById($attrs.target)
// 				);
// 				ContextMenuService.element = event.target;
// 				//console.log('set', ContextMenuService.element);

// 				event.preventDefault();
// 				event.stopPropagation();
// 				$scope.$apply(function() {
// 					$scope.callback({ $event: event });
// 				});
// 				$scope.$apply(function() {
// 					open(event, ContextMenuService.menuElement);
// 				});
	                
// 	        });

// 	        function handleClickEvent(event) {
// 	            if (opened &&
// 	              (event.button !== 2 ||
// 	               event.target !== ContextMenuService.element)) {
// 	              $scope.$apply(function() {
// 	                close(ContextMenuService.menuElement);
// 	              });
// 	            }
//           	}


//           	// Firefox treats a right-click as a click and a contextmenu event
//           	// while other browsers just treat it as a contextmenu event
// 			$document.bind('click', handleClickEvent);
// 			$document.bind('contextmenu', handleClickEvent);

// 			$scope.$on('$destroy', function() {
// 				//console.log('destroy');
// 				$document.unbind('click', handleClickEvent);
// 				$document.unbind('contextmenu', handleClickEvent);
// 			});
// 	    }
//     };
// }]);

// .directive('ngRightClick', function($parse) {
//     return function(scope, element, attrs) {
//         var fn = $parse(attrs.ngRightClick);
//         element.bind('contextmenu', function(event) {
//             scope.$apply(function() {
//                 event.preventDefault();
//                 fn(scope, {$event:event});
//             });
//         });
//     };
// });

// .directive('contextMenu', [
//     '$document',
//     'ContextMenuService',
//     function($document, ContextMenuService) {
//       return {
//         restrict: 'A',
//         scope: {
//           'callback': '&contextMenu',
//           'disabled': '&contextMenuDisabled'
//         },
//         link: function($scope, $element, $attrs) {
//           var opened = false;

//           function open(event, menuElement) {
//             menuElement.addClass('open');

//             var doc = $document[0].documentElement;
//             var docLeft = (window.pageXOffset || doc.scrollLeft) -
//                           (doc.clientLeft || 0),
//                 docTop = (window.pageYOffset || doc.scrollTop) -
//                          (doc.clientTop || 0),
//                 elementWidth = menuElement[0].scrollWidth,
//                 elementHeight = menuElement[0].scrollHeight;
//             var docWidth = doc.clientWidth + docLeft,
//               docHeight = doc.clientHeight + docTop,
//               totalWidth = elementWidth + event.pageX,
//               totalHeight = elementHeight + event.pageY,
//               left = Math.max(event.pageX - docLeft, 0),
//               top = Math.max(event.pageY - docTop, 0);

//             if (totalWidth > docWidth) {
//               left = left - (totalWidth - docWidth);
//             }

//             if (totalHeight > docHeight) {
//               top = top - (totalHeight - docHeight);
//             }

//             menuElement.css('top', top + 'px');
//             menuElement.css('left', left + 'px');
//             opened = true;
//           }

//           function close(menuElement) {
//             menuElement.removeClass('open');
//             opened = false;
//           }

//           $element.bind('contextmenu', function(event) {
//             if (!$scope.disabled()) {
//               if (ContextMenuService.menuElement !== null) {
//                 close(ContextMenuService.menuElement);
//               }
//               ContextMenuService.menuElement = angular.element(
//                 document.getElementById($attrs.target)
//               );
//               ContextMenuService.element = event.target;
//               //console.log('set', ContextMenuService.element);

//               event.preventDefault();
//               event.stopPropagation();
//               $scope.$apply(function() {
//                 $scope.callback({ $event: event });
//               });
//               $scope.$apply(function() {
//                 open(event, ContextMenuService.menuElement);
//               });
//             }
//           });

//           function handleKeyUpEvent(event) {
//             //console.log('keyup');
//             if (!$scope.disabled() && opened && event.keyCode === 27) {
//               $scope.$apply(function() {
//                 close(ContextMenuService.menuElement);
//               });
//             }
//           }

//           function handleClickEvent(event) {
//             if (!$scope.disabled() &&
//               opened &&
//               (event.button !== 2 ||
//                event.target !== ContextMenuService.element)) {
//               $scope.$apply(function() {
//                 close(ContextMenuService.menuElement);
//               });
//             }
//           }

//           $document.bind('keyup', handleKeyUpEvent);
//           // Firefox treats a right-click as a click and a contextmenu event
//           // while other browsers just treat it as a contextmenu event
//           $document.bind('click', handleClickEvent);
//           $document.bind('contextmenu', handleClickEvent);

//           $scope.$on('$destroy', function() {
//             //console.log('destroy');
//             $document.unbind('keyup', handleKeyUpEvent);
//             $document.unbind('click', handleClickEvent);
//             $document.unbind('contextmenu', handleClickEvent);
//           });
//         }
//       };
//     }
//   ]);

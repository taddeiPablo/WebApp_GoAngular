var app = angular.module('main', ['ngRoute','oc.lazyLoad']);

app.config(['$routeProvider', '$httpProvider', function($routeProvider) {
  $routeProvider.
      when('/', {
        templateUrl: '/public/views/main.html',
        controller: 'indexController',
        resolve: {
            lazy: ['$ocLazyLoad', function($ocLazyLoad) {
                return $ocLazyLoad.load([{
                    name: 'index',
                    files: ['/public/controllers/main_controller.js']
                }]);
            }]
        }
      }).
      when('/data', {
        templateUrl: '/public/views/data.html',
        controller: 'searchController',
        resolve: {
            lazy: ['$ocLazyLoad', function($ocLazyLoad) {
                return $ocLazyLoad.load([{
                    name: 'search',
                    files: ['/app/controllers/search_controller.js']
                }]);
            }]
        }
      }).
      otherwise({
        redirectTo: '/'
      })
}]);
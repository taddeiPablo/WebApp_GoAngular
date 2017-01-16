'use strict';

var app = angular.module('main', ['ngRoute']);

/*aqui configuramos el ruteo atravez de angularjs*/
app.config(['$routeProvider', '$httpProvider', function($routeProvider) {
  $routeProvider.
      when('/', {
        templateUrl: '/public/views/main.html',
        controller: 'indexController'
      }).
      otherwise({
        redirectTo: '/'
      })
}]);
/*aqui el controller que corresponde a la view main*/
app.controller('indexController', ['$scope', function($scope){
    $scope.titulo = "Bienvenido a WebGO";
    $scope.subtitulo = "Framework creado en go y angularjs";
}]);

var index = angular.module('index', []);

index.controller('indexController', ['$scope', function($scope){
	$scope.titulo = "Bienvenidos a WebAppGo"
	$scope.subtitulo = "El Framework Web basado en go y angular mas simple de todos";
}]);


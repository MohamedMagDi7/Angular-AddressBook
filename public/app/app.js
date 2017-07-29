'use strict';

// Declare app level module which depends on views, and components
var app = angular.module('myApp', [
  'ngRoute',
  'ngAnimate',
  'myApp.login',
  'myApp.user',
  'myApp.contactus',
  'myApp.about',
  'myApp.version',
  'ngResource',
  'LocalStorageModule',
  'myApp.contactResource',
  'myApp.loginResource'
])

app.config(['$locationProvider', '$routeProvider', function($locationProvider, $routeProvider) {

  $locationProvider.hashPrefix('!');
  $routeProvider.when('/login', {
          disapleCache: true,
          template: '<login-view></login-view>'
        });
  $routeProvider.when('/usergo', {
          disapleCache: true,
          template: '<user-view></user-view>'
        });
   $routeProvider.when('/contactus', {
      template: '<contactus-view></contactus-view>'
    });
   $routeProvider.when('/about', {
         template: '<about-view></about-view>'
       });

  $routeProvider.otherwise({redirectTo: '/login'});

}]);

app.controller('AppHandler',  function AppHandler ($scope ,localStorageService ,sharedData, $location ){
$scope.UserIn = localStorageService.get("Userin");


});


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
  'LocalStorageModule'
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

app.factory('loginResource', ["$resource", function($resource){
        var serviceObject = {
        attemptToLogin : function (username, password){
        return $resource("/signin").save({}, {username: username, password: password}).$promise; //this promise will be fulfilled when the response is retrieved for this call
        },
        attemptToSignup : function (username, password){
        return $resource("/register").save({}, {username: username, password: password}).$promise; //this promise will be fulfilled when the response is retrieved for this call
        }
        };
        return serviceObject;
    }]);


app.factory('contactResource', ["$resource", function($resource){
        var addContact = $resource('/addnewcontact', {} , {
        Add : {method : 'POST' , headers: { 'Content-Type': 'application/x-www-form-urlencoded' } }
        });
        var serviceObject = {
        deleteNum : function (contactId , numberId){
            return $resource("/deletenumber").delete({contactid:contactId , contactnumber:numberId}).$promise; //this promise will be fulfilled when the response is retrieved for this call
        },
        deleteContact : function (contactId){
                    return $resource("/deletecontact").delete({contactid:contactId}).$promise; //this promise will be fulfilled when the response is retrieved for this call
        },
        addNewContact : function (contact){
                    console.log(contact);
                    return addContact.Add({}, contact).$promise; //this promise will be fulfilled when the response is retrieved for this call
        }
        };
        return serviceObject;
    }]);





loginResourceModule = angular.module('myApp.loginResource' , []);

loginResourceModule.factory('loginResource', ["$resource", function($resource){
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
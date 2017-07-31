


loginResourceModule = angular.module('myApp.loginResource' , []);

loginResourceModule.factory('loginResource', ["$resource", function($resource){
        var login = $resource('/signin', {} , {
                Post : {method : 'POST' , headers: { 'Content-Type': 'application/x-www-form-urlencoded' } }
        });
        var register = $resource('/register', {} , {
                        Post : {method : 'POST' , headers: { 'Content-Type': 'application/x-www-form-urlencoded' } }
                });
        var serviceObject = {
        attemptToLogin : function (logins){
        return login.Post({}, logins).$promise; //this promise will be fulfilled when the response is retrieved for this call
        },
        attemptToSignup : function (logins){
        return register.Post({}, logins).$promise; //this promise will be fulfilled when the response is retrieved for this call
        }
        };
        return serviceObject;
    }]);
/**
*  Module
*
* Description
*/
var app = angular.module('myApp.login', ['ngRoute' , 'ngAnimate'])

app.component('loginView', {
        templateUrl: 'login/login.html',
        controller: 'LoginCtrl'
      });


app.controller('LoginCtrl',  function LoginCtrl ($scope, $http , $location , localStorageService , loginResource ){

this.loginFlag = true;
this.aboutFlag = false;
this.contactusFlag = false;
this.Error = "";

this.SignIn = function(){
	if(this.loginForm.$valid)
	{
		loginResource.attemptToLogin(this.username , this.password).
		then(
		function mySuccess(response) {
           this.userdata = response;
           console.log(this.userdata);
           if(this.userdata.In == true)
            {
                localStorageService.set( "Userdata" ,this.userdata);
                $location.path('/usergo');
            }
           else{
                this.Error = this.userdata.Error;
               }
    }, function myError(response) {
        this.Error = response.statusText;
    });
	}
	else {
		this.Error = "No Request sent!!!";
	}
	};

this.Register = function(){
		if(this.loginForm.$valid)
	{
		loginResource.attemptToSignup( this.username , this.password ).
		then(
		function mySuccess(response) {
         this.userdata = response;
          console.log(response);

              if(this.userdata.In == true)
                {
                	    localStorageService.set( "Userdata" ,this.userdata);
                     $location.path('/usergo');
                 }
              else{
              	this.Error = this.userdata.Error;
              }
    }, function myError(response) {
        this.Error = "error";
    });
	}
	else {
		this.Error = "No Registeration happened";
	}
	};

this.ShowLogin = function() {
	   this.loginFlag = true;
       this.aboutFlag = false;
      this.contactusFlag = false;
	};

this.ShowAbout = function() {
    	   this.loginFlag = false;
           this.aboutFlag = true;
           this.contactusFlag = false;
    	};

this.ShowContactus = function() {
        	   this.loginFlag = false;
               this.aboutFlag = false;
               this.contactusFlag = true;
        };

})

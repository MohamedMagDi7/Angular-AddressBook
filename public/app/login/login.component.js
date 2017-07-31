/**
*  Module
*
* Description
*/
var loginModule = angular.module('myApp.login')

loginModule.component('loginView', {
        templateUrl: 'login/login.html',
        controller: 'LoginCtrl'
      });


loginModule.controller('LoginCtrl',  function LoginCtrl ($scope, $location , localStorageService , loginResource ){

$scope.loginFlag = true;
$scope.aboutFlag = false;
$scope.contactusFlag = false;
$scope.Error = "";

$scope.SignIn = function(){
	if($scope.loginForm.$valid)
	{
		loginResource.attemptToLogin($scope.username , $scope.password).
		then(
		function mySuccess(response) {
           $scope.userdata = response;
           console.log($scope.userdata);
           if($scope.userdata.In == true)
            {
                localStorageService.set( "Userdata" ,$scope.userdata);
                $location.path('/usergo');
            }
           else{
                $scope.Error = $scope.userdata.Error;
               }
    }, function myError(response) {
        $scope.Error = response.statusText;
    });
	}
	else {
		$scope.Error = "No Request sent!!!";
	}
	};

$scope.Register = function(){
		if($scope.loginForm.$valid)
	{
		loginResource.attemptToSignup( $scope.username , $scope.password ).
		then(
		function mySuccess(response) {
         $scope.userdata = response;
          console.log(response);

              if($scope.userdata.In == true)
                {
                	    localStorageService.set( "Userdata" ,$scope.userdata);
                     $location.path('/usergo');
                 }
              else{
              	$scope.Error = $scope.userdata.Error;
              }
    }, function myError(response) {
        $scope.Error = "error";
    });
	}
	else {
		$scope.Error = "No Registeration happened";
	}
	};

$scope.ShowLogin = function() {
	   $scope.loginFlag = true;
       $scope.aboutFlag = false;
      $scope.contactusFlag = false;
	};

$scope.ShowAbout = function() {
    	   $scope.loginFlag = false;
           $scope.aboutFlag = true;
           $scope.contactusFlag = false;
    	};

$scope.ShowContactus = function() {
        	   $scope.loginFlag = false;
               $scope.aboutFlag = false;
               $scope.contactusFlag = true;
        };

})

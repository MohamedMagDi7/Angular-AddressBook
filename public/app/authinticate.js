/**
*  Module
*
* Description
*/
var app = angular.module('myApp.login', ['ngRoute'])

app.config(['$routeProvider', function($routeProvider) {
      $routeProvider.when('/login', {
        templateUrl: 'login/login.html',
        controller: 'LoginCtrl'
      });
    }]);

app.controller('LoginCtrl',  function LoginCtrl ($scope , $http , $location , sharedData ){

	$scope.SignIn = function(){
	if($scope.loginForm.$valid)
	{
		$http({
        method : "POST",
        url : "/signin",
        data:
        	  {
        		username : $scope.username , 
        		password : $scope.password
       		  }
    }).then(function mySuccess(response) {
         $scope.userdata = response.data;
       	sharedData.setData(response.data);


       	if($scope.userdata.In == true)
       	{
            $location.path('/usergo');
       }
       else{
       	$scope.myError = $scope.userdata.Error;
       }
    }, function myError(response) {
        $scope.myWelcome = response.statusText;
    });
	}
	else {
		$scope.myWelcome = "No Request sent!!!";
	}
	};

	$scope.Register = function(){
		if($scope.loginForm.$valid)
	{
		$http({
        method : "POST",
        url : "/register" ,
        data: {
            username : $scope.username ,
        	password : $scope.password
              }
    }).then(function mySuccess(response) {
         	$scope.userdata=response.data;
       		sharedData.setData(response.data);
             if($scope.userdata.In == true)
              {
                    $location.path('/usergo');
              }
              else{
              	$scope.myError = $scope.userdata.Error;
              }
    }, function myError(response) {
        $scope.myWelcome = "error";
    });
	}
	else {
		$scope.myWelcome = "No Registeration happened";
	}
	};
})

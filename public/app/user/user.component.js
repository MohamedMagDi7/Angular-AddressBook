/**
*  Module
*
* Description
*/
var userModule = angular.module('myApp.user')

userModule.component('userView', {
        templateUrl: 'user/user.html',
        controller: 'UserCtrl'
      });

userModule.controller('UserCtrl',  function UserCtrl ($scope , $location ,contactResource , localStorageService){

 $scope.contactFlag = true ;
 $scope.userdata = localStorageService.get("Userdata");

if($scope.userdata == undefined)
    {
        $location.path('/login');
          return;
    }

$scope.formdata={};
$scope.numbers = $scope.formdata.phones=[{id: 'phone1'}];
$scope.contactFlag =false;
$scope.contacts = $scope.userdata.Userdata.Contacts ;

$scope.DeleteContact = function(ContactId) {
   contactResource.deleteContact(ContactId).
   then(
   function mySuccess(response) {
       if(response[0] === undefined){
            angular.forEach($scope.contacts , function(value, key) {
               if(value.Id === ContactId )
               {
                   $scope.contacts.splice(key, 1);
                   $scope.userdata.Userdata.Contacts=$scope.contacts;
                   localStorageService.set( "Userdata" ,$scope.userdata);
                     return;
               }
               });
               }
          else
          {
               $scope.Error = response.error;
                 return;
          }
      }, function myError(response) {
          $scope.Error = "error";
            return;
      });

}

$scope.DeleteNumber= function(NumberId , ContactId) {
  contactResource.deleteNum(ContactId , NumberId).
  then(
  function mySuccess(response) {
      if(response[0] === undefined){
           angular.forEach($scope.contacts , function(ContactValue, ContactKey) {
               if(ContactValue.Id === ContactId )
               {
                   angular.forEach(ContactValue.PhoneNumbers , function(value, key) {
                       if(key === NumberId )
                       {
                           $scope.contacts[ContactKey].PhoneNumbers.splice(key, 1);
                           $scope.userdata.Userdata.Contacts=$scope.contacts;
                           localStorageService.set( "Userdata" ,$scope.userdata);
                           return;
                       }
                       });
               }
               });
               }
               else{
               $scope.Error = response.error;
                 return;
               }
      }, function myError(response) {
          $scope.Error = "error";
            return;
      });

}

$scope.ShowAddContact = function() {
$scope.addcontactFlag = true;
$scope.contactFlag = false;
$scope.aboutFlag = false;
$scope.contactusFlag = false;
  return;
}

$scope.ShowContacts = function() {
$scope.addcontactFlag = false;
$scope.contactFlag = true;
$scope.aboutFlag = false;
$scope.contactusFlag = false;
  return;
}

$scope.ShowAbout = function() {
$scope.addcontactFlag = false;
$scope.contactFlag = false;
$scope.aboutFlag = true;
$scope.contactusFlag = false;
  return;
}

$scope.ShowContactus = function() {
$scope.addcontactFlag = false;
$scope.contactFlag = false;
$scope.aboutFlag = false;
$scope.contactusFlag = true;
  return;
}

$scope.Logout = function(){
    localStorageService.remove("Userdata");
    $location.path('/login');
      return;
}

$scope.AddNumberField = function(){
    var newItemNumber = $scope.formdata.phones.length+1;
    $scope.formdata.phones.push({'id':'phone'+newItemNumber});
      return;
}

$scope.RemoveNumberField = function(item){

    $scope.formdata.phones.splice(item,1);
      return;

}

$scope.AddContact = function(){
        if($scope.userForm.$valid)
        	{
        	    contactResource.addNewContact($("#contact-form").serialize()).
        	    then(
        	    function mySuccess(response) {
                                  console.log(response);
                                  if(!response.Id)
                                  {
                                  $scope.Error = "Error Adding";
                                  }
                                  else{
                                  if($scope.contacts === null)
                                  {
                                  $scope.contacts = [];
                                  }
                                  $scope.contacts.push(response);
                                  $scope.AfterSubmit();
                                  return;
                                  }

                             }, function myError(response) {

        		                    $scope.Error = "No Request sent!!!";

                              });
 }
 }

$scope.AfterSubmit = function () {
                                   $scope.userdata.Userdata.Contacts=$scope.contacts;
                                   localStorageService.set( "Userdata" ,$scope.userdata);
                                   $scope.addcontactFlag=false;
                                   $scope.userForm.$setUntouched();
                                   $scope.userForm.$setPristine();
                                   $scope.userForm.firstname.$dirty = false;
                                   $scope.userForm.lastname.$dirty = false;
                                   $scope.userForm.email.$dirty = false;
                                   $scope.formdata={};
                                   $scope.formdata.phones=[{id: 'phone1'}];
                                   return;
 }

});

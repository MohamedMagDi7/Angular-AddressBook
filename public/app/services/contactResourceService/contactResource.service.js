

contactResourceModule = angular.module('myApp.contactResource' , []);

contactResourceModule.factory('contactResource', ["$resource", function($resource){
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

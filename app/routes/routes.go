// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tApp struct {}
var App tApp


func (_ tApp) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Index", args).URL
}

func (_ tApp) Signin(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Signin", args).URL
}

func (_ tApp) Signup(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Signup", args).URL
}


type tUser struct {}
var User tUser


func (_ tUser) GetUser(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("User.GetUser", args).URL
}

func (_ tUser) AddNewContact(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("User.AddNewContact", args).URL
}

func (_ tUser) Delete(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("User.Delete", args).URL
}

func (_ tUser) DeleteNum(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("User.DeleteNum", args).URL
}

func (_ tUser) Logout(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("User.Logout", args).URL
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).URL
}

func (_ tTestRunner) Suite(
		suite string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	return revel.MainRouter.Reverse("TestRunner.Suite", args).URL
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).URL
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).URL
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).URL
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).URL
}



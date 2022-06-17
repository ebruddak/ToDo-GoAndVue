# Golang and Vue 3 Poc: Todo List Project

It is a project written using Goland in the backend and Vue 3 in the frondend. MongoDb (Online MongoDbAtlas) has been used as database. JWT is used for authorization operations. Fiber was used for write Rest API. Visual studio code was used as an IDE.

##Database Connection (.env)

```sh
MONGOURI=mongodb+srv://ebrududak:ebru123456@cluster0.mvtbu.mongodb.net/?retryWrites=true&w=majority
```

##Client Port

```sh
http://localhost:8080
```
##Running Client

```sh
npm run serve
```
##Backend Port

```sh
http://localhost:3000
```
##Running Backend

```sh
go run main.go
```
##Data Models

The project consists of three models. These; are user, group and todo. Collection contents are as follows.

##User Case

The user can log in and create a new record.<br/>
User can add and remove groups.<br/>
User can add new todo.<br/>
Todos can be marked as complete.<br/>
Todos can be prioritized and group assigned.<br/>

## Resources

 ![ScreenShot](https://i.ibb.co/5xwfsCb/9999999.png)<br/>

 
## Resources

https://www.youtube.com/watch?v=rx6CPDK_5mU&list=LL&index=5 <br/>
https://www.youtube.com/watch?v=vDrSxxLXLcM&list=LL&index=4 <br/>
https://www.udemy.com/course/working-with-vue-3-and-go/



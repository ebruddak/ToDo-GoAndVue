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
##Test User

```sh
email: ebrududak@hotmail.com
password: 1
```

##Data Models

The project consists of three models. These; are user, group and todo. Collection contents are as follows.

```sh
type Group struct {
	Id     primitive.ObjectID `json:"id,omitempty" bson:"id"`
	Name   string             `json:"name,omitempty" bson:"name"`
	UserId primitive.ObjectID `json:"userid,omitempty" bson:"userid"`
}
type Todo struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	Title    string             `json:"title,omitempty"`
	State    bool               `json:"state,omitempty"`
	Content  string             `json:"content,omitempty"`
	Priority string             `json:"priority,omitempty"`
	UserId   primitive.ObjectID `json:"userId,omitempty"`
	GroupId  primitive.ObjectID `json:"groupId,omitempty"`
	DueDate  time.Time          `json:"dueDate,omitempty"`
}
type User struct {
	Id       primitive.ObjectID `json:"id,omitempty" bson:"id"`
	UserName string             `json:"userName,omitempty" bson:"username"`
	Email    string             `json:"eMail,omitempty" bson:"email"`
	Password []byte             `json:"password,omitempty" bson:"password"`
}


```
##User Case

The user can log in and create a new record.<br/>
User can add and remove groups.<br/>
User can add new todo.<br/>
Todos can be marked as complete.<br/>
Todos can be prioritized and group assigned.<br/>

## ScreeenShots

 <a href="https://imgbb.com/"><img src="https://i.ibb.co/kmcCgZt/1.png" alt="1" border="0"></a><br/>
 <a href="https://imgbb.com/"><img src="https://i.ibb.co/4W6kFKZ/2.png" alt="2" border="0"></a><br/>
 <a href="https://ibb.co/sHmfKfx"><img src="https://i.ibb.co/C21cbcC/3.png" alt="3" border="0"></a><br/>
 <a href="https://ibb.co/ySFkJp7"><img src="https://i.ibb.co/RyT6WNd/4.png" alt="4" border="0"></a>

 
## Resources

https://www.youtube.com/watch?v=rx6CPDK_5mU&list=LL&index=5 <br/>
https://www.youtube.com/watch?v=vDrSxxLXLcM&list=LL&index=4 <br/>
https://www.udemy.com/course/working-with-vue-3-and-go/



#  User API Detali

### 회원가입 [  *GET*  ] 
클라이언트가 회원가입을 할 수 있는 API 입니다.

### Request
```URL
/auth/singup
```

```json
{
    "username": "JunBeomHan",
    "email": "email",
    "password": "1234!"
}
```

### Response 

[ *SUCCESS* ] HTTP Status **201**

```json
{
    "status": 201,
    "message": "Your membership has been successfully completed."
}
```

[ *FAIL* ] HTTP Status **404**

```json
{
    "status": 404,
    "message": "There are already registered or existing username."
}
```

---

### 가입 [  *GET* ] 
클라이언트가 가입할 수 있는 API 입니다.

### Request
```url
/auth/login
```


```json
{
    "username": "JunBeomHan",
    "email": "email",
    "password": "1234!"
}
```

### Response


[ *SUCCESS* ] HTTP Status **200**
```json
{
    "status": 200,
    "message": "Your subscription is complete."
}
```

[ *FAIL* ] HTTP Status **400**
```json
{
    "status": 400,
    "message": "There are no registered or existing user names."
}
```

---


### 로그아웃 [  *GET* ] 
- Auth관련은 후반에 작업합시다!
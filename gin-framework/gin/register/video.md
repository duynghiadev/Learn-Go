- **Lưu ý:** 3 video số `7, 8, 9` mình phải xem xong khoá mysql trước, để biết cách config server mysql rồi sau đó mới quay lại khoá Golang Gin Web Framework để xem tiếp video `7, 8, 9`

---

[video 7](https://www.youtube.com/watch?v=djU1_308M8E&list=PLDZ_9qD1hkzMdre6oedUdyDTgoJYq-_AY&index=7)

For the MySQL refresher please check out my playlist at [MySQL for Golang Web Dev](https://www.youtube.com/playlist?list=PLDZ_9qD1hkzN0VVRS7cemYtZPutyY8hea)

In this video we register users. Here are some of the things we check prior to registering users and provide helpful responses to users.

1. username length
2. username is only alphanumeric characters
3. password is of sufficient randomness and complexity
4. email syntax
5. no disposable email address
6. provide email domain suggestion if typed wrong
7. email address domain has Mx Record setup properly
8. make sure username doesn't already exist
   If all checks pass, a new user is created and then a verification email will be sent to user's email with a verification link inside. We cover how to handle this request and grab the values. Comparing passwords with saved password hash is covered as well.

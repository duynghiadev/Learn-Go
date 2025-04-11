# Email Verification API Documentation

## Overview

This API provides an endpoint to verify email addresses using the AfterShip email-verifier package. It checks for valid syntax, disposable emails, reachability via SMTP, and MX record presence.

## Base URL

```
http://localhost:8080
```

## Endpoints

### 1. GET `/verifyemail`

**Description** : Renders the email verification form.

**Response** :

- **200 OK** : Returns an HTML form (`ver-email.html`).

---

### 2. POST `/verifyemail`

**Description** : Verifies the provided email address for validity and reachability.

**Request Parameters** :

- **email** (string, required): The email address to be verified.
- in this video i using this email: duynghia@gmail.com

  **Request Example** :

```http
POST /verifyemail HTTP/1.1
Host: localhost:8080
Content-Type: application/x-www-form-urlencoded

email=test@example.com
```

**Response** :

| Status Code               | Description                                                                 |
| ------------------------- | --------------------------------------------------------------------------- |
| 200 OK                    | Email is valid and can be registered.                                       |
| 400 Bad Request           | Email has invalid syntax, is disposable, unreachable, or has no MX records. |
| 500 Internal Server Error | Server error occurred during verification.                                  |

**Response Example** (Success - 200 OK):

```html
<!DOCTYPE html>
<html>
  <head>
    <title>Email Verification Result</title>
  </head>
  <body>
    <h1>Email Verified Successfully</h1>
    <p>Email: test@example.com</p>
  </body>
</html>
```

**Response Example** (Failure - 400 Bad Request):

```html
<!DOCTYPE html>
<html>
  <head>
    <title>Verification Failed</title>
  </head>
  <body>
    <h1>Email Verification Failed</h1>
    <p>Message: email address syntax is invalid</p>
  </body>
</html>
```

---

## Error Handling

- If the email syntax is invalid: `400 Bad Request`
- If the email is from a disposable domain: `400 Bad Request`
- If the email is unreachable via SMTP: `400 Bad Request`
- If the domain lacks proper MX records: `400 Bad Request`
- If an internal server error occurs: `500 Internal Server Error`

---

## Technologies Used

- **Golang** with **Gin framework**
- **AfterShip email-verifier** for email validation
- **HTML templates** for UI rendering

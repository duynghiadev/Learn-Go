# Backend GitHub Trending API Documentation

## Authentication

### Sign Up

- **URL**: `/user/sign-up`
- **Method**: `POST`
- **Description**: Register a new user
- **Request Body**:

```json
{
  "email": "user@example.com",
  "password": "your_password"
}
```

### Sign In

- **URL**: `/user/sign-in`
- **Method**: `POST`
- **Description**: Login to get access token
- **Request Body**:

```json
{
  "email": "user@example.com",
  "password": "your_password"
}
```

## User Profile

All endpoints in this section require JWT Authentication

### Get Profile

- **URL**: `/user/profile`
- **Method**: `GET`
- **Description**: Get user profile information
- **Headers**:

```
Authorization: Bearer {your_jwt_token}
```

### Update Profile

- **URL**: `/user/profile/update`
- **Method**: `PUT`
- **Description**: Update user profile information
- **Headers**:

```
Authorization: Bearer {your_jwt_token}
```

- **Request Body**:

```json
{
  "name": "John Doe",
  "email": "john@example.com"
}
```

## GitHub Trending

Requires JWT Authentication

### Get Trending Repositories

- **URL**: `/github/trending`
- **Method**: `GET`
- **Description**: Get list of trending GitHub repositories
- **Headers**:

```
Authorization: Bearer {your_jwt_token}
```

## Bookmarks

All endpoints in this section require JWT Authentication

### List Bookmarks

- **URL**: `/bookmark/list`
- **Method**: `GET`
- **Description**: Get list of user's bookmarked repositories
- **Headers**:

```
Authorization: Bearer {your_jwt_token}
```

### Add Bookmark

- **URL**: `/bookmark/add`
- **Method**: `POST`
- **Description**: Add a repository to bookmarks
- **Headers**:

```
Authorization: Bearer {your_jwt_token}
```

- **Request Body**:

```json
{
  "repo": "username/repository"
}
```

### Delete Bookmark

- **URL**: `/bookmark/delete`
- **Method**: `DELETE`
- **Description**: Remove a repository from bookmarks
- **Headers**:

```
Authorization: Bearer {your_jwt_token}
```

- **Request Body**:

```json
{
  "repo": "username/repository"
}
```

## Authentication

All protected endpoints require a valid JWT token in the Authorization header. The token can be obtained through the sign-in endpoint.

Example of protected endpoint header:

```
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

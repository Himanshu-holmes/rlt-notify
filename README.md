# Real-Time Notification System

This project is a **real-time notification system** built using the Go programming language and the Echo web framework. It provides functionality for creating and retrieving notifications for users in a fast and efficient manner.

---

## Features

### 1. **Notification Creation**
- Endpoint: `POST /notify`
- Accepts a JSON payload to create a notification for a specific user.
- Example payload:
```json
{
  "userID": 123,
  "unreadMessage": {
    "count": 5
  },
  "unreadWorkRequest": null
}
```
- Returns a 201 Created response on success.

### 2. **Notification Retrieval**
- Endpoint: `GET /listen/:id`
- Retrieves all notifications for a specific user (identified by id).
- If no notifications are available, the server waits for new notifications using a signaling mechanism.

### 3. **Real-Time Notification Logic**
- Notifications are stored in memory using either a list-based or channel-based storage mechanism.
- A publish-subscribe system is implemented to notify clients when new notifications are available.

### 4. **Timeout Management**
- Both notification creation and retrieval use context-based timeouts to ensure operations do not run indefinitely.

## How It Works

**Client Sends a Notification:**
1. A client sends a POST request to `/notify` with the notification details.
2. The server processes the request and stores the notification using the Tuntun service.

**Client Listens for Notifications:**
1. A client sends a GET request to `/listen/:id` to retrieve notifications for a specific user.
2. If notifications are available, they are returned immediately.
3. If no notifications are available, the server waits for new notifications using a signaling mechanism.

**Notification Storage:**
- Notifications are stored in memory using either a list-based or channel-based storage implementation.
- The storage mechanism ensures efficient management and retrieval of notifications.

## Endpoints

### 1. POST /notify
**Description**: Creates a new notification for a user.

**Request Body**:
```json
{
  "userID": 123,
  "unreadMessage": {
    "count": 5
  },
  "unreadWorkRequest": null
}
```

**Response**:
- 201 Created: Notification successfully created.
- 400 Bad Request: Invalid request payload.

### 2. GET /listen/:id
**Description**: Retrieves notifications for a specific user.

**Response**:
- 200 OK: Returns a list of notifications.
- 204 No Content: No notifications available.

## Installation

1. Clone the repository:
```
git clone https://github.com/himanshu-holmes/rlt-notify.git
cd rlt-notify
```

2. Install dependencies:
```
go mod tidy
```

3. Run the server:
```
go run *.go
```

## Technologies Used
- **Go**: Programming language.
- **Echo**: Web framework for building HTTP servers.
- **In-Memory Storage**: For managing notifications.
- **Publish-Subscribe Mechanism**: For real-time notification delivery.

## Future Enhancements
- Add persistent storage (e.g., database) for notifications.
- Implement WebSocket support for real-time updates.
- Add authentication and authorization for endpoints.
- Improve error handling and logging.
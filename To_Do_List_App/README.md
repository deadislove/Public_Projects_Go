# To-do list app

### Test

- For Listing Tasks: Open your browser or use curl to make a GET request:

```
curl http://localhost:9090/tasks
```

- For Creating a Task: Use curl with a POST request:

```
curl -X POST -d '{"title":"New Task", "content":"Task content"}' -H "Content-Type: application/json" http://localhost:9090/tasks
```

- For Updating a Task: Use curl with a PUT request:

```
curl -X PUT -d '{"title":"Updated Task", "content":"Updated content"}' -H "Content-Type: application/json" http://localhost:9090/tasks/1
```

- For Deleting a Task: Use curl with a DELETE request:

```
curl -X DELETE http://localhost:9090/tasks/1
```
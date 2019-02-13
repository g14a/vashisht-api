## Vashisht API
#### This is the backend architecture of the technical fest of IIITDM Kancheepuram.

For clients, the following routes might be useful

1. "/events" --> GetAllEvents - Method "GET"
2. "/events/{id}" --> GetEventByID
3. "/events" --> AddEvent - Method "POST"
4. "/events" --> UpdateEvent - Method "PUT"
5. "/events/{id}" --> DeleteEvent - Method "DELETE"
6. "/events/{eventid}/users" --> GetUsersForEvent - Method "GET"

<br><br> 

1. "/users" --> AddUser - Method "POST"
2. "/users" --> GetAllUsers - Method "GET"
3. "/users/login --> Login - Method "POST"
4. "/users/{userid}/events" --> GetEventsOfUsers - Method "GET"
5. "/users/{userid}/events/{eventid}/register" --> AddRegistration - Method "POST"
6. "/users/{userid}/events/{eventid}/cancel" --> CancelRegistration - Method "DELETE"
7. "/users/{userid}/events/{eventid}/check" --> CheckIfUserRegisteredForEvent - Method "GET"
8. "/users/{mongoid}/events/{eventid}/checkMongoID" --> CheckIfUserRegisteredForEventByMongoID - Method "GET"


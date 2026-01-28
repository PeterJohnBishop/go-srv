# go-crypt

# websocket package

# hub
- there is a single instance of the hub that maintains a list of connected clients, and contains channels that register and deregister clients, and transfer data.
- when a client app connects to the websocket a new websocket client is created and registers with the hub
- the hub then waits for data

# client
- when a client app sends data to the websocket, the client readPump sends that data to hub.broadcast where's its processed - when the hub recieves data in the hub.broadcast channel it loops throgh the clients list and sends the message to each client's send channel
- when data arrives in the send channel the writePump sends that data to the websocket and off to the client app

# totp
- generate a secret 'room' key and email as a QR code
- 
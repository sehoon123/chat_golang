# chat_golang
To use this chat app, you can run the following steps:

Save the above code to a file named chat.go.
Open a terminal and navigate to the directory where you saved the file.
Run the command go build to compile the code into an executable file.
Run the command ./chat to start the chat app.
Open another terminal and run the command nc localhost 8080 to connect to the chat app.
Type messages in the terminal running nc and they should be printed in the terminal running the chat app.
You can also connect to the chat app using telnet instead of nc by running the command telnet localhost 8080.

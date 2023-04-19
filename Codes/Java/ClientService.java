package Codes.Java;

public class ClientService{
    
    // create a method that saves a client
    public void save(ClientModel client) {
        // create a connection
        ClientConnection connection = new ClientConnection();
        // connect to the database
        connection.connect();
        // insert a client
        connection.insert(client);
        // disconnect from the database
        connection.disconnect();
    }

    // create a method that updates a client
    public void update(ClientModel client) {
        // create a connection
        ClientConnection connection = new ClientConnection();
        // connect to the database
        connection.connect();
        // update a client
        connection.update(client);
        // disconnect from the database
        connection.disconnect();
    }

    // create a method that deletes a client
    public void delete(ClientModel client) {
        // create a connection
        ClientConnection connection = new ClientConnection();
        // connect to the database
        connection.connect();
        // delete a client
        connection.delete(client);
        // disconnect from the database
        connection.disconnect();
    }

    // create a method that returns a list of clients and export to a file
    public void list() {
        // create a connection
        ClientConnection connection = new ClientConnection();
        // connect to the database
        connection.connect();
        // get a list of clients
        connection.list();
        // export to Client.json file
        System.out.println("Exported to Client.json file!");

        // disconnect from the database
        connection.disconnect();
    }

    // create a method to write a json file with the list of clients
    public void writeJsonFile() {


    }




}

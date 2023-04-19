package Codes.Java;

public class ClientController {

    private ClientModel client;
    
    // create a new ClientModel object using the constructor
    public ClientController() {
        client = new ClientModel("John Doe", "john.doe@email.com", "123456");
    }

    // create a method to save a client
    public void save() {
        // create a service
        ClientService service = new ClientService();
        // save a client
        service.save(client);
    }

    // create a method to update a client
    public void update() {
        // create a service
        ClientService service = new ClientService();
        // update a client
        service.update(client);
    }

    // create a method to delete a client
    public void delete() {
        // create a service
        ClientService service = new ClientService();
        // delete a client
        service.delete(client);
    }

}

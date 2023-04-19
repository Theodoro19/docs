package Codes.Java;

public class ClientModel {

    private String name;
    private String email;
    private String password;
    
    // create a constructor
    public ClientModel(String name, String email, String password) {
        this.name = name;
        this.email = email;
        this.password = password;
    }

    // create a getter
    public String getName() {
        return name;
    }

    // create a setter
    public void setName(String name) {
        this.name = name;
    }

    // create a getter
    public String getEmail() {
        return email;
    }

    // create a setter
    public void setEmail(String email) {
        this.email = email;
    }

    // create a getter
    public String getPassword() {
        return password;
    }

    // create a setter
    public void setPassword(String password) {
        this.password = password;
    }

    // create a method to generato a json file
    public String toJson() {
        return "{" + "\"name\": \"" + name + "\", \"email\": \"" + email + "\", \"password\": \"" + password + "\"}";
    }

}
// create a quarkus project

import io.quarkus.runtime.Quarkus;
import io.quarkus.runtime.annotations.QuarkusMain;

@QuarkusMain
public class Quarkus {
    public static void main(String... args) {
        Quarkus.run(args);
    }

    // create a method that returns a string
    public String hello() {
        return "Hello World!";
    }
}

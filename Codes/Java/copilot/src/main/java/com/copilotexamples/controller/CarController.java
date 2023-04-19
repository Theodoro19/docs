package com.copilotexamples.controller;

import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;

import com.copilotexamples.persistence.model.Car;
import com.copilotexamples.service.CarService;

@RestController
@RequestMapping("/api/car")
public class CarController {
    
    // inject car service
    @Autowired
    private CarService carService;

    // returns all cars
    @RequestMapping(value = "/all", method = RequestMethod.GET)
    public List<Car> findAll() {
        return carService.findAll();
    }

    // returns a car by id
    @RequestMapping(value = "/{id}", method = RequestMethod.GET)
    public Car findOne(@PathVariable("id") long id) {
        return carService.findOne(id);
    }

    // saves a car
    @RequestMapping(value = "/save", method = RequestMethod.POST)
    public Car save(@RequestBody Car car) {
        return carService.save(car);
    }

    // deletes a car by id
    @RequestMapping(value = "/delete/{id}", method = RequestMethod.DELETE)
    public void delete(@PathVariable("id") long id) {
        carService.delete(id);
    }

    // returns a car by make
    @RequestMapping(value = "/make/{make}", method = RequestMethod.GET)
    public Car findByMake(@PathVariable("make") String make) {
        return carService.findByMake(make);
    }

    // returns a car by model
    @RequestMapping(value = "/model/{model}", method = RequestMethod.GET)
    public Car findByModel(@PathVariable("model") String model) {
        return carService.findByModel(model);
    }

    // returns a car by year
    @RequestMapping(value = "/year/{year}", method = RequestMethod.GET)
    public Car findByYear(@PathVariable("year") String year) {
        return carService.findByYear(year);
    }

    // returns a car by color
    @RequestMapping(value = "/color/{color}", method = RequestMethod.GET)
    public Car findByColor(@PathVariable("color") String color) {
        return carService.findByColor(color);
    }

    // returns a car by price
    @RequestMapping(value = "/price/{price}", method = RequestMethod.GET)
    public Car findByPrice(@PathVariable("price") String price) {
        return carService.findByPrice(price);
    }

    // returns a car by description
    @RequestMapping(value = "/description/{description}", method = RequestMethod.GET)
    public Car findByDescription(@PathVariable("description") String description) {
        return carService.findByDescription(description);
    }

}

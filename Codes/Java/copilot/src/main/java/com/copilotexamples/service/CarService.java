package com.copilotexamples.service;

import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import com.copilotexamples.persistence.model.Car;
import com.copilotexamples.persistence.repo.CarRepository;

@Component
public class CarService {

    @Autowired
    private CarRepository carRepository;

    // returns all cars
    public List<Car> findAll() {
        return carRepository.findAll();
    }

    // returns a car by id using the Car class
    public Car findOne(long id) {
        return carRepository.findOne(id);
    }

    // saves a car using the Car class
    public Car save(Car car) {
        return carRepository.save(car);
    }

    // deletes a car by id
    public void delete(long id) {
        carRepository.delete(id);
    }

    // returns a car by make using the Car class
    public Car findByMake(String make) {
        return carRepository.findByMake(make);
    }

    // returns a car by model using the Car class
    public Car findByModel(String model) {
        return carRepository.findByModel(model);
    }

    // returns a car by year using the Car class
    public Car findByYear(String year) {
        return carRepository.findByYear(year);
    }

    // returns a car by color using the Car class
    public Car findByColor(String color) {
        return carRepository.findByColor(color);
    }

    // returns a car by price using the Car class
    public Car findByPrice(String price) {
        return carRepository.findByPrice(price);
    }

    // returns a car by description using the Car class
    public Car findByDescription(String description) {
        return carRepository.findByDescription(description);
    }

}

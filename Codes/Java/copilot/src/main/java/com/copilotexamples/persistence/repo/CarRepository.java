package com.copilotexamples.persistence.repo;

import java.util.List;

import org.springframework.data.repository.CrudRepository;
import org.springframework.stereotype.Repository;

import com.copilotexamples.persistence.model.Car;

@Repository
public interface CarRepository extends CrudRepository<Car, Long>{
    
    Car findByMake(String make);
    
    Car findOne(long id);

    Car findByModel(String model);

    Car findByYear(String year);

    Car findByColor(String color);

    Car findByPrice(String price);

    Car findByDescription(String description);

    List<Car> findAll();
    
}

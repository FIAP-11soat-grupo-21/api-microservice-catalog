Feature: Category Management

  Scenario: Create a new category
    Given the category data is valid with name "Bebidas"
    When I send a request to create a new category
    Then the category should be created successfully
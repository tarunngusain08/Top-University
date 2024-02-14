## Top-University

The **Top-University** project retrieves information from a database of universities worldwide using the HTTP GET method. By querying [https://jsonmock.hackerrank.com/api/universities](https://jsonmock.hackerrank.com/api/universities), you can access all the records. The query result is paginated, and additional pages can be retrieved by appending `?page=num` to the query string, where `num` is the page number.

### Response Structure
The response is a JSON object with the following fields:
- **page**: The current page of the results. (Number)
- **per_page**: The maximum number of results returned per page. (Number)
- **total**: The total number of results. (Number)
- **total_pages**: The total number of pages with results. (Number)
- **data**: Either an empty array or an array with a single object that contains the universities' records.

### University Object Schema
Each university record in the `data` array follows this schema:
- **university**: The name of the university (String)
- **rank_display**: The rank of the university according to the 2022 QS Rankings (String).
- **score**: The score of the university according to the 2022 QS Rankings (Number).
- **type**: The type of university (String)
- **student_faculty_ratio**: The ratio of the number of students to the number of faculty. (Number)
- **international_students**: The number of international students (String).
- **faculty_count**: The number of faculty (String)
- **location**: An object containing the location details with the following schema:
  - **city**: (String)
  - **country**: (String)
  - **region**: (String)

### Function Description
The **highestInternationalStudents** function, when given the names of two cities as parameters, returns the name of the university with the highest number of international students in the first city. If the first city does not have a university within the data, it returns the university with the highest number of international students in the second city.

#### Function Parameters
- **firstCity**: Name of the first city (String)
- **secondCity**: Name of the second city (String)

#### Returns
- **string**: The name of the university with the highest number of international students.

#### Constraints
- There is always a university in one of the two cities.

<img width="1511" alt="Screenshot 2024-02-14 at 3 54 08 PM" src="https://github.com/tarunngusain08/Top-University/assets/36428256/f6d04f8e-e7b4-498e-8c04-c8239f09f206">

<img width="1510" alt="Screenshot 2024-02-14 at 4 12 12 PM" src="https://github.com/tarunngusain08/Top-University/assets/36428256/f4d5c262-fc58-40cc-bdb5-9688aa5f5c6b">

<img width="1512" alt="Screenshot 2024-02-14 at 5 30 36 PM" src="https://github.com/tarunngusain08/Top-University/assets/36428256/a9f96fc5-e142-4654-a319-7bb2c48b0f3c">

<img width="1508" alt="Screenshot 2024-02-14 at 4 11 52 PM" src="https://github.com/tarunngusain08/Top-University/assets/36428256/3acebe4f-fdcb-4b83-8515-21ff25e8edd5">

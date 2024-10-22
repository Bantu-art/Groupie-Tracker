# Groupie-Tracker-Visualization
![APIs](/static/images/RDME.png)

## Description

Groupie Tracker visualization is a project designed to provide detailed information about musical artists, including their concerts and members, in an interactive and visually appealing way. This project allows users to explore artist details and concert locations with ease.

## Authors
- [Brian Bantu](https://github.com/Bantu-art)
- [Antony Oduor](https://github.com/oduortoni/oduortoni) 

## Usage

### Requirements

- Go (version 1.15 or higher recommended)
- Access to a terminal or command prompt

### Running the Application

1. Clone the repository:
   ```bash
   git clone https://learn.zone01kisumu.ke/git/anoduor/groupie-tracker-visualizations
   cd groupie-tracker-visualization
   ```
2. Run the server:
   ```bash
    go run .  
    ```
3. Open a web browser and visit:
    ```bash
   http://localhost:9000
    ```

## Implementation Details

### Endpoints

- **GET /**: Serves the main page where users can see the landing page of our project.

- **GET /artists**: Displays a list of all available artists, allowing users to click on an artist to view more details.

- **GET /artists/{id}**: Serves a detailed page for a specific artist, including information about the artist's members, first album, and a list of concert dates and locations.

- **GET /search**: Allows users to search for artists based on their name or other attributes.


### Features

- **Artist Listing**: Users can browse a list of artists, each with a brief overview including the name and an image.

- **Artist Details**: Clicking on an artist's name takes users to a detailed page with information about the artist, such as the members, first album, and a list of concert dates and locations.

- **Concerts Overview**: Users can view a list of concert locations and dates for each artist, giving them a comprehensive look at where and when the artist has performed.

- **Responsive Design**: The project is designed to be responsive, ensuring that users can easily navigate and view artist information on both desktop and mobile devices.

- **Error Handling**: Users are presented with appropriate error messages if they attempt to access invalid artist URLs or if there is an issue retrieving artist data from the API.


  ### Algorithm

The Groupie Tracker project involves several key steps to retrieve and display artist information and concert details from an external API. Here’s a detailed breakdown:

1. **Receiving User Request**:
   - The main page (`GET /`) allows users to view a list of available artists. Each artist's name is linked to a detailed view.
   - When a user clicks on an artist's name, a `GET` request is sent to the `/artists/{id}` endpoint to fetch detailed information about that specific artist.

2. **Processing the URL**:
   - The server extracts the artist's ID from the URL path. This ID is then used to construct API requests to fetch the artist's details and related concert information.

3. **Fetching Artist Information**:
   - The server sends an API request to an external service (e.g., `https://groupietrackers.herokuapp.com/api/artists/{id}`) to retrieve the artist's detailed information, including name, image, first album, and members.
   - If the artist is found, the server processes the JSON response into Go structs that represent the artist's data.

4. **Fetching Concert Information**:
   - The server makes another API request to fetch the concert locations and dates associated with the artist (e.g., `https://groupietrackers.herokuapp.com/api/relation/{id}`).
   - This data is then processed into a map structure, where keys are locations, and values are slices of dates.

5. **Combining Data for Display**:
   - The artist information and concert data are combined into a single template data structure.
   - The server uses Go templates to render the artist's detail page, embedding the artist's information and a list of concerts into the HTML.

6. **Displaying the Result**:
   - The server returns the rendered HTML page, displaying the artist's name, image, first album, members, and a list of concert locations and dates.
   - Users can view this detailed information directly in their browser.

7. **Error Handling**:
   - The server checks for potential errors throughout the process, such as invalid artist IDs, missing data, or API request failures.
   - If an error occurs (e.g., the artist ID does not exist), the server responds with an appropriate error message, such as `404 Not Found` or `500 Internal Server Error`.


## HTTP Status Codes

- `200 OK`: Returned when the artist details and concert information are successfully retrieved and displayed.
- `404 Not Found`: Returned when a requested artist ID does not exist or when a non-existent page or resource is accessed.
- `400 Bad Request`: Returned when the client sends an invalid request, such as an incorrect URL format.
- `500 Internal Server Error`: Returned when the server encounters an unexpected error while processing the request, such as a failure to retrieve data from the external API.


### Technologies Used

- **Go**: The primary programming language used to build the server-side logic, manage routes, and handle API interactions.
- **HTML Templates**: Utilized for rendering dynamic content on the web pages, including artist details and concert information.
- **Net/HTTP Package**: Employed for handling HTTP requests and responses, enabling communication between the client and server.
- **JSON**: Used for data interchange between the server and the Groupie Trackers API, allowing the application to fetch and display artist and concert information.


## Contributing

Contributions to this project are welcome. Please fork the repository and submit a pull request with your changes.
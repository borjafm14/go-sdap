# web-client/web-client/README.md

# Web Client for Management Server

This project is a web client application that interacts with the management server to provide functionalities for user management, including logging in as an admin, viewing users, creating new users, and deleting existing users.

## Project Structure

```
web-client
├── public
│   ├── index.html          # Main HTML file for the web application
├── src
│   ├── components          # React components for the application
│   │   ├── Login.tsx      # Component for user authentication
│   │   ├── UserList.tsx    # Component to display the list of users
│   │   ├── CreateUser.tsx  # Component to create new users
│   │   └── DeleteUser.tsx  # Component to delete existing users
│   ├── services            # Service layer for API interactions
│   │   └── managementService.ts # Functions to interact with the management server
│   ├── App.tsx            # Main application component
│   ├── index.tsx          # Entry point for the React application
│   └── types              # TypeScript types and interfaces
│       └── index.ts
├── package.json            # npm configuration file
├── tsconfig.json           # TypeScript configuration file
└── README.md               # Project documentation
```

## Setup Instructions

1. **Clone the repository:**
   ```
   git clone <repository-url>
   cd web-client
   ```

2. **Install dependencies:**
   ```
   npm install
   ```

3. **Run the application:**
   ```
   npm start
   ```

4. **Open your browser and navigate to:**
   ```
   http://localhost:3000
   ```

## Usage

- **Login:** Use the admin credentials to log in.
- **View Users:** After logging in, you can view the list of users.
- **Create User:** Fill out the form to create a new user.
- **Delete User:** Select a user from the list and delete them.

## License

This project is licensed under the MIT License.
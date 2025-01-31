import React, { useState } from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import Login from './components/Login';
import UserList from './components/UserList';
import CreateUser from './components/CreateUser';
import DeleteUser from './components/DeleteUser';

const App: React.FC = () => {
    const [isLoggedIn, setIsLoggedIn] = useState(false);

    const handleLogin = () => {
        setIsLoggedIn(true);
    };

    const handleLogout = () => {
        setIsLoggedIn(false);
    };

    return (
        <Router>
            <div>
                <Switch>
                    <Route path="/" exact>
                        {isLoggedIn ? (
                            <UserList onLogout={handleLogout} />
                        ) : (
                            <Login onLogin={handleLogin} />
                        )}
                    </Route>
                    <Route path="/create-user">
                        {isLoggedIn ? <CreateUser /> : <Login onLogin={handleLogin} />}
                    </Route>
                    <Route path="/delete-user">
                        {isLoggedIn ? <DeleteUser /> : <Login onLogin={handleLogin} />}
                    </Route>
                </Switch>
            </div>
        </Router>
    );
};

export default App;